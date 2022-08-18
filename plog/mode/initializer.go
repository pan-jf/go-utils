package mode

import (
	"runtime"

	"go.uber.org/zap"

	"github.com/pan-jf/go-utils/plog/util"
)

// zap config 配置常量值
const (
	reserveKeyTimeStamp = "ts"
)

type logInitializer interface {
	logInit(*util.Options) (zap.AtomicLevel, *zap.Logger, error)
}

// LogInit 初始化日志
func LogInit(config *util.Options) (zap.AtomicLevel, *zap.Logger, error) {
	var (
		zapInit logInitializer
		level   zap.AtomicLevel
		logger  *zap.Logger
		err     error
	)

	if runtime.GOOS == "linux" {
		// 线上使用。不输出到控制台，保存到文件里
		zapInit = &prodLogInitializer{}
	} else {
		//调试使用。不保存日志，在控制台查看
		zapInit = &localLogInitializer{}
	}

	if level, logger, err = zapInit.logInit(config); err != nil {
		return level, logger, err
	}

	return level, logger, nil
}
