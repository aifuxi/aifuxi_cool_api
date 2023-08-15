package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger

func Setup() {

	ws := getLogWriter()
	e := getEncoder()

	core := zapcore.NewCore(e, ws, zap.DebugLevel)

	L = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")

	return zapcore.AddSync(file)
}
