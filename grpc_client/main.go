// +build client

package main

import (
	"demo_grpc/protos"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	targetPort = flag.String("port", "10000", "Specify port to send request to")
)

func main() {
	var conn *grpc.ClientConn
	flag.Parse()
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *targetPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := protos.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &protos.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
