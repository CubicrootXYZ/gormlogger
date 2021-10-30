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
func NewLogger(debug bool) *Logger {
	var err error
	var log *zap.Logger

	if debug {
		log, err = zap.NewDevelopment(zap.AddCallerSkip(1))
	} else {
		log, err = zap.NewProduction(zap.AddCallerSkip(1))

	}
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
	return logger
}

// Info log
func (logger *Logger) Info(ctx context.Context, msg string, args ...interface{}) {
	logger.logger.Debug(msg, args)
}

// Warn log
func (logger *Logger) Warn(ctx context.Context, msg string, args ...interface{}) {
	logger.logger.Warn(msg, args)
}

// Error log
func (logger *Logger) Error(ctx context.Context, msg string, args ...interface{}) {
	logger.logger.Error(msg, args)
}

// Trace log
func (logger *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
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
