package gormwrapper

import (
	"context"
	"errors"
	"fmt"
	"github.com/bdcp-ops/alpha/alog"
	"github.com/bdcp-ops/alpha/autil/ahttp/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type LoggerWrapper struct {
	*logger
	PrintRecordNotFoundError bool
}

func NewLoggerWrapper(sugarLogger *zap.SugaredLogger, config Config, printRecordNotFoundError bool) gormlogger.Interface {
	lw := &LoggerWrapper{
		New(sugarLogger, config).(*logger),
		printRecordNotFoundError,
	}
	return lw
}

// Trace print sql message
func (lw *LoggerWrapper) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l := lw
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || lw.PrintRecordNotFoundError):
			sql, rows := fc()
			if rows == -1 {
				l.sugarLogger.Errorw(
					fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql),
					alog.RequestIdKey,
					request.RequestIdValue(ctx))
			} else {
				l.sugarLogger.Errorw(
					fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql),
					alog.RequestIdKey,
					request.RequestIdValue(ctx))
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn, l.LogLevel >= gormlogger.Info:
			l.logger.Trace(ctx, begin, fc, err)
		}
	}
}
