package main

import (
	"fmt"
	"log"

	"github.com/ThisJohan/snapp-assignment/internal/app/logistics"
	"github.com/ThisJohan/snapp-assignment/internal/app/uber"
	"github.com/ThisJohan/snapp-assignment/pkg/di"
	"github.com/ThisJohan/snapp-assignment/pkg/env"
	"github.com/ThisJohan/snapp-assignment/pkg/grpcext"
	"google.golang.org/grpc"
)

type Config struct {
	Grpc      grpcext.Config
	Logistics logistics.Config
}

const (
	service = "3pl"
)

var (
	configs Config
)

func init() {
	if err := env.LoadConfig(service, &configs); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	fmt.Println(configs)

	runGrpcServer()
}

func runGrpcServer() {
	s := grpc.NewServer(
		di.GrpcProvide("Hi", "There"),
	)

	uber.NewService(s)
	logistics.NewService(s, configs.Logistics)

	if err := grpcext.Serve(s, configs.Grpc); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
