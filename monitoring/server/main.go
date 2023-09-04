package main

import (
	"log"
	"monitoring/server/metrics"
	"net"

	"google.golang.org/grpc"

	v1 "monitoring/proto/gen/go/monitoring/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterMetricsServiceServer(s, &metrics.MetricsServer{})

	log.Println("Starting gRPC server on :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
