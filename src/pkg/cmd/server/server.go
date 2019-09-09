package cmd

import (
	"context"
	"github.com/c-beel/amoozgar/src/pkg/service"
	"github.com/c-beel/amoozgar/src/pkg/protocol/grpc"
	"github.com/c-beel/amoozgar/src/configman"
	"flag"
	"log"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	configFileAddress := flag.String("conf", "config.yaml", "The path to the service config file")
	autoMigrate := flag.Bool("migrate", true, "Auto-migrate models")
	flag.Parse()

	// get configuration
	cfg, err := configman.ImportConfigFromFile(*configFileAddress)
	if err != nil {
		log.Fatalf("Failed to parse config file with error %v", err)
	}

	api, err := service.NewAmoozgarServer(cfg)
	if err != nil {
		log.Fatalf("failed to start service : %v", err)
	}
	if *autoMigrate {
		log.Println("Starting auto migrate...")
		if err := api.AutoMigrate(); err != nil {
			log.Fatalf("failed to auto migrate : %v", err)
		}
		log.Println("Auto migrate done.")
	}
	return grpc.RunServer(ctx, *api, cfg.ListenPort)
}
