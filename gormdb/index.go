package gormdb

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
)

type DbConfig struct {
	Url          string
	MaxIdleConns *int
	MaxOpenConns *int
	MaxLife      *time.Duration
}

func Connect(dbconfig DbConfig, otp ...gorm.Option) (*gorm.DB, error) {
	if instance == nil {
		var err error

		instance, err = gorm.Open(postgres.Open(dbconfig.Url), otp...)
		if err != nil {
			return nil, err
		}
		sqlDB, err := instance.DB()
		if err != nil {
			return nil, err
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		var maxIdle = 10
		var maxOpen = 100
		var maxLife = time.Minute
		if dbconfig.MaxIdleConns != nil {
			maxIdle = *dbconfig.MaxIdleConns
		}
		if dbconfig.MaxOpenConns != nil {
			maxOpen = *dbconfig.MaxOpenConns
		}
		if dbconfig.MaxLife != nil {
			maxLife = *dbconfig.MaxLife
		}
		sqlDB.SetMaxIdleConns(maxIdle)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(maxOpen)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(maxLife)
		return instance, nil
	} else {
		return instance, nil
	}
}
