package plog

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pan-jf/go-utils/plog/util"
)

type payload struct {
	Level *zapcore.Level `json:"level"`
}

func TestJsonPayLoad(t *testing.T) {
	level := zapcore.Level(-1)
	pl := payload{
		Level: &level,
	}
	jsonStr, err := json.Marshal(pl)
	if err != nil {
		t.Errorf("failed err:%v", err)
	}
	t.Logf("json str:%v", string(jsonStr))
}

func TestLog(t *testing.T) {
	nums := 10
	var s = ""

	for i := 0; i < nums; i++ {
		s += fmt.Sprintf("%d, ", i)
		Debug("new", zap.String("msg", s))
	}
	_ = Sync()

}

func TestMaxMsgSize(t *testing.T) {
	_ = initZapLog(util.MaxMsgSize(3000000))
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for range ticker.C {
			Debug("a1")
			Debug("max")
			Info("max")
		}

	}()
	time.Sleep(time.Second * 2)

}

func TestWithPid(t *testing.T) {
	_ = initZapLog(util.AddPidToLogFile())
	Debug("test")
}

func TestDurationField(t *testing.T) {
	tt := time.Now()
	n := tt.Add(time.Millisecond * 100)
	Info("duration", zap.Duration("d", n.Sub(tt)))
	Info("duration", zap.Duration("d", time.Millisecond*10))
	Info("duration", zap.Duration("d", time.Millisecond*100))
	Info("duration", zap.Duration("d", time.Millisecond*1000))
	Info("duration", zap.Duration("d", time.Second*75))

}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Debug("test")

	}
}

func TestWithOption(t *testing.T) {
	wrap()
	go func() {
		wrap()
	}()
	Info("aa", zap.String("key", "string"))
	time.Sleep(1 * time.Second)
}

func wrap() {
	WithOptions().Info("aa", zap.String("key", "string"))
}

func BenchmarkWrap(b *testing.B) {
	org, withOpt := "origin", "with_option"
	benchmarks := []struct {
		name string
	}{
		{org}, {withOpt},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				switch bm.name {
				case org:
					WithOptions(zap.AddCallerSkip(0)).Info("aa")
				case withOpt:
					Info("aa")
				}
			}
		})
	}
}
