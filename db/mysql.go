package db

import (
	"time"

	"go_sampler/logging"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var logger = logging.GetSugar()

type Database struct {
	*gorm.DB
}

var DB *Database

func InitMysql(databaseURL string, debug bool) error {
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: &CustomLogger{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			BaseLevel:                 dbLogger.Warn, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		},
	}
	if debug {
		gormConfig.Logger = gormConfig.Logger.LogMode(dbLogger.Info)
	}

	db, err := gorm.Open(mysql.Open(databaseURL), gormConfig)
	if err != nil {
		logger.Errorf("fail to connect to database: %v", err)
		return err
	}

	DB = &Database{
		DB: db,
	}
	return nil
}
