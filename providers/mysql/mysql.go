package mysql

import (
	"context"
	"log/slog"
	"time"

	"go_sampler/providers/config"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *Database

func NewMysql(lc fx.Lifecycle, cfg *config.Config, logger *slog.Logger) *Database {
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: &CustomLogger{
			logger:                    logger,
			SlowThreshold:             time.Second,   // Slow SQL threshold
			BaseLevel:                 dbLogger.Warn, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		},
	}
	if cfg.Debug {
		gormConfig.Logger = gormConfig.Logger.LogMode(dbLogger.Info)
	}

	var db = &Database{}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			gdb, err := gorm.Open(mysql.Open(cfg.DatabaseURL), gormConfig)
			if err != nil {
				slog.Error("fail to connect to database", "error", err)
				return err
			}

			db.DB = gdb
			return nil
		},
	})

	return db
}
