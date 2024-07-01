package mysql

import (
	"context"
	"go.uber.org/fx"
	"go_sampler/providers/config"
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

func NewMysql(lc fx.Lifecycle, config *config.Config) *Database {
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: &CustomLogger{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			BaseLevel:                 dbLogger.Warn, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		},
	}
	if config.Debug {
		gormConfig.Logger = gormConfig.Logger.LogMode(dbLogger.Info)
	}

	var db = &Database{}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			gdb, err := gorm.Open(mysql.Open(config.DatabaseURL), gormConfig)
			if err != nil {
				logger.Errorf("fail to connect to database: %v", err)
				return err
			}

			db.DB = gdb
			return nil
		},
	})

	return db
}
