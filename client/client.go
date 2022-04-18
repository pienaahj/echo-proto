package main

import (
	"context"
	"fmt"

	"github.com/pienaahj/grpc-todd/echo-proto/client/echopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	// create a context
	ctx := context.Background()

	// create a connection by dialing it with a context
	// creds := credentials.NewClientTLSFromCert(nil, "test cert")

	conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials().Clone()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create a new echo server client
	e := echopb.NewEchoServerClient(conn)
	// grpc to server
	resp, err := e.Echo(ctx, &echopb.EchoRequest{
		Message: "Hi, Im a GRPC message!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Got from server: ", resp)
}