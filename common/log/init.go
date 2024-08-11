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
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
	levelPanic = "PANIC"
	levelFatal = "FATAL"
)

var runSugarLogger *zap.SugaredLogger

func Init(cfg *Config, exitSig chan struct{}) {
	// init run logger
	logger := zap.New(zapcore.NewCore(
		getEncoder(), newSyncWriterWithConfig(&cfg.RunWriter), logLevelMap(cfg.Level),
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
				//#nosec G104
				runSugarLogger.Sync()
			case <-exitSignal:
				// sync log data before goroutine exit
				fmt.Println("receive exit signal")
				runSugarLogger.Sync() //#nosec G104
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

func newSyncWriterWithConfig(cfg *WriterConfig) zapcore.WriteSyncer {
	if cfg.StdPrint {
		return newSyncWriter(
			newLumberjackLogger(cfg),
			os.Stdout,
		)
	}

	return newSyncWriter(newLumberjackLogger(cfg))
}

func newSyncWriter(writers ...io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(io.MultiWriter(writers...))
}

func newLumberjackLogger(cfg *WriterConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
}

// logLevelMap: mapping severity level to standard log level of zap
func logLevelMap(level string) zapcore.Level {
	switch level {
	case levelDebug:
		return zapcore.DebugLevel
	case levelInfo:
		return zapcore.InfoLevel
	case levelWarn:
		return zapcore.WarnLevel
	case levelPanic:
		return zapcore.ErrorLevel
	case levelFatal:
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
	if logger == nil {
		return
	}

	switch level {
	case levelDebug:
		logger.Debug(info)
	case levelInfo:
		logger.Info(info)
	case levelWarn:
		logger.Warn(info)
	case levelError:
		logger.Error(info)
	case levelPanic:
		logger.Panic(info)
	case levelFatal:
		logger.Fatal(info)
	default:
		logger.Info(info)
	}
}
