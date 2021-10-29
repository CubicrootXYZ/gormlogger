package gormlogger

import (
	"context"
	"time"

	"go.uber.org/zap"
	gormlog "gorm.io/gorm/logger"
)

// Logger gormlogger
type Logger struct {
	logger   *zap.SugaredLogger
	logLevel gormlog.LogLevel
}

// NewLogger makes a new logger instance
func NewLogger() *Logger {
	log, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	l := log.Sugar()

	return &Logger{
		logger:   l,
		logLevel: gormlog.Info,
	}
}

// Sync call defer logger.Sync() to empty the buffer
func (logger *Logger) Sync() {
	logger.logger.Sync()
}

// LogMode sets the log level
func (logger *Logger) LogMode(level gormlog.LogLevel) gormlog.Interface {
	logger.logLevel = level
	return logger
}

// Info log
func (logger *Logger) Info(ctx context.Context, msg string, args ...interface{}) {
	if logger.logLevel < gormlog.Info {
		return
	}
	logger.logger.Info(msg, args)
}

// Warn log
func (logger *Logger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if logger.logLevel < gormlog.Warn {
		return
	}
	logger.logger.Warn(msg, args)
}

// Error log
func (logger *Logger) Error(ctx context.Context, msg string, args ...interface{}) {
	if logger.logLevel < gormlog.Error {
		return
	}
	logger.logger.Error(msg, args)
}

// Trace log
func (logger *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if logger.logLevel < gormlog.Info {
		return
	}
	var sql string
	var affected int64
	var errMsg string

	if fc != nil {
		sql, affected = fc()
	}

	if err != nil {
		errMsg = err.Error()
	}

	logger.logger.Debugw(errMsg,
		"begin", begin.UTC(),
		"sql", sql,
		"rows_affected", affected)
}
