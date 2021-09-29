package main

import(
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	pr "github.com/Farrukh/charlen/proto"
)

type CounterServer struct{
	pr.UnimplementedCounterServer
}

func (s *CounterServer) CharLength(ctx context.Context, in *pr.Request) (*pr.Response, error){
	
	return &pr.Response { Length: int64(len(in.GetText())),}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()

	pr.RegisterCounterServer(s, &CounterServer {})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("%v", err)
	}
}