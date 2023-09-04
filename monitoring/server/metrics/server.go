package metrics

import (
	"context"
	"go.opentelemetry.io/otel/sdk/metric"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "monitoring/proto/gen/go/monitoring/v1"
	"monitoring/server/store"
)

type MetricsServer struct {
	v1.UnimplementedMetricsServiceServer
	Meter *metric.MeterProvider
}

var _ v1.MetricsServiceServer = &MetricsServer{}

var (
	counters *store.ConcurrentMap[string, *v1.Counter]
)

func init() {
	counters = store.NewConcurrentMap[string, *v1.Counter]()
}

func (m MetricsServer) CreateMetric(ctx context.Context, request *v1.CreateMetricRequest) (*v1.CreateMetricResponse, error) {
	switch request.Type {
	case v1.MetricType_COUNTER:
		counter := request.Metric.(*v1.CreateMetricRequest_Counter).Counter
		log.Infof(ctx, "creating counter: %s=%d", counter.Name, counter.LastValue)
		found, ok := counters.Find(counter.Name)
		if !ok {
			counter.Count = 1
			counters.Store(counter.Name, &v1.Counter{
				Name:      counter.Name,
				Sum:       counter.LastValue,
				Count:     1,
				Timestamp: timestamppb.New(time.Now().UTC()),
			})
		} else {
			counters.Store(counter.Name, &v1.Counter{
				Name:      counter.Name,
				Sum:       counter.LastValue + found.Sum,
				Count:     found.Count + 1,
				Timestamp: timestamppb.New(time.Now().UTC()),
			})
		}
		return &v1.CreateMetricResponse{
			Type:   request.Type,
			Metric: &v1.CreateMetricResponse_Counter{Counter: counter},
		}, nil
	default:
		log.Errorf(ctx, "unsupported metric type: %f", request.Type.String())
		return nil, errors.New("unsupported metric type")
	}
}

func (m MetricsServer) GetMetric(ctx context.Context, request *v1.GetMetricsRequest) (*v1.GetMetricsResponse, error) {
	switch request.Type {
	case v1.MetricType_COUNTER:
		var metrics []*v1.Metric
		for _, counter := range counters.Values() {
			metrics = append(metrics, &v1.Metric{
				Metric: &v1.Metric_Counter{
					Counter: counter,
				},
			})
		}
		return &v1.GetMetricsResponse{
			Metrics: metrics,
		}, nil
	default:
		log.Errorf(ctx, "unsupported metric type: %f", request.Type.String())
		return nil, errors.New("unsupported metric type")
	}
}
