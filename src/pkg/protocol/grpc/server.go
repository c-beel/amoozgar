package grpc

import (
	"net"
	"os"
	"os/signal"
	"context"
	"log"
	"google.golang.org/grpc"
	"github.com/c-beel/amoozgar/src/pkg/api"
)

func RunServer(ctx context.Context, api amoozgar.AmoozgarServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	amoozgar.RegisterAmoozgarServiceServer(server, api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("Starting gRPC server...")
	return server.Serve(listen)
}
