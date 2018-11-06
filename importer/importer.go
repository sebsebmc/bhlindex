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
	ctx := srv.Context()
	stmt, err := db.Prepare("INSERT INTO page_imports WHERE page_id=? and data = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err != io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		id := req.Id
		json := req.JsonData
		_, err = stmt.Exec(id, json)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bhlServer) IngestTitleData(protob.Importer_IngestTitleDataServer) error {
	ctx := srv.Context()
	stmt, err := db.Prepare("INSERT INTO title_imports WHERE title_id=? and data = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err != io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		id := req.Id
		json := req.JsonData
		_, err = stmt.Exec(id, json)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bhlServer) IngestPageNameData(protob.Importer_IngestPageNameDataServer) error {
	ctx := srv.Context()
	stmt, err := db.Prepare("INSERT INTO page_name_imports WHERE page_id=? and data = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err != io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		id := req.Id
		json := req.JsonData
		_, err = stmt.Exec(id, json)
		if err != nil {
			return err
		}
	}
	return nil
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
