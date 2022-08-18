package mode

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pan-jf/go-utils/plog/async"
	"github.com/pan-jf/go-utils/plog/util"
)

type prodLogInitializer struct {
}

func epochFullTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func (prod *prodLogInitializer) logInit(config *util.Options) (zap.AtomicLevel, *zap.Logger, error) {
	var (
		lLevel  zap.AtomicLevel
		lZapLog *zap.Logger
	)

	// Initialize Zap.
	encConf := zap.NewProductionEncoderConfig()
	encConf.TimeKey = reserveKeyTimeStamp
	encConf.EncodeTime = epochFullTimeEncoder
	encoder := zapcore.NewJSONEncoder(encConf)

	lLevel = zap.NewAtomicLevelAt(zap.DebugLevel)

	core, err := async.NewAsyncFileCore(lLevel, encoder, util.GetLogFilePath(config))

	if err != nil {
		return lLevel, lZapLog, err
	}

	lZapLog = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel))

	return lLevel, lZapLog, nil
}
