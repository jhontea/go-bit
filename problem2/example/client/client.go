package main

import (
	"context"
	"fmt"
	"log"

	"go-bit/entities/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	exampleCallSearch(conn)
	// exampleCallGetDetail(conn)
}

func exampleCallSearch(conn *grpc.ClientConn) {
	fmt.Println("call search")
	imdbService := proto.NewIMDBServiceClient(conn)
	req := proto.SearchRequest{
		Search: "Batman",
		Page:   1,
	}
	result, err := imdbService.Search(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result is", result)
}

func exampleCallGetDetail(conn *grpc.ClientConn) {
	fmt.Println("call get detail")
	imdbService := proto.NewIMDBServiceClient(conn)
	req := proto.GetDetailRequest{
		Id: "tt2011118",
	}
	result, err := imdbService.GetDetail(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result is", result)
}
