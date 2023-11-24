package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type CustomLogger struct {
	BaseLevel                 dbLogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func (l *CustomLogger) LogMode(level dbLogger.LogLevel) dbLogger.Interface {
	newLogger := *l
	newLogger.BaseLevel = level
	return &newLogger
}

func (l *CustomLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.BaseLevel >= dbLogger.Info {
		logger.Infof(s, i...)
	}
}

func (l *CustomLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.BaseLevel >= dbLogger.Warn {
		logger.Warnf(s, i...)
	}
}

func (l *CustomLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.BaseLevel >= dbLogger.Error {
		logger.Errorf(s, i...)
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.BaseLevel <= dbLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.BaseLevel >= dbLogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Error(ctx, "%s %v [%.3fms] %s", utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			l.Error(ctx, "%s %v [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.BaseLevel >= dbLogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Warn(ctx, "%s %s [%.3fms] %s", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			l.Warn(ctx, "%s %s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.BaseLevel == dbLogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Info(ctx, "%s [%.3fms] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, sql)
		} else {
			l.Info(ctx, "%s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
