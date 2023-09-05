package logs

import (
	"github.com/sirupsen/logrus"
	v1 "monitoring/proto/gen/go/monitoring/v1"
)

func convertToLogrusLevel(level v1.ProtoLogLevel) logrus.Level {
	switch level {
	case v1.ProtoLogLevel_DEBUG:
		return logrus.DebugLevel
	case v1.ProtoLogLevel_INFO:
		return logrus.InfoLevel
	case v1.ProtoLogLevel_WARN:
		return logrus.WarnLevel
	case v1.ProtoLogLevel_ERROR:
		return logrus.ErrorLevel
	case v1.ProtoLogLevel_FATAL:
		return logrus.FatalLevel
	case v1.ProtoLogLevel_PANIC:
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel // or some other default
	}
}

func convertToProtoLevel(level logrus.Level) v1.ProtoLogLevel {
	switch level {
	case logrus.DebugLevel:
		return v1.ProtoLogLevel_DEBUG
	case logrus.InfoLevel:
		return v1.ProtoLogLevel_INFO
	case logrus.WarnLevel:
		return v1.ProtoLogLevel_WARN
	case logrus.ErrorLevel:
		return v1.ProtoLogLevel_ERROR
	case logrus.FatalLevel:
		return v1.ProtoLogLevel_FATAL
	case logrus.PanicLevel:
		return v1.ProtoLogLevel_PANIC
	default:
		return v1.ProtoLogLevel_INFO // or some other default
	}
}
