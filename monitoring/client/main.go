package main

import (
	"context"
	"fmt"
	"github.com/smartpcr/go-otel/pkg/ot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	v1 "monitoring/proto/gen/go/monitoring/v1"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	conn, err := grpc.Dial(
		"localhost:5000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	stopChan := make(chan os.Signal, 1) // Setup a signal handler for Ctrl+C
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(stopChan)
		cancel()
	}()

	// Create a new ticker that ticks every 5 seconds
	metricClient := v1.NewMetricsServiceClient(conn)
	metricTicker := time.NewTicker(5 * time.Second)
	defer metricTicker.Stop()

	// spans
	traceClient := v1.NewTracesServiceClient(conn)
	traceTicker := time.NewTicker(15 * time.Second)
	defer traceTicker.Stop()

	// logs
	logClient := v1.NewLogServiceClient(conn)
	logTicker := time.NewTicker(10 * time.Second)
	defer logTicker.Stop()

	// run until Ctrl+C
	go func() {
		select {
		case <-metricTicker.C:
			generateCounters(ctx, metricClient)
		case <-traceTicker.C:
			generateSpans(ctx, traceClient)
		case <-logTicker.C:
			generateLogs(ctx, logClient)
		case <-stopChan:
			cancel()
			log.Println("Received Ctrl+C. Exiting...")
			return
		}
	}()
}

func generateCounters(ctx context.Context, metricClient v1.MetricsServiceClient) {
	createMetricResp, err := metricClient.CreateMetric(ctx, &v1.CreateMetricRequest{
		Metric: &v1.CreateMetricRequest_Counter{
			Counter: &v1.Counter{
				Name:      "test_counter",
				LastValue: 1,
				Sum:       0,
				Count:     1,
				Timestamp: nil,
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to create metric: %v", err)
	}
	counter := createMetricResp.Metric.(*v1.CreateMetricResponse_Counter).Counter
	log.Printf(
		"created counter of type %s: %s with value %d",
		createMetricResp.Type.String(),
		counter.Name,
		counter.LastValue,
	)
}

func generateSpans(ctx context.Context, traceClient v1.TracesServiceClient) {
	ctx, span, _ := ot.StartSpanLogger(ctx)
	defer span.End()
	span.AddEvent("span1")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter.

		go func(i int) {
			defer wg.Done()
			generateNestedSpans(ctx, i, traceClient)
		}(i)
	}
	wg.Wait()
}

func generateNestedSpans(ctx context.Context, index int, traceClient v1.TracesServiceClient) {
	_, childSpan, _ := ot.StartSpanLogger(ctx)
	defer childSpan.End()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	childSpan.AddEvent(fmt.Sprintf("child_span_%d", index))
}

func generateLogs(ctx context.Context, logClient v1.LogServiceClient) {
	logResp, err := logClient.Log(ctx, &v1.LogMessage{
		Level:   v1.ProtoLogLevel_INFO,
		Message: "test log",
	})
	if err != nil {
		log.Fatalf("failed to log: %v", err)
	}
	log.Printf("logged: %s", logResp.Message)
}
