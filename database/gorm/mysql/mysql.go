package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(opts ...Option) (*gorm.DB, error) {
	options := newOptions(opts...)
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
	db, err := gorm.Open(mysql.Open(dsn), &options.GormConfig)
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
