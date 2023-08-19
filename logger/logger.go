package logger

import (
	"os"

	"api.aifuxi.cool/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	var lg *zap.Logger

	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)
	lg = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	if settings.AppConfig.Mode == "debug" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   settings.LogConfig.Filename,
		MaxSize:    settings.LogConfig.MaxSize,
		MaxBackups: settings.LogConfig.MaxBackups,
		MaxAge:     settings.LogConfig.MaxAge,
	}

	if settings.AppConfig.Mode == "debug" {
		return zapcore.AddSync(os.Stdout)
	}

	return zapcore.AddSync(lumberjackLogger)
}
