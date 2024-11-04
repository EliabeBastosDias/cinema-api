package logger

import (
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (l *Logger) Info(message string, ctx ...interface{}) {
	switch len(ctx) {
	case 0:
		l.logger.Info(message)
	case 1:
		context := ctx[0]

		jsonContext, err := json.Marshal(context)
		if err != nil {
			panic(err)
		}

		logData := logStruct{
			Context: string(jsonContext),
		}

		zapFields := composeLogFields(logData, context)

		l.logger.Info(message, zapFields...)
	default:
		panic("too many arguments to logger.Info")
	}
}

func composeLogFields(logData logStruct, context interface{}) []zapcore.Field {
	var fields []zapcore.Field
	fields = append(fields, zap.Any("info", logData))

	// todo: check for holders once middlewares are implemented(EchoToken, Timespan, etc...) and add their fields
	// if echoTokenHolder, ok := context.(echoTokenHolder); ok {
	// 	fields = append(fields, zap.String("echoToken", echoTokenHolder.GetEchoToken()))
	// }

	return fields
}
