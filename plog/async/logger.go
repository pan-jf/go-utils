package async

import (
	"context"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// 缓存切片容量
	maxChanSize = 256 * 1024
	maxFileSize = 4 * 1024 // 4GBytes
	maxBackups  = 10
	maxAge      = 7
)

// WriterLogger 异步写日志
type WriterLogger struct {
	writer  *lumberjack.Logger
	closed  bool
	msgChan chan string
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

// NewAsyncWriteLogger 对外接口
func NewAsyncWriteLogger(filename string) (*WriterLogger, error) {
	l := WriterLogger{
		writer: &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxFileSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   false, // 是否压缩日志
			LocalTime:  true,
		},
		msgChan: make(chan string, maxChanSize),
	}

	l.ctx, l.cancel = context.WithCancel(context.Background())
	l.wg.Add(1)
	go func(lg *WriterLogger) {
		for {
			if logLoop(lg) {
				break
			}
		}
		lg.wg.Done()
	}(&l)
	return &l, nil
}

// WriteString 写日志
func (l *WriterLogger) WriteString(msg string) error {
	_, err := l.writer.Write([]byte(msg))
	return err
}

// Close 关闭log
func (l *WriterLogger) Close() error {
	if l.closed {
		return nil
	}
	l.closed = true
	if l.cancel != nil {
		l.cancel()
	}
	l.wg.Wait()
	return l.writer.Close()
}

func logLoop(l *WriterLogger) bool {
	defer func() {
		recover()
	}()

	var msg string
	closed := false

	for {
		if !closed {
			select {
			case msg = <-l.msgChan:
			case <-l.ctx.Done():
				closed = true
			}
		} else {
			select {
			case msg = <-l.msgChan:
			default:
				return true
			}
		}
		if len(msg) > 0 {
			_ = l.WriteString(msg)
			msg = ""
		}
	}
}
