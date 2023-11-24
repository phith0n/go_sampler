package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger zap.Logger
var sugar zap.SugaredLogger

func GetLogger() *zap.Logger {
	return &logger
}

func GetSugar() *zap.SugaredLogger {
	return &sugar
}

func InitLogger(debug bool) error {
	if t, err := NewZapLogger(debug); err != nil {
		return nil
	} else {
		logger = *t
		sugar = *logger.Sugar()
	}
	return nil
}

func NewZapLogger(debug bool) (*zap.Logger, error) {
	var config = zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "time",
			LevelKey:         "level",
			NameKey:          "name",
			CallerKey:        "caller",
			FunctionKey:      "function",
			MessageKey:       "message",
			StacktraceKey:    zapcore.OmitKey,
			ConsoleSeparator: "|",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339Nano))
			},
			EncodeName:     zapcore.FullNameEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	if debug {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.ConsoleSeparator = " "
		config.Encoding = "console"
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Development = true
	}

	l, err := config.Build()
	if err != nil {
		return nil, err
	} else {
		return l, nil
	}
}
