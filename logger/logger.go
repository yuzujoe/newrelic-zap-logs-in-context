package logger

import "go.uber.org/zap"

var (
	Logger *zap.Logger
	err    error
)

func NewLogger() (*zap.Logger, error) {
	Logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return Logger, nil
}
