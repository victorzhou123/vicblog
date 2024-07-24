package log

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelPanic = "PANIC"
	LevelFatal = "FATAL"
)

var runSugarLogger *zap.SugaredLogger

func Init(cfg *Config, exitSig chan struct{}) {
	// init run logger
	logger := zap.New(zapcore.NewCore(
		getEncoder(), newWriter(&cfg.RunWriter), logLevelMap(cfg.Level),
	))
	runSugarLogger = logger.Sugar()

	// switch sync flush on
	syncFlush(cfg.FlushTime, exitSig)
}

// SyncFlush: flush all buffers to local file within a fixed cycle(seconds)
func syncFlush(t int, exitSignal chan struct{}) {
	ticker := time.NewTicker(time.Duration(t) * time.Second)

	go func(t *time.Ticker) {
		for {
			select {
			case <-t.C:
				runSugarLogger.Sync() // #nosec G104
			case <-exitSignal:
				// sync log data before goroutine exit
				fmt.Println("receive exit signal")
				runSugarLogger.Sync() // #nosec G104
				return
			}

		}
	}(ticker)
}

func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = ""
	encodeConfig.LevelKey = ""

	return zapcore.NewConsoleEncoder(encodeConfig)
}

func newWriter(cfg *WriterConfig) zapcore.WriteSyncer {
	writer := &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}

	var ws io.Writer
	if cfg.StdPrint {
		ws = io.MultiWriter(writer, os.Stdout)
	} else {
		ws = io.MultiWriter(writer)
	}

	return zapcore.AddSync(ws)
}

// logLevelMap: mapping severity level to standard log level of zap
func logLevelMap(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelPanic:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

// toLog: convert log struct to string
func toLog(r any) string {
	var s strings.Builder

	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		s.WriteString(fmt.Sprintf("%s|", v.Field(i)))
	}

	return strings.TrimSuffix(s.String(), "|")
}

func writeLog(logger *zap.SugaredLogger, level string, info any) {
	switch level {
	case LevelDebug:
		logger.Debug(info)
	case LevelInfo:
		logger.Info(info)
	case LevelWarn:
		logger.Warn(info)
	case LevelError:
		logger.Error(info)
	case LevelPanic:
		logger.Panic(info)
	case LevelFatal:
		logger.Fatal(info)
	default:
		logger.Info(info)
	}
}
