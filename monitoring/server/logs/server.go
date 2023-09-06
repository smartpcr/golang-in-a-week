package logs

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/smartpcr/go-otel/pkg/ot"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "monitoring/proto/gen/go/monitoring/v1"
	"monitoring/server/store"
	"time"
)

type LogServer struct {
	v1.UnimplementedLogServiceServer
	Logger *logrus.Entry
}

type Log struct {
	Message   *v1.LogMessage
	Timestamp time.Time
}

var _ v1.LogServiceServer = &LogServer{}

var (
	logs *store.ConcurrentQueue[Log]
)

func init() {
	logs = store.NewConcurrentQueue[Log]()
}

func (l LogServer) Log(ctx context.Context, request *v1.LogMessage) (*v1.LogMessageResponse, error) {
	ctx, span, logger := ot.StartSpanLogger(ctx)
	defer span.End()
	logger.Infof(
		"new log message: level=%s, from %s.%s:%d\n%s",
		request.Level.String(),
		request.FileName,
		request.FunctionName,
		request.LineNumber,
		request.Message,
	)
	span.AddEvent(request.FunctionName)

	logs.Enqueue(Log{
		Message: &v1.LogMessage{
			Level:      request.Level,
			Message:    request.Message,
			LineNumber: request.LineNumber,
			FileName:   request.FileName,
		},
		Timestamp: time.Now().UTC(),
	})
	return &v1.LogMessageResponse{
		Message: &v1.LogMessage{
			Level:      request.Level,
			Message:    request.Message,
			LineNumber: request.LineNumber,
			FileName:   request.FileName,
		},
		Timestamp: timestamppb.New(time.Now().UTC()),
	}, nil
}

func (l LogServer) GetLogs(ctx context.Context, empty *emptypb.Empty) (*v1.LogMessages, error) {
	logMessages := make([]*v1.LogMessage, 0, logs.Len())
	for {
		msg, ok := logs.Dequeue()
		if !ok {
			break
		}
		logMessages = append(logMessages, &v1.LogMessage{
			Level:      msg.Message.Level,
			Message:    msg.Message.Message,
			LineNumber: msg.Message.LineNumber,
			FileName:   msg.Message.FileName,
		})
	}
	return &v1.LogMessages{
		Messages: logMessages,
	}, nil
}
