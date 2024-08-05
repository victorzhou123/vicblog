package mysql

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *Config) error {
	dbInst, err := gorm.Open(mysql.Open(cfg.toDSN()), &gorm.Config{TranslateError: true})
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

func AutoMigrate(table interface{}) error {
	// pointer non-nil check
	if db == nil {
		err := errors.New("empty pointer of *gorm.DB")

		return err
	}

	return db.AutoMigrate(table)
}
