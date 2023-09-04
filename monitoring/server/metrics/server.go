package metrics

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "monitoring/proto/gen/go/monitoring/v1"
	"monitoring/server/store"
	"sync"
	"time"
)

type MetricsServer struct {
	v1.UnimplementedMetricsServiceServer
}

var _ v1.MetricsServiceServer = &MetricsServer{}

var (
	counters     *store.ConcurrentMap[string, int64]
	once_counter sync.Once
)

func init() {
	once_counter.Do(func() {
		counters = store.NewConcurrentMap[string, int64]()
	})
}

func (m MetricsServer) CreateMetric(ctx context.Context, request *v1.CreateMetricRequest) (*v1.CreateMetricResponse, error) {
	switch request.Type {
	case v1.MetricType_COUNTER:
		counter := request.Metric.(*v1.CreateMetricRequest_Counter)
		counter.Counter.Timestamp = timestamppb.New(time.Now().UTC())
		log.Infof(ctx, "creating counter: %s=%d", counter.Counter.Name, counter.Counter.Value)
		found, ok := counters.Find(counter.Counter.Name)
		if !ok {
			counters.Store(counter.Counter.Name, counter.Counter.Value)
		} else {
			counter.Counter.Value += found
			counters.Store(counter.Counter.Name, counter.Counter.Value)
		}
		return &v1.CreateMetricResponse{
			Type:   request.Type,
			Metric: &v1.CreateMetricResponse_Counter{Counter: counter.Counter},
		}, nil
	default:
		log.Errorf(ctx, "unsupported metric type: %f", request.Type.String())
		return nil, errors.New("unsupported metric type")
	}
}

func (m MetricsServer) GetMetric(ctx context.Context, request *v1.GetMetricsRequest) (*v1.GetMetricsResponse, error) {
	switch request.Type {
	case v1.MetricType_COUNTER:

		counter := request.Metric.(*v1.CreateMetricRequest_Counter)
		log.Infof(ctx, "creating counter: %s=%d", counter.Counter.Name, counter.Counter.Value)
		counters.Store(counter.Counter.Name, counter.Counter.Value)
		return &v1.CreateMetricResponse{}, nil
	default:
		log.Errorf(ctx, "unsupported metric type: %f", request.Type.String())
		return nil, errors.New("unsupported metric type")
	}
}
