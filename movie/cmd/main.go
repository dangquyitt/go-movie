package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dangquyitt/go-movie/gen"
	"github.com/dangquyitt/go-movie/movie/internal/business/movie"
	metadatagateway "github.com/dangquyitt/go-movie/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/dangquyitt/go-movie/movie/internal/gateway/rating/http"
	grpchandler "github.com/dangquyitt/go-movie/movie/internal/handler/grpc"
	"github.com/dangquyitt/go-movie/pkg/discovery"
	"github.com/dangquyitt/go-movie/pkg/discovery/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

const serviceName = "movie"

func main() {
	var cfg config
	f, err := os.Open("../configs/base.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.API.Port
	log.Printf("Starting the movie service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(ctx, instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	biz := movie.New(ratingGateway, metadataGateway)
	h := grpchandler.New(biz)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
