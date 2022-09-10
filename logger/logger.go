package logger

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"

	lic "github.com/newrelic/go-agent/v3/integrations/logcontext"
)

var Lgr *Client

type Client struct {
	Logger
}

type Logger interface {
	Debug(msg string, txn *newrelic.Transaction)
	Info(msg string, txn *newrelic.Transaction)
	Warn(msg string, txn *newrelic.Transaction)
	Error(msg string, txn *newrelic.Transaction)
}

type logger struct {
	Logger *zap.Logger
}

func NewLogger() (*Client, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(log)

	Lgr = &Client{
		Logger: newZapLogger(log),
	}

	return Lgr, nil
}

func newZapLogger(z *zap.Logger) *logger {
	return &logger{
		Logger: z,
	}
}

func (l logger) Debug(msg string, txn *newrelic.Transaction) {
	l.Logger.Debug(msg, transformAttributes(txn)...)
}

func (l *logger) Info(msg string, txn *newrelic.Transaction) {
	l.Logger.Info(msg, transformAttributes(txn)...)
}

func (l *logger) Warn(msg string, txn *newrelic.Transaction) {
	l.Logger.Warn(msg, transformAttributes(txn)...)
}

func (l *logger) Error(msg string, txn *newrelic.Transaction) {
	l.Logger.Error(msg, transformAttributes(txn)...)
}

func transformAttributes(txn *newrelic.Transaction) []zap.Field {
	if txn == nil {
		return []zap.Field{}
	}
	md := txn.GetLinkingMetadata()

	return []zap.Field{
		zap.String(lic.KeyTraceID, md.TraceID),
		zap.String(lic.KeySpanID, md.SpanID),
		zap.String(lic.KeyEntityName, md.EntityName),
		zap.String(lic.KeyEntityType, md.EntityType),
		zap.String(lic.KeyEntityGUID, md.EntityGUID),
		zap.String(lic.KeyHostname, md.Hostname),
	}
}
