package main

import (
	"log"
	"rader/grpcdemo/rpc"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	//noraml case
	// cc, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())

	//send or receive big messages
	// cc, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(math.MaxInt32),
	// grpc.MaxCallSendMsgSize(math.MaxInt32)))

	//using gzip to compress message
	cc, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure(), grpc.WithDecompressor(grpc.NewGZIPDecompressor()))

	if err != nil {
		log.Fatal("faild to dial")
	}
	grpcClient := rpc.NewHelloServiceClient(cc)
	req := new(rpc.HelloRequest)
	req.Greeting = "test"
	resp, err := grpcClient.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to say hello. %s", err.Error())
	}
	log.Println("connect to server")
	reply := resp.GetReply()
	log.Printf("message size:%d", len(reply))
}
