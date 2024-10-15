package log

import (
	"log"
	"time"

	"gorm.io/gorm/logger"
)

func NewSqlLog(cfg *WriterConfig) logger.Interface {
	return logger.New(
		log.New(newLumberjackLogger(cfg), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                  // Don't include params in the SQL log
			Colorful:                  false,                  // Disable color
		},
	)
}
