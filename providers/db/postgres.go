package db

import (
	"go_sampler/providers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"log/slog"
	"time"
)

func NewPostgres(cfg *config.Config, logger *slog.Logger) (*Database, error) {
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

	var err error
	var db = &Database{Config: cfg}
	db.DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), gormConfig)
	if err != nil {
		slog.Error("fail to connect to database", "error", err)
		return nil, err
	}

	DB = db
	return db, nil
}
