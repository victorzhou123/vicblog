package log

import (
	"bytes"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	buf = &bytes.Buffer{}
)

func init() {
	// log init
	logger := zap.New(zapcore.NewCore(
		getEncoder(), newSyncWriter(buf), logLevelMap(levelDebug),
	))
	runSugarLogger = logger.Sugar()
}

func TestInfo(t *testing.T) {
	type args struct {
		TraceID  string
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				TraceID:  "test_trace_id_01",
				template: "%s ate an apple",
				args:     []any{"victor"},
			},
			want: "|INFO||vicBlog|test_trace_id_01|||run_log_test.go:47|victor ate an apple|\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.TraceID, tt.args.template, tt.args.args...)

			index := strings.IndexRune(buf.String(), '|')
			if index == -1 {
				t.Errorf("assert failed")
				return
			}
			got := buf.String()[index:]

			if got != tt.want {
				t.Errorf("got %s want %s", got, tt.want)
			}
		})
	}
}
