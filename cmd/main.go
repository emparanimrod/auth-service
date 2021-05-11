package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"auth/configs"
	"auth/core"
	"auth/endpoints"
	"auth/storage/postgres"
	"auth/transport/grpc_service"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file from current directory into env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error reading .env file: %v", err)
		return
	}

	// load env vars into struct
	envconfig, err := configs.GetEnvConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	// transform envvar config into app config
	config := core.GetConfig(*envconfig)

	// instantiate a postgres database instance
	database, err := postgres.NewDatabase(*config)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	postgres.Migrate(database)

	// initialize domain
	domain := core.New(database)

	// initialize endpoints layer
	serviceEndpoints := endpoints.New(domain, *config)

	// init our grpc service and get a grpc server instance
	grpcServer := grpc_service.NewServer(serviceEndpoints)

	// create tcp listener for grpc use
	port := fmt.Sprintf(":%v", config.GRPCPort)
	grpcListener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error opening tcp listener for grpc: %v", err)
		return
	}

	// serve grpc server on the tcp port
	log.Printf("GRPC server started on: %v", config.GRPCPort)
	if err = grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("error starting grpc server: %v", err)
		return
	}
}
