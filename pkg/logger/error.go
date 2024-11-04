package logger

import (
	errs "github.com/EliabeBastosDias/cinema-api/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (l *Logger) Error(message string, err error) {
	if httpErr, ok := err.(*errs.HTTPError); ok {
		l.logError(httpErr, message)
	} else {
		regularErr := errs.E(err, errs.UnexpectedError)
		l.logError(regularErr.(*errs.HTTPError), message)
	}
}

func (l *Logger) logError(e *errs.HTTPError, message string) {
	errData := logStruct{
		Error:   e.Error(),
		Stack:   e.StackTrace(),
		Context: e.JSONParams(),
	}

	zapFields := []zapcore.Field{
		zap.Any("error", errData),
		zap.String("echoToken", e.EchoToken()),
	}

	l.logger.Error(message, zapFields...)
}
