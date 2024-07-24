package log

type Config struct {
	Level     string
	FlushTime int // unit second

	InterfaceWriter WriterConfig
	RunWriter       WriterConfig
}

func (cfg *Config) SetDefault() {
	cfg.Level = LevelInfo
	cfg.FlushTime = 5

	cfg.InterfaceWriter.setDefault()
	cfg.RunWriter.setDefault()
}

type WriterConfig struct {
	FilePath   string
	MaxSize    int // unit MB
	MaxAge     int // unit day
	MaxBackups int // max backup logs
	LocalTime  bool
	Compress   bool // compress historical log
	StdPrint   bool // if print to os.Stdout
}

func (w *WriterConfig) setDefault() {
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
