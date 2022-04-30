package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"src/gRPC/protos"

	"google.golang.org/grpc"
)

func main() {
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	c := protos.NewEchoServiceClient(cli)
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res, err := c.GetUnaryEcho(context.Background(), &protos.EchoRequest{Req: str})
		if err != nil {
			panic(err)
		}
		fmt.Println(res.GetRes())
	}
}
