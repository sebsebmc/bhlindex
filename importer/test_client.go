package main

import (
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

func listPages(client pb.BHLIndexClient) {
	ctx := context.Background()
	stream, err := client.Pages(ctx, &pb.PagesOpt{WithText: false})
	if err != nil {
		log.Fatalf("Could not stream pages: %v", err)
	}
	for {
		page, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Recv error: %v", err)
		}
		log.Println(page.TitleId)
	}
}

func main() {
	inConn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to bhlindex server: %v", err)
	}
	defer inConn.Close()
	inClient := pb.NewBHLIndexClient(inConn)
	listPages(inClient)
}
