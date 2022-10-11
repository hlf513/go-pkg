package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var dbs sync.Map

func Connect(opts ...Option) error {
	options := newOptions(opts...)
	var logLevel logger.LogLevel
	switch options.LogLevel {
	case Silent:
		logLevel = logger.Silent
	case Error:
		logLevel = logger.Error
	case Warn:
		logLevel = logger.Warn
	case Info:
		logLevel = logger.Info
	}
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html#SQLite
	db, err := gorm.Open(sqlite.Open(options.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return err
	}
	// set connection pool
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(options.MaxIdleConn)
	sqlDB.SetMaxOpenConns(options.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(options.MaxLifeTime)

	dbs.Store(options.Name, db)

	return nil
}

func GetDB(name ...string) *gorm.DB {
	var key = "default"
	if len(name) > 0 {
		key = name[0]
	}

	if db, ok := dbs.Load(key); ok {
		return db.(*gorm.DB)
	}

	log.Fatal("sqlite instance[" + key + "] not found")

	return nil
}