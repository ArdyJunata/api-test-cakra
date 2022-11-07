package logger

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Infof(ctx context.Context, format string, msg ...interface{})
	Errorf(ctx context.Context, format string, msg ...interface{})
	build(ctx context.Context) *logrus.Entry
}

type log struct {
	logger *logrus.Logger
}

func NewLog() Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	return &log{
		logger: logger,
	}
}

func (l *log) build(ctx context.Context) *logrus.Entry {

	var fields logrus.Fields = generateFields(ctx)
	entry := l.logger.WithFields(fields)
	return entry
}

func (l *log) Infof(ctx context.Context, format string, msg ...interface{}) {
	l.build(ctx).Infof(format, msg...)
}

func (l *log) Errorf(ctx context.Context, format string, msg ...interface{}) {
	l.build(ctx).Errorf(format, msg...)
}

func generateFields(ctx context.Context) map[string]interface{} {
	val := ctx.Value(DATA)
	var data map[string]interface{} = map[string]interface{}{}

	if val == nil {
		return data
	}
	for k, v := range val.(map[LogKey]interface{}) {
		data[strings.ToLower(string(k))] = v
	}

	return data
}
