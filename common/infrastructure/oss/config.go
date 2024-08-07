package oss

import "errors"

type Config struct {
	Id               string `json:"id"`
	Secret           string `json:"secret"`
	Endpoint         string `json:"endpoint"`
	Bucket           string `json:"bucket"`
	ConnTimeout      int64  `json:"conn_timeout"`       // unit is seconds
	ReadWriteTimeout int64  `json:"read_write_timeout"` // unit is seconds

	RootDir           string `json:"root_dir"` // root dir of file storage
	PictureFolderName string `json:"picture_folder_name"`
}

func (cfg *Config) Validate() error {
	if cfg.Id == "" {
		return errors.New("id can not be empty")
	}

	if cfg.Secret == "" {
		return errors.New("secret can not be empty")
	}

	if cfg.Endpoint == "" {
		return errors.New("endpoint can not be empty")
	}

	if cfg.Bucket == "" {
		return errors.New("bucket can not be empty")
	}

	if cfg.RootDir == "" {
		return errors.New("RootDir can not be empty")
	}

	if cfg.PictureFolderName == "" {
		return errors.New("PictureFolderName can not be empty")
	}

	return nil
}

func (cfg *Config) SetDefault() {
	if cfg.ConnTimeout == 0 {
		cfg.ConnTimeout = 10
	}

	if cfg.ReadWriteTimeout == 0 {
		cfg.ReadWriteTimeout = 20
	}
}
