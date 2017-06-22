package main

import (
	"log"
	"rader/grpcdemo/rpc"
	"strings"

	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	//normal case
	// grpcServer := grpc.NewServer()

	//send or receive big messages
	// grpcServer := grpc.NewServer(grpc.MaxSendMsgSize(math.MaxInt32), grpc.MaxRecvMsgSize(math.MaxInt32))

	//use grzip to compression message
	grpcServer := grpc.NewServer(grpc.RPCCompressor(grpc.NewGZIPCompressor()))

	rpc.RegisterHelloServiceServer(grpcServer, &Server{})
	l, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("failed to listen")
	}
	log.Println("server is listening")
	grpcServer.Serve(l)
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, r *rpc.HelloRequest) (*rpc.HelloResponse, error) {
	defer log.Println("msg has been sent to client")

	//msg has a bggier size than default 4M. To send this big message, the default max message size should be changed.
	bigMsg := strings.Repeat("s", 1024*1024*8)
	log.Println("get request from client")
	resp := new(rpc.HelloResponse)
	resp.Reply = bigMsg
	return resp, nil
}
