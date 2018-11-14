package importer

import (
	"strconv"

	pb "github.com/gnames/bhlindex/protob"
	"github.com/lib/pq"
	context "golang.org/x/net/context"
)

func (bhlServer) RequestNewColumn(ctx context.Context, name *pb.ColumnName) (*pb.ColumnId, error) {
	getNewId := "INSERT INTO import_cols (name) VALUES ($1) RETURNING id;"
	var colId int
	err := db.QueryRow(getNewId, name.GetName()).Scan(&colId)
	if err != nil {
		return &pb.ColumnId{Id: -1}, err
	}
	idStr := strconv.Itoa(colId)
	colName := pq.QuoteIdentifier("import_col_" + idStr)

	query := "ALTER TABLE page_imports ADD COLUMN " + colName + " jsonb; ALTER TABLE title_imports ADD COLUMN " + colName + "jsonb;"
	_, err = db.Exec(query)
	if err != nil {
		return &pb.ColumnId{Id:-1}, err
	}
	return &pb.ColumnId{Id:int32(colId)}, nil
}

func (bhlServer) GetColumns(context.Context, *pb.Void) (*pb.ColumnList, error) {
	rows, err := db.Query("SELECT id, name FROM import_cols;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols := make([]*pb.ColumnList_ColumnInfo, 0)
	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		ci := new(pb.ColumnList_ColumnInfo)
		ci.Id = &pb.ColumnId{Id:int32(id)}
		ci.Name = &pb.ColumnName{Name:name}
		cols = append(cols, ci)
	}
	return &pb.ColumnList{Columns:cols}, nil
}

func (bhlServer) GetColumnByName(ctx context.Context, name *pb.ColumnName) (*pb.ColumnId, error) {
	rows := db.QueryRow("SELECT id FROM import_cols WHERE name = $1;", name.GetName())
	var colId int
	if err := rows.Scan(&colId); err != nil {
		return nil, err
	}
	return &pb.ColumnId{Id:int32(colId)}, nil
}
