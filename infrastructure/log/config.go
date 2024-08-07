package log

type Config struct {
	Level     string `json:"level"`
	FlushTime int    `json:"flush_time"` // unit second

	InterfaceWriter WriterConfig `json:"interface_writer"`
	RunWriter       WriterConfig `json:"run_writer"`
}

func (cfg *Config) SetDefault() {
	cfg.Level = levelInfo
	cfg.FlushTime = 5

	cfg.InterfaceWriter.setDefault()
	cfg.RunWriter.setDefault()
}

type WriterConfig struct {
	FilePath   string `json:"file_path"`
	MaxSize    int    `json:"max_size"`    // unit MB
	MaxAge     int    `json:"max_age"`     // unit day
	MaxBackups int    `json:"max_backups"` // max backup logs
	LocalTime  bool   `json:"local_time"`
	Compress   bool   `json:"compress"`  // compress historical log
	StdPrint   bool   `json:"std_print"` // if print to os.Stdout
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
