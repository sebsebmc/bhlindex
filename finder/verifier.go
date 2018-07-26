package finder

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gnames/bhlindex"
	"github.com/gnames/bhlindex/models"
	"github.com/gnames/gnfinder/resolver"
	gfutil "github.com/gnames/gnfinder/util"
	"github.com/lib/pq"
)

const BATCH_SIZE = 50000

var empty struct{}

func Verify(db *sql.DB, foundNames <-chan []models.DetectedName,
	wgFinish *sync.WaitGroup) {

	defer wgFinish.Done()

	namesVerified := make(chan bool)
	counter := make(chan int)

	go populateUniqueNames(db, foundNames)
	go verifyLog(counter)
	go verifyNames(db, counter, namesVerified)

	// wait till all names are imported
	<-namesVerified
	close(counter)
}

func verifyNames(db *sql.DB, counter chan<- int, namesVerified chan<- bool) {
	time.Sleep(20 * time.Second)
	for {
		time.Sleep(200 * time.Microsecond)
		verifyNamesQuery(db, counter)
		status := readStatus(db)
		if status.Status == AllNamesVerified {
			break
		}
	}
	namesVerified <- true
}

func verifyNamesQuery(db *sql.DB, counter chan<- int) int {
	m := gfutil.NewModel()
	m.Workers = 15
	q := `
		WITH temp AS (
			SELECT name FROM name_statuses
				WHERE processed=false LIMIT $1)
		UPDATE name_statuses ns SET processed=true
		FROM temp
			WHERE ns.name = temp.name
		RETURNING temp.name`

	rows, err := db.Query(q, BATCH_SIZE)
	gfutil.Check(err)
	defer rows.Close()

	var name string
	names := make([]string, 0, BATCH_SIZE)

	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	err = rows.Err()
	bhlindex.Check(err)

	namesSize := len(names)

	status := readStatus(db)
	if status.Status == AllNamesVerified && namesSize == 0 {
		updateStatus(db, &metaData{Status: AllErrorsProcessed})
	} else if status.Status == AllNamesHarvested && namesSize == 0 {
		updateStatus(db, &metaData{Status: AllNamesVerified})
	}

	counter <- namesSize
	time1 := time.Now().UnixNano()
	verified := resolver.Verify(names, m)
	if namesSize > 0 {
		timeSpent := float64(time.Now().UnixNano()-time1) / 1000000000
		speed := int(float64(namesSize) / timeSpent)
		log.Printf("\033[40;32;1mVerification speed %d names/sec\033[0m\n", speed)
		saveVerifiedNameStrings(db, verified)
	}
	return namesSize
}

func verifyLog(counter <-chan int) {
	total := 0
	count_same := 0
	for i := range counter {
		total += i
		if i > 0 {
			count_same = 0
			log.Printf("\033[40;32;1mVerifying %d name-strings\033[0m\n", total)
		} else {
			count_same++
			if count_same%100 == 0 {
				fmt.Printf("\033[40;31;1mVerification que is empty %d times\033[0m\n",
					count_same)
			}
		}
	}
}

func saveVerifiedNameStrings(db *sql.DB, verified resolver.VerifyOutput) {
	var errStr sql.NullString
	now := time.Now()
	columns := []string{"name", "match_type", "edit_distance",
		"matched_name", "current_name", "classification", "datasource_id",
		"datasources_number", "retries", "error", "updated_at"}
	transaction, err := db.Begin()
	bhlindex.Check(err)

	stmt, err := transaction.Prepare(pq.CopyIn("name_strings", columns...))
	bhlindex.Check(err)

	for name, v := range verified {
		if v.Error == nil {
			errStr = sql.NullString{}
		} else {
			errStr.Scan(v.Error.Error())
		}
		_, err = stmt.Exec(name, v.MatchType, v.EditDistance, v.MatchedName,
			v.CurrentName, v.ClassificationPath, v.DataSourceID, v.DatabasesNum,
			v.Retries, errStr, now)
		bhlindex.Check(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(`
Bulk import of titles data failed, probably you need to empty all data
and start with empty database.`)
		log.Fatal(err)
	}

	err = stmt.Close()
	bhlindex.Check(err)

	err = transaction.Commit()
	bhlindex.Check(err)
}

func populateUniqueNames(db *sql.DB, foundNames <-chan []models.DetectedName) {
	batch := 50000
	names := make([]string, 0, batch)
	uniqueNames := make(map[string]struct{})
	count := -1
	for fNames := range foundNames {
		for _, n := range fNames {
			name := n.NameString
			if _, ok := uniqueNames[name]; ok {
				continue
			}

			uniqueNames[name] = empty
			count++
			if count < batch {
				names = append(names, name)
			} else {
				count = -1
				uniqueNames = make(map[string]struct{})
				insertToUnique(db, names)
			}
		}
	}
	insertToUnique(db, names)
	updateStatus(db, &metaData{Status: AllNamesHarvested})
}

func insertToUnique(db *sql.DB, names []string) {
	namesJoined := strings.Join(names, "'),\n('")

	q := fmt.Sprintf(
		`INSERT INTO name_statuses (name)
				VALUES ('%s')
				ON CONFLICT (name) DO NOTHING`, namesJoined)

	stmt, err := db.Prepare(q)
	bhlindex.Check(err)

	_, err = stmt.Exec()
	bhlindex.Check(err)

	err = stmt.Close()
	bhlindex.Check(err)
}
