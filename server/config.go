package server

type Config struct {
	Port              int `json:"port"`
	ReadTimeout       int `json:"read_timeout"`        // unit Millisecond
	ReadHeaderTimeout int `json:"read_header_timeout"` // unit Millisecond
}

func (cfg *Config) SetDefault() {
	if cfg.Port == 0 {
		cfg.Port = 8080
	}

	if cfg.ReadTimeout == 0 {
		cfg.ReadHeaderTimeout = 10000
	}

	if cfg.ReadHeaderTimeout == 0 {
		cfg.ReadHeaderTimeout = 2000
	}
}
