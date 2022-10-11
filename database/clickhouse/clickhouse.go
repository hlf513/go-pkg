package clickhouse

import (
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var dbs sync.Map

func Connect(opts ...Option) error {
	options := newOptions(opts...)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=%d&write_timeout=%d",
		options.Host,
		options.Port,
		options.Database,
		options.Username,
		options.Password,
		options.ReadTimeout,
		options.WriteTimeout,
	)
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
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{
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

	log.Fatal("clickhouse instance[" + key + "] not found")

	return nil
}
