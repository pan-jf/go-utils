package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logCommonKeyGoID = "GoID"
)

// GoID 协程ID
func GoID(goID int64) zapcore.Field {
	return zap.Int64(logCommonKeyGoID, goID)
}
