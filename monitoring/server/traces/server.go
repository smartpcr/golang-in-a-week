package traces

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "monitoring/proto/gen/go/monitoring/v1"
	"monitoring/server/store"
	"time"
)

type TraceServer struct {
	v1.UnimplementedTracesServiceServer
}

var _ v1.TracesServiceServer = &TraceServer{}

var (
	spans *store.ConcurrentTree[*v1.Span]
)

func init() {
	spans = store.NewConcurrentTree[*v1.Span](&v1.Span{
		TraceId:      "root",
		SpanId:       "root",
		ParentSpanId: "",
		Name:         "root",
		StartTime:    timestamppb.New(time.Now().UTC()),
	})
}

func (t TraceServer) CreateSpan(ctx context.Context, request *v1.CreateSpanRequest) (*v1.CreateSpanResponse, error) {
	parent := spans.Root
	predicate := func(span *v1.Span) bool {
		return span.SpanId == request.ParentSpanId
	}
	if request.ParentSpanId != "" {
		parent = spans.Root.Find(predicate)
		if parent == nil {
			parent = spans.Root
		}
	}
	parent.AddChild(&store.Node[*v1.Span]{
		Value: &v1.Span{
			TraceId:      request.TraceId,
			SpanId:       request.SpanId,
			ParentSpanId: request.ParentSpanId,
			Name:         request.Name,
			StartTime:    timestamppb.New(time.Now().UTC()),
		},
	})
	return &v1.CreateSpanResponse{
		TraceId:      request.TraceId,
		SpanId:       request.SpanId,
		ParentSpanId: request.ParentSpanId,
		Name:         request.Name,
		StartTime:    timestamppb.New(time.Now().UTC()),
	}, nil
}

func (t TraceServer) GetTrace(ctx context.Context, request *v1.GetTraceRequest) (*v1.GetTraceResponse, error) {
	parent := spans.Root
	predicate := func(span *v1.Span) bool {
		return span.TraceId == request.TraceId
	}
	found := parent.FindAll(predicate)
	spansFound := make([]*v1.Span, 0, len(found))
	for _, span := range found {
		spansFound = append(spansFound, span.Value)
	}
	return &v1.GetTraceResponse{
		TraceId: request.TraceId,
		Spans:   spansFound,
	}, nil
}

func (t TraceServer) GetAllTraces(ctx context.Context, empty *emptypb.Empty) (*v1.Traces, error) {
	getTraceId := func(span *v1.Span) string {
		return span.TraceId
	}
	allTraceIds := spans.Root.AllUniqueValue(getTraceId)
	return &v1.Traces{
		TraceIds: allTraceIds,
	}, nil
}
