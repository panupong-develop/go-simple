package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Environment int

const (
	Development Environment = 1 << iota
	Production  Environment = 1 << iota
)

func NewZapLogger(e Environment, filePath string) *zap.Logger {
	var config zapcore.EncoderConfig
	var defaultLogLevel zapcore.Level

	if e == Development {
		config = zap.NewDevelopmentEncoderConfig()
		defaultLogLevel = zapcore.DebugLevel
	} else {
		config = zap.NewProductionEncoderConfig()
		defaultLogLevel = zapcore.InfoLevel
	}

	// configurations
	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.LevelKey = "level"
	config.MessageKey = "message"
	config.NameKey = "name"
	config.StacktraceKey = "stacktrace"
	config.TimeKey = "time"
	config.CallerKey = "caller"
	config.FunctionKey = "func"

	// create file encoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)

	// create console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	// creates a Core that duplicates log entries into two or more underlying Cores.
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
