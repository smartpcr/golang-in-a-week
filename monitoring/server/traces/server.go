package traces

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "monitoring/proto/gen/go/monitoring/v1"
)

type TraceServer struct {
	v1.UnimplementedTracesServiceServer
	Tracer
}

func (t TraceServer) CreateTrace(ctx context.Context, request *v1.CreateTraceRequest) (*v1.CreateTraceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TraceServer) GetTrace(ctx context.Context, request *v1.GetTraceRequest) (*v1.GetTraceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TraceServer) QueryTraces(ctx context.Context, request *v1.QueryTraceRequest) (*v1.QueryTraceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TraceServer) GetAllTraces(ctx context.Context, empty *emptypb.Empty) (*v1.Trace, error) {
	//TODO implement me
	panic("implement me")
}

var _ v1.TracesServiceServer = &TraceServer{}
