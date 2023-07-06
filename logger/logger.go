package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// 	start : Create a zapcore log file from the generated log
	// if err = os.MkdirAll("./log", os.ModePerm); err != nil {
	// 	panic(err)
	// }
	// logFilePath := "./log/logger.log"
	// logFile, _ := os.Create(logFilePath)
	// defer logFile.Close()
	// config.OutputPaths = []string{logFilePath}
	// end : logfile

	if log, err = config.Build(zap.AddCallerSkip(1)); err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
