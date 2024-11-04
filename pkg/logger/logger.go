package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logStruct struct {
	Stack   string `json:"stacktrace,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Context string `json:"context"`
}

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Close() error {
	return l.logger.Sync()
}

func New() *Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		LineEnding:  zapcore.DefaultLineEnding,
	}

	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger := &Logger{
		logger: zapLogger,
	}

	return logger
}

type Provider interface {
	Warn(message string, err error)
	DPanic(recoveryData any, msg ...string)
	Panic(auditUuid string, recoveryData any)
	Error(message string, err error)
	Info(message string, ctx ...interface{})
}
