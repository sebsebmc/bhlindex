package importer

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/gnames/bhlindex"
	"github.com/gnames/bhlindex/protob"
	"google.golang.org/grpc"
)

var db *sql.DB

type bhlServer struct{}

func (bhlServer) IngestPageData(srv protob.Importer_IngestPageDataServer) error {

	for {
		req, err := srv.Recv()
		if err == io.EOF {
    		srv.SendAndClose(&protob.Success{Successful:true})
    		return nil
		}
		if err != nil {
			return err
		}

		id := req.GetId()
		json := req.GetJsonData()
		col := req.GetCol().GetId()
		//Cannot use a prepared statement because we are dynamically adding the column name
	    query := fmt.Sprintf("INSERT INTO page_imports (page_id, import_col_%d) VALUES ($1, $2) ON CONFLICT (page_id) DO UPDATE SET import_col_%d = EXCLUDED.import_col_%d;", col, col, col)
		_, err = db.Exec(query, id, json)
		if err != nil {
    		log.Printf("Error inserting: %v", err)
			return err
		}
	}
	return nil
}

// Bidirectional stream for ingesting data to be joined to a title id.
func (bhlServer) IngestTitleData(srv protob.Importer_IngestTitleDataServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
    		srv.SendAndClose(&protob.Success{Successful:true})
    		return nil
		}
		if err != nil {
			return err
		}

		id := req.GetId()
		json := req.GetJsonData()
		col := req.GetCol().GetId()
		//Cannot use a prepared statement because we are dynamically adding the column name
    	query := fmt.Sprintf("INSERT INTO title_imports (title_id, import_col_%d) VALUES ($1, $2) ON CONFLICT (title_id) DO UPDATE SET import_col_%d = EXCLUDED.import_col_%d;", col, col, col)
		_, err = db.Exec(query, id, json)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bhlServer) IngestPageNameData(srv protob.Importer_IngestPageNameDataServer) error {
    return fmt.Errorf("Not Implemented!")
}

func initDB() *sql.DB {
	db, err := bhlindex.DbInit()
	bhlindex.Check(err)
	return db
}

func Serve(port int) {
	srv := grpc.NewServer()
	db = initDB()
	var bhl bhlServer
	protob.RegisterImporterServer(srv, bhl)
	address := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen at %s\n%v", address, err)
	}
	log.Fatal(srv.Serve(l))
}
