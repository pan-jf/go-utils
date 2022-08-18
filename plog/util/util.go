package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ProcessName() string {
	pName := path.Base(os.Args[0])
	if runtime.GOOS == "windows" {
		return filepath.Base(pName)
	} else {
		return strings.TrimSuffix(pName, "_d") + "_rd"
	}
}

// GetLogFilePath 由于日志文件配套工具有相关限制，故不提供灵活的文件路径
func GetLogFilePath(opt *Options) string {
	baseDir := os.TempDir() + string(os.PathSeparator) + "plog" + string(os.PathSeparator)

	if runtime.GOOS == "linux" {
		baseDir = linuxBaseDir
	}

	if opt.logFileHasPid {
		return baseDir + fmt.Sprintf(logFileWithPidPath, ProcessName(), ProcessName(), os.Getpid())
	}
	return baseDir + fmt.Sprintf(defaultLogPath, ProcessName(), ProcessName())
}
