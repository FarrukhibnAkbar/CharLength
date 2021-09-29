package main

import(
	"os"
	"log"
	"context"
	pr "github.com/Farrukh/charlen/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost: 8000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil{
		log.Fatalf("%v", err)
	}

	defer conn.Close()

	stub := pr.NewCounterClient(conn)

	res, err := stub.CharLength(context.Background(), &pr.Request { Text: os.Args[1]})

	if err != nil{
		log.Fatalf("%v", err)
	}

	log.Printf("%+v", res)
}