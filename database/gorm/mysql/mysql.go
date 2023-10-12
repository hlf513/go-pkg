package mysql

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
	"sync"
	"time"
)

var dbs sync.Map

func Connect(opts ...Option) error {
	options := NewOptions(opts...)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.Database,
	)
	// log level https://github.com/go-gorm/gorm/issues/3544
	//database.Logger = logger.Default.LogMode(logger.Silent)
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return err
	}
	// set connection pool
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(options.MaxIdleConn)
	sqlDB.SetMaxOpenConns(options.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(options.MaxLifeTime) * time.Second)

	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		return err
	}

	dbs.Store(options.Name, db)

	return nil
}

func GetDB(ctx context.Context, name ...string) *gorm.DB {
	var key = "default"
	if len(name) > 0 {
		key = name[0]
	}

	if db, ok := dbs.Load(key); ok {
		return db.(*gorm.DB).WithContext(ctx)
	}

	log.Fatal("mysql pool[" + key + "] not found")

	return nil
}
