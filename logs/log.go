package logs

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	formattedPath := fmt.Sprintf("logs/logs_data/%d-%s-%d.log", time.Now().Day(), time.Now().Month().String(), time.Now().Year())
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	config.OutputPaths = []string{formattedPath, "stdout"}
	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, field ...zap.Field) {
	log.Info(message, field...)

}

func Debug(message string, field ...zap.Field) {
	log.Debug(message, field...)
}

func Error(message interface{}, field ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), field...)
	case string:
		log.Error(v, field...)
	}
}
