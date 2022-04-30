package main

import (
	"context"
	"fmt"
	"net"
	"src/gRPC/protos"
	"strings"

	"google.golang.org/grpc"
)

/*
type EchoServiceServer interface {
	GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedEchoServiceServer()
}
*/
type echoService struct {
	protos.UnimplementedEchoServiceServer
}

func (es *echoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	var builder strings.Builder
	builder.WriteString("received ")
	builder.WriteString(req.GetReq())
	str := builder.String()
	fmt.Println(str)
	return &protos.EchoResponse{Res: str}, nil
}

func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, new(echoService))
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
