package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	user3 "github.com/msh2107/auth/internal/api/user"
	user2 "github.com/msh2107/auth/internal/service/user"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/msh2107/auth/internal/config"
	"github.com/msh2107/auth/internal/config/env"
	"github.com/msh2107/auth/internal/repository/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/msh2107/auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to parse gRPC config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to parse PG config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	repository := user.NewRepository(pool)
	service := user2.NewService(repository)

	imp := user3.NewImplementation(service)
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, imp)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("auth service listening on %s", grpcConfig.Address())

	<-interrupt

	log.Printf("Shutting down...")
	s.GracefulStop()

}
