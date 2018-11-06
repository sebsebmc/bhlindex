package importer

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/gnames/bhlindex/protob"
	"google.golang.org/grpc"
)

var db *sql.DB

type bhlServer struct{}

func (bhlServer) IngestPageData(srv protob.Importer_IngestPageDataServer) error {
	ctx := srv.Context()
	stmt, err := db.Perpare("INSERT INTO page_imports WHERE page_id=? and data = ?")
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

		id := req.PageId
		json := req.JsonData
		_, err := stmt.Exec(id, json)
		if err != nil {
    		return err
		}
	}
}

func (bhlServer) IngestTitleData(protob.Importer_IngestTitleDataServer) error {

}

func (bhlServer) IngestPageNameData(protob.Importer_IngestPageNameDataServer) error {

}

func Serve(port int, ver string) {
	srv := grpc.NewServer()
	db = initDB()
	var bhl bhlServer
	protob.RegisterImporterServer(srv, bhlServer)
	address := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen at %s\n%v", address, err)
	}
	log.Fatalf(srv.Serve(l))
}
