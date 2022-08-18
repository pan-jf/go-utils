package mode

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pan-jf/go-utils/plog/util"
)

type localLogInitializer struct {
}

func (mLog *localLogInitializer) logInit(config *util.Options) (zap.AtomicLevel, *zap.Logger, error) {
	var (
		zapConfig zap.Config
		zapLevel  zap.AtomicLevel
		logger    *zap.Logger
		err       error
	)

	zapConfig = zap.NewDevelopmentConfig()

	zapConfig.DisableStacktrace = true
	zapConfig.EncoderConfig.TimeKey = reserveKeyTimeStamp
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err = zapConfig.Build()
	zapLevel = zapConfig.Level

	return zapLevel, logger, err
}
