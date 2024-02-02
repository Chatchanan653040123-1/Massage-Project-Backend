package logs

import (
	"fmt"
	"time"

	"github.com/gtuk/discordwebhook"
	"github.com/spf13/viper"
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
	//DiscordBot(message, "Info")
	log.Info(message, field...)

}

func Debug(message string, field ...zap.Field) {
	//DiscordBot(message, "Debug")
	log.Debug(message, field...)
}

func Error(message interface{}, field ...zap.Field) {
	switch v := message.(type) {
	case error:
		//DiscordBot(v.Error(), "Error")
		log.Error(v.Error(), field...)
	case string:
		//DiscordBot(v, "Error")
		log.Error(v, field...)
	}
}
func DiscordBot(messageOfLog string, messageType string) {
	var url = viper.GetString("DISCORD_WEBHOOK_URL")
	message := discordwebhook.Message{
		Username: &messageType,
		Content:  &messageOfLog,
	}
	discordwebhook.SendMessage(url, message)
}
