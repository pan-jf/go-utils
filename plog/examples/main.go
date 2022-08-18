package main

import (
	"time"

	"github.com/pan-jf/go-utils/plog"
)

func main() {
	for {
		plog.Debug("this is a debug log")
		plog.Info("this is a info log")
		plog.Warn("this is a warn log")
		time.Sleep(time.Second)
	}
}
