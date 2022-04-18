package main

import (
	"context"
	"fmt"
	"net"

	"github.com/pienaahj/grpc-todd/echo-proto/server/echopb"
	"google.golang.org/grpc"
)

type EchoServer struct {
	echopb.UnimplementedEchoServerServer // needs this to implement the interface
}



// Echo creates the client 
func (e *EchoServer) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	// respond with:
	return &echopb.EchoResponse{
		Response: "My Echo: " + req.Message,
	}, nil
}


// newServer creates a new echopb server
func newServer() *EchoServer {
	return &EchoServer{}
}

func main() {
	// GRPC server does not listen to a port by default
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// set up GRPC server
	grpcServer := grpc.NewServer()

	// create the echo server
	srv := newServer()

	// register our echo server with the GRPC server
	echopb.RegisterEchoServerServer(grpcServer, srv)

	// serve the listener
	fmt.Println("Now serving at port 8080")
	err = grpcServer.Serve(lst)
	if err != nil {
		panic(err)
	}
}