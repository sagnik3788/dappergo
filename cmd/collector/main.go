package main

import (
	"log"
	"net"

	"github.com/sagnik3788/dappergo/pkg/collector"
	"github.com/sagnik3788/dappergo/pkg/proto/tracerpb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4317")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	svc := collector.NewCollectorService()
	tracerpb.RegisterCollectorServer(grpcServer, svc)

	log.Println("Collector running on port 4317...")
	grpcServer.Serve(lis)

}
