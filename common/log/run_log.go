package log

import (
	"fmt"

	"victorzhou123/vicblog/common/constant"
	"victorzhou123/vicblog/common/util"
)

type runLog struct {
	Time            string
	Severity        string
	Host            string
	Service         string // optional
	TranslocationID string
	ThreadID        string // optional
	UserID          string // optional
	Position        string // optional
	Info            string
	AdditionalInfo  string // optional, format: json
}

func (r *runLog) setFromLogItem(item *runLogItem) {
	r.Severity = item.Level
	r.TranslocationID = item.TraceID
	r.Info = item.Info
}

func newDefaultRunLog() *runLog {
	return &runLog{
		Time:     util.TimeNowBaseSecond(),
		Severity: LevelInfo,
		Service:  constant.ServerName,
		Position: util.GetCallStackInfo(4),
	}
}

type runLogItem struct {
	Level   string
	TraceID string
	Info    string
}

func newRunLog(item *runLogItem) *runLog {
	log := newDefaultRunLog()
	log.setFromLogItem(item)

	return log
}

func writeRunLog(item *runLogItem) {
	runLog := newRunLog(item)

	writeLog(runSugarLogger, item.Level, toLog(*runLog))
}

func Debug(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelDebug,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Info(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelInfo,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Warn(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelWarn,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Error(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelError,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Panic(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelPanic,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Fatal(TraceID, template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelFatal,
		TraceID: TraceID,
		Info:    fmt.Sprintf(template, args...),
	})
}

func Debugf(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelDebug,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}

func Infof(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelInfo,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}

func Warnf(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelWarn,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}

func Errorf(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelError,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}

func Panicf(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelPanic,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}

func Fatalf(template string, args ...any) {
	writeRunLog(&runLogItem{
		Level:   LevelFatal,
		TraceID: "",
		Info:    fmt.Sprintf(template, args...),
	})
}
