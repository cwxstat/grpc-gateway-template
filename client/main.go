package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pbExample "github.com/cwxstat/grpc-gateway-template/proto"
	structpb "google.golang.org/protobuf/types/known/structpb"

	"context"
	"flag"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "0.0.0.0:10000", "the address to connect to")
)

func getCert() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile("./certs/proto-certs.pem",
		"localhost")
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func main() {
	flag.Parse()
	// Set up a connection to the server.

	creds, err := getCert()
	if err != nil {
		log.Fatalf("could not find certs")
	}
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbExample.NewNamespaceServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateNamespace(ctx, &pbExample.CreateNamespaceRequest{
		Namespace: "client-test",
		Metadata:  &structpb.Struct{},
	})

	if err != nil {
		log.Fatalf("could not addUser: %v", err)
	}
	log.Printf("Create time: %s", r.CreateTime)

}
