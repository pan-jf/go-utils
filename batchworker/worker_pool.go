package batchworker

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type poolWorker struct {
	worker  Worker
	errChan chan error
}

// WorkerPool 工作协程池
//
// Notes: 能够多个处理同时使用或多次处理复用同一个协程池
type WorkerPool struct {
	// 最大同时处理并发数
	MaxSize int64

	size int64
	work chan poolWorker
	stop chan interface{}
}

// NewWorkerPool ...
func NewWorkerPool(maxSize int64, chanSize int64) *WorkerPool {
	return &WorkerPool{
		MaxSize: maxSize,
		work:    make(chan poolWorker, chanSize),
		stop:    make(chan interface{}, 16),
	}
}

// Do ...
func (p *WorkerPool) Do(w Worker, c chan error) {
	s := atomic.AddInt64(&p.size, 1)
	if s <= p.MaxSize {
		go p.doWork()
	}

	p.work <- poolWorker{
		worker:  w,
		errChan: c,
	}
}

// Stop ...
func (p *WorkerPool) Stop() {
	s := p.size
	if p.size > p.MaxSize {
		s = p.MaxSize
	}
	for i := int64(0); i < s; i++ {
		p.stop <- nil
	}
}

func (p *WorkerPool) doWork() {
	for {
		select {
		case pw := <-p.work:
			func() {
				defer func() {
					if pn := recover(); pn != nil {
						pw.errChan <- errors.New(fmt.Sprintf("worker panic recovered: %v", pn))
					}
				}()
				err := pw.worker()
				if err != nil {
					pw.errChan <- err
				}
			}()
		case <-p.stop:
			return
		}
	}
}
