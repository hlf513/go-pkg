package postgresSQL

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var dbs sync.Map

func Connect(opts ...Option) error {
	options := newOptions(opts...)
	// refer https://gorm.io/zh_CN/docs/connecting_to_the_database.html#PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		options.Host,
		options.Username,
		options.Password,
		options.Database,
		options.Port,
		options.TimeZone,
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
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
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

	log.Fatal("postgreSQL instance[" + key + "] not found")

	return nil
}