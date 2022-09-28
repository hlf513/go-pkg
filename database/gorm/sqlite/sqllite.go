package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(opts ...Option) (*gorm.DB, error) {
	options := newOptions(opts...)
	// log level https://github.com/go-gorm/gorm/issues/3544
	//database.Logger = logger.Default.LogMode(logger.Silent)
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html#SQLite
	db, err := gorm.Open(sqlite.Open(options.DSN), &gorm.Config{})
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
