package dlog

import (
	"context"
	"io"

	"github.com/inconshreveable/log15"
)

type Logger interface {
	log15.Logger
	io.Closer
}

type Log interface {
	LoggerFunc(namespace, instance string) (Logger, error)
	NamespaceLogger(namespace string) (Logger, error)
	QueryLogs(ctx context.Context, instance string, limit, offset int) (QueryReponse, error)
	StreamLogs(ctx context.Context, instance string) (chan interface{}, error)
	DeleteNamespaceLogs(namespace string) error
	DeleteInstanceLogs(instance string) error
}

type LogEntry struct {
	Level     string            `json:"lvl"`
	Timestamp int64             `json:"time"`
	Message   string            `json:"msg"`
	Context   map[string]string `json:"ctx"`
}

type QueryReponse struct {
	Count  int        `json:"count"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset"`
	Logs   []LogEntry `json:"data"`
}
