package util

const (
	linuxBaseDir       = "/data/plog"
	defaultLogPath     = "/%s/%s.log"
	logFileWithPidPath = "/%s/%s_%d.log"
)

// Options 参数配置
type Options struct {
	logFileHasPid bool // 设置后在保存日志时会带上pid
	maxMsgLen     int
}

var DefaultLogOptions = Options{
	logFileHasPid: false,
	maxMsgLen:     8 * 1024 * 1024, //
}

// Option 配置函数
type Option func(*Options)

// AddPidToLogFile 是否在日志输出中加入pid信息
func AddPidToLogFile() Option {
	return func(o *Options) {
		o.logFileHasPid = true
	}
}

// MaxMsgSize 单条日志最大长度，默认为2048
func MaxMsgSize(size int) Option {
	return func(o *Options) {
		o.maxMsgLen = size
	}
}
