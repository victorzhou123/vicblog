package mysql

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var db *gorm.DB

func Init(cfg *Config) error {
	dbInst, err := gorm.Open(mysql.Open(cfg.toMasterDSN()), &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	// use master-slave replication
	dbInst.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(cfg.toMasterDSN())},
		Replicas: []gorm.Dialector{mysql.Open(cfg.toSlave01DSN()), mysql.Open(cfg.toSlave02DSN())},
		Policy:   dbresolver.RandomPolicy{},
	}))

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
