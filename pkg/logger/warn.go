package logger

import (
	errs "github.com/EliabeBastosDias/cinema-api/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (l *Logger) Warn(message string, err error) {
	if httpErr, ok := err.(*errs.HTTPError); ok {
		l.logWarn(httpErr, message)
	} else {
		regularErr := errs.E(err, errs.UnexpectedError)
		l.logWarn(regularErr.(*errs.HTTPError), message)
	}
}

func (l *Logger) logWarn(e *errs.HTTPError, message string) {
	errData := logStruct{
		Error:   e.Error(),
		Stack:   e.StackTrace(),
		Context: e.JSONParams(),
	}

	zapFields := []zapcore.Field{
		zap.Any("error", errData),
	}

	l.logger.Warn(message, zapFields...)
}
