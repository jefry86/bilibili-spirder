//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/16 8:30 下午

package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func NewLogger() Logger {
	cfg := loggerCfg()
	logger, err := cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))
	if err != nil {
		panic(err)
	}
	return Logger{
		ZapLog: logger,
	}
}

func loggerCfg() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig(),
		OutputPaths:      []string{"stderr", fmt.Sprintf("%s%s%s", Config.Path.Logs, string(os.PathSeparator), "info.log")},
		ErrorOutputPaths: []string{"stderr", fmt.Sprintf("%s%s%s", Config.Path.Logs, string(os.PathSeparator), "error.log")},
	}
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

type Logger struct {
	ZapLog *zap.Logger
}

func (l *Logger) Info(msg ...string) {
	l.ZapLog.Info(strings.Join(msg, " "))
	defer l.ZapLog.Sync()
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.ZapLog.Info(fmt.Sprintf(format, a))
	defer l.ZapLog.Sync()
}

func (l *Logger) Warn(msg ...string) {
	l.ZapLog.Warn(strings.Join(msg, " "))
	defer l.ZapLog.Sync()
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.ZapLog.Warn(fmt.Sprintf(format, a))
	defer l.ZapLog.Sync()
}

func (l *Logger) Error(msg ...string) {
	l.ZapLog.Error(strings.Join(msg, " "))
	defer l.ZapLog.Sync()
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.ZapLog.Error(fmt.Sprintf(format, a))
	defer l.ZapLog.Sync()
}
