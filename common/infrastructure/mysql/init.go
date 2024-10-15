package mysql

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/victorzhou123/vicblog/common/log"
)

var db *gorm.DB

func Init(cfg *Config, logCfg *log.WriterConfig) error {
	dbInst, err := gorm.Open(mysql.Open(cfg.toDSN()), &gorm.Config{
		TranslateError: true,
		Logger:         log.NewSqlLog(logCfg),
	})
	if err != nil {
		return err
	}

	sqlDb, err := dbInst.DB()
	if err != nil {
		return err
	}

	sqlDb.SetConnMaxLifetime(cfg.maxLifTime())
	sqlDb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)

	db = dbInst

	return nil
}

func DB() *gorm.DB {
	return db
}

func AutoMigrate(tables ...any) error {
	// pointer non-nil check
	if db == nil {
		err := errors.New("empty pointer of *gorm.DB")

		return err
	}

	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			return err
		}
	}

	return nil
}
