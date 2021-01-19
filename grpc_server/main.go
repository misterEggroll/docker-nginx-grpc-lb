// +build server

package main

import (
	"demo_grpc/protos"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := protos.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	protos.RegisterPingServer(grpcServer, &s)

	// start the server
	log.Printf("Started GRPC Server on Addr: \"%s\", Network \"%s\"...\n",
		lis.Addr().String(),
		lis.Addr().Network(),
	)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
