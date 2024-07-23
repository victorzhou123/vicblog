package log

type Config struct {
	Level string

	InterfaceWriter Writer
	RunWriter       Writer
}

func (cfg *Config) SetDefault() {
	cfg.Level = LevelInfo

	cfg.InterfaceWriter.setDefault()
	cfg.RunWriter.setDefault()
}

type Writer struct {
	FilePath   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

func (w *Writer) setDefault() {
	if w.FilePath == "" {
		w.FilePath = "./log/std.log"
	}

	if w.MaxSize == 0 {
		w.MaxSize = 100
	}

	if w.MaxAge == 0 {
		w.MaxAge = 7
	}

	if w.MaxBackups == 0 {
		w.MaxBackups = 15
	}
}
