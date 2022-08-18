package plog

import (
	"fmt"

	"github.com/v2pro/plz/gls"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pan-jf/go-utils/plog/mode"
	"github.com/pan-jf/go-utils/plog/util"
)

var (
	zapLogger *zap.Logger
)

// 支持直接启动
func init() {
	_ = initZapLog()
}

// initZapLog 根据options的设置,初始化日志系统。
// 注意默认是测试环境模式,需要设置线上模式的需要设置TestEnv(false)
func initZapLog(opts ...util.Option) error {
	var err error
	config := util.DefaultLogOptions

	// 自定义配置
	for _, opt := range opts {
		opt(&config)
	}

	if _, zapLogger, err = mode.LogInit(&config); err != nil {
		return err
	}

	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(1))

	return nil
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:debug,msg:", msg)
		return
	}
	zapLogger.Debug(msg, addGoID(fields)...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:info,msg:", msg)
		return
	}
	zapLogger.Info(msg, addGoID(fields)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:warn,msg:", msg)
		return
	}
	zapLogger.Warn(msg, addGoID(fields)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:error,msg:", msg)
		return
	}
	zapLogger.Error(msg, addGoID(fields)...)
	// 可能需要发通知 todo
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:DPanic,msg:", msg)
		return
	}
	zapLogger.DPanic(msg, addGoID(fields)...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:panic,msg:", msg)
		return
	}
	zapLogger.Panic(msg, addGoID(fields)...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, fields ...zapcore.Field) {
	if zapLogger == nil {
		fmt.Println("logger not init!!!level:fatal,msg:", msg)
		return
	}
	zapLogger.Fatal(msg, addGoID(fields)...)
}

func addGoID(fields []zapcore.Field) []zapcore.Field {
	var ret = make([]zapcore.Field, 0, len(fields)+1)
	ret = append(ret, util.GoID(gls.GoID()))
	ret = append(ret, fields...)
	return ret
}

// Sync calls the underlying syslogCore's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() error {
	Info("logger closed")
	if zapLogger != nil {
		return zapLogger.Sync()
	}
	return nil
}

// WithOptions clones the zapLogger, applies the supplied Options, and
// returns the resulting AsyncWriterLogger. It's safe to use concurrently.
func WithOptions(opt ...zap.Option) *zap.Logger {
	if zapLogger == nil {
		fmt.Println("logger not init!!!")
		return nil
	}
	return zapLogger.WithOptions(opt...).With(util.GoID(gls.GoID()))
}
