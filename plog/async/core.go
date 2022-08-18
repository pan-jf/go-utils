package async

import (
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

// FileCore ...
type FileCore struct {
	zapcore.LevelEnabler

	encoder     zapcore.Encoder
	asyncLogger *WriterLogger
}

// NewAsyncFileCore ...
func NewAsyncFileCore(
	zapLevelEnabler zapcore.LevelEnabler,
	encoder zapcore.Encoder,
	fileName string) (*FileCore, error) {
	asyncLogger, err := NewAsyncWriteLogger(fileName)
	if err != nil {
		return nil, err
	}
	return &FileCore{
		LevelEnabler: zapLevelEnabler,
		encoder:      encoder,
		asyncLogger:  asyncLogger,
	}, nil
}

// With ...
func (c *FileCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	for _, field := range fields {
		field.AddTo(clone.encoder)
	}
	return clone
}

// Check ...
func (c *FileCore) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checked.AddCore(entry, c)
	}
	return checked
}

// Write ...
func (c *FileCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// Generate the message.
	buffer, err := c.encoder.EncodeEntry(entry, fields)
	if err != nil {
		return errors.Wrap(err, "failed to encode log entry")
	}

	msg := buffer.String()
	err = c.asyncLogger.WriteString(msg)
	buffer.Free()
	return err

}

// Sync ...
func (c *FileCore) Sync() error {
	if c.asyncLogger != nil {
		return c.asyncLogger.Close()
	}
	return nil
}

func (c *FileCore) clone() *FileCore {
	return &FileCore{
		LevelEnabler: c.LevelEnabler,
		encoder:      c.encoder.Clone(),
		asyncLogger:  c.asyncLogger,
	}
}
