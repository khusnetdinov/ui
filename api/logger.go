package api

import (
	"io"
	"log/slog"
)

type Slog struct {
	debug  bool
	logger *slog.Logger
}

func NewSlog(writer io.Writer, debug bool) *Slog {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: level,
	}))

	return &Slog{
		debug: debug,
		logger: logger,
	}
}

func (slog Slog) Debug(msg string, args ...interface{}) {
	slog.logger.Debug(msg, args...)
}

func (slog Slog) Info(msg string, args ...interface{}) {
	slog.logger.Info(msg, args...)
}

func (slog Slog) Warn(msg string, args ...interface{}) {
	slog.logger.Warn(msg, args...)
}

func (slog Slog) Error(msg string, args ...interface{}) {
	slog.logger.Error(msg, args...)
}
