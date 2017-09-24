package models

import (
	"database/sql"
	"fmt"

	"github.com/GlobalNamesArchitecture/bhlindex"
	"github.com/lib/pq"
)

// Count returns number of rows in a table
func Count(db *sql.DB, table string) int {
	var count int
	q := fmt.Sprintf(`SELECT COUNT(*) FROM %s`, pq.QuoteIdentifier(table))
	err := db.QueryRow(q).Scan(&count)
	bhlindex.Check(err)
	return count
}

// Truncate removes all the rows from a table
func Truncate(db *sql.DB, table string) {
	q := fmt.Sprintf("TRUNCATE TABLE %s", pq.QuoteIdentifier(table))
	_, err := db.Query(q)
	bhlindex.Check(err)
}