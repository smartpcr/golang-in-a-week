package logs

import (
	"context"
	"github.com/sirupsen/logrus"
	v1 "monitoring/proto/gen/go/monitoring/v1"
)

type LogServer struct {
	v1.UnimplementedLogServiceServer
	Logger *logrus.Entry
}

func (l LogServer) Log(ctx context.Context, request *v1.LogMessageRequest) (*v1.LogMessageResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ v1.LogServiceServer = &LogServer{}
