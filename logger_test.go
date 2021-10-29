package gormlogger

import (
	"context"
	"testing"

	gormlog "gorm.io/gorm/logger"
)

func Test1(t *testing.T) {
	logger := NewLogger()
	logger.LogMode(gormlog.Warn)
	logger.Info(context.Background(), "gdfg")
}
