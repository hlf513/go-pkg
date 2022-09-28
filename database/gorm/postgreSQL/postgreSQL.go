package mysql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(opts ...Option) (*gorm.DB, error) {
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
	// log level https://github.com/go-gorm/gorm/issues/3544
	//database.Logger = logger.Default.LogMode(logger.Silent)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
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
