package main

import (
	"context"
	"fmt"
	"github.com/smartpcr/go-otel/pkg/ot"
	"log"
	"monitoring/server/logs"
	"monitoring/server/metrics"
	"monitoring/server/traces"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	v1 "monitoring/proto/gen/go/monitoring/v1"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	// Run until signaled or the context expires.
	go func() {
		select {
		case <-c:
			fmt.Println("CTRL-C")
			gracefulStop()
			cancel()
		case <-ctx.Done():
			gracefulStop()
		}
	}()

	fmt.Println("registering logger")
	logger := ot.RegisterLogger(ctx)
	logger.Infof("starting %s", ServiceName)

	logger.Infof("registering tracing at %s", config.Receiver.Endpoint)
	if err := ot.RegisterTracing(ctx, config.Receiver.Endpoint, ServiceName, logger); err != nil {
		panic(err)
	}
	ctx, span, logger := ot.StartSpanLogger(ctx, "startup")
	defer span.End()
	span.AddEvent("startup")

	logger.Infof("registering metrics at %s", config.Receiver.Endpoint)
	metric, err := ot.RegisterOtelMetrics(ctx, config.Receiver.Endpoint, ServiceName)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterLogServiceServer(s, &logs.LogServer{
		Logger: logger,
	})
	v1.RegisterMetricsServiceServer(s, &metrics.MetricsServer{
		Meter: metric,
	})
	v1.RegisterTracesServiceServer(s, &traces.TraceServer{})

	log.Println("Starting gRPC server on :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func gracefulStop() {
	fmt.Println("stopping server...")
}
