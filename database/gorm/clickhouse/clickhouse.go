package mysql

import (
	"fmt"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Connect(opts ...Option) (*gorm.DB, error) {
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
	// log level https://github.com/go-gorm/gorm/issues/3544
	//database.Logger = logger.Default.LogMode(logger.Silent)
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// set connection pool
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(options.MaxIdleConn)
	sqlDB.SetMaxOpenConns(options.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(options.MaxLifeTime)

	return db, nil
}
