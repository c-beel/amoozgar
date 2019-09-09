package main

import (
	"flag"
	"log"
	"time"
	"context"
	"google.golang.org/grpc"
	"github.com/c-beel/amoozgar/src/pkg/api"
	"fmt"
)

func main() {
	// get configuration
	address := flag.String("server", "localhost:8000", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := amoozgar.NewAmoozgarServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	fmt.Println(c, ctx)

}
