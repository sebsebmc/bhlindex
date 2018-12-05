package main

import (
	"fmt"
	"io"
	"log"

	context "golang.org/x/net/context"

	pb "github.com/gnames/bhlindex/protob"
	"google.golang.org/grpc"
)

// ImporterClient is the client API for Importer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
// type ImporterClient interface {
// 	RequestNewColumn(ctx context.Context, in *ColumnName, opts ...grpc.CallOption) (*ColumnId, error)
// 	GetColumns(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ColumnList, error)
// 	GetColumnByName(ctx context.Context, in *ColumnName, opts ...grpc.CallOption) (*ColumnId, error)
// 	IngestPageData(ctx context.Context, opts ...grpc.CallOption) (Importer_IngestPageDataClient, error)
// 	IngestTitleData(ctx context.Context, opts ...grpc.CallOption) (Importer_IngestTitleDataClient, error)
// 	IngestPageNameData(ctx context.Context, opts ...grpc.CallOption) (Importer_IngestPageNameDataClient, error)
// }
// BHLIndexClient is the client API for BHLIndex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
// type BHLIndexClient interface {
// 	Ver(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Version, error)
// 	Pages(ctx context.Context, in *PagesOpt, opts ...grpc.CallOption) (BHLIndex_PagesClient, error)
// 	Titles(ctx context.Context, in *TitlesOpt, opts ...grpc.CallOption) (BHLIndex_TitlesClient, error)
// }

func listPages(client pb.BHLIndexClient, pages chan<- *pb.Page) {
	ctx := context.Background()
	stream, err := client.Pages(ctx, &pb.PagesOpt{WithText: false})
	if err != nil {
		log.Fatalf("Could not stream pages: %v", err)
	}
	for {
		page, err := stream.Recv()
		if err == io.EOF {
			close(pages)
			break
		}
		if err != nil {
			close(pages)
			log.Fatalf("Recv error: %v", err)
		}
		// log.Println(page.TitleId)
		pages <- page
	}
}

func getDummyColumnId(client pb.ImporterClient) *pb.ColumnId {
	ctx := context.Background()
	resp, err := client.GetColumnByName(ctx, &pb.ColumnName{Name: "dummy"})
	if err != nil {
		log.Printf("Could not get column id for dummy: %v\n", err)
	}
	if resp.GetId() != 0 {
		return resp
	}
	//Otherwise we need to create the column
	colResp, err := client.RequestNewColumn(ctx, &pb.ColumnName{Name: "dummy"})
	if err != nil {
		log.Printf("Cannot request new column: %v\n", err)
	}
	return colResp
}

func importDummyData(inClient pb.BHLIndexClient, outClient pb.ImporterClient) {
	ctx := context.Background()
	cid := getDummyColumnId(outClient)
	//Uh, I don't use channels much, so idk if this is a good buffer size
	pages := make(chan *pb.Page, 10)
	go listPages(inClient, pages)
	stream, err := outClient.IngestPageData(ctx)
	if err != nil {
		log.Fatalf("Cannot start import stream: %v", err)
	}
	count := 0
	for page := range pages {
		if err := stream.Send(
			&pb.PageData{
				Id:       page.GetId(),
				JsonData: fmt.Sprintf(`{"test":%d}`, count),
				Col:      &pb.ColumnId{Id: cid.GetId()},
			}); err != nil {
			// fmt.Printf("Error sending: %v\n", err)
		}
		count++
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv error: %v", err)
	}
	fmt.Printf("Success: %t\n", res.GetSuccessful())
}

func main() {
	inConn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to bhlindex server: %v", err)
	}
	defer inConn.Close()
	inClient := pb.NewBHLIndexClient(inConn)
	log.Printf("Connected to bhlindex server\n")
	outConn, err := grpc.Dial(":8889", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to bhlindex server: %v", err)
	}
	defer outConn.Close()
	outClient := pb.NewImporterClient(outConn)
	log.Printf("Connected to importer server\n")

	importDummyData(inClient, outClient)
}
