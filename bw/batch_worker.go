package bw

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// Worker 工作函数
type Worker func() error

// Error ...
type Error struct {
	Index int64
	Err   error
}

// Error ...
func (e *Error) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// ErrorList ...
type ErrorList []error

// Error ...
func (l ErrorList) Error() string {
	msg := strings.Builder{}
	for i, e := range l {
		if i != 0 {
			msg.WriteString(", ")
		}
		msg.WriteString(fmt.Sprintf("[%v]: ", i))
		msg.WriteString(e.Error())
	}
	return msg.String()
}

// BatchWorker 批量工作
type BatchWorker struct {
	// 是否根据调用 Do 的顺序对错误进行排序, 默认 true
	SortErrors bool

	pool        *WorkerPool
	workerIndex int64
	workerWg    sync.WaitGroup

	handleErrWg    sync.WaitGroup
	handleErrStart sync.Once

	errChan chan error
	errs    []error
}

// NewBatchWorker ...
func NewBatchWorker(pool *WorkerPool) *BatchWorker {
	errChan := make(chan error, 64)
	errs := make([]error, 0, 64)

	return &BatchWorker{
		SortErrors: true,

		pool:    pool,
		errChan: errChan,
		errs:    errs,
	}
}

// Do ...
func (bw *BatchWorker) Do(w Worker) {
	bw.handleErrStart.Do(func() {
		bw.handleErrWg.Add(1)
		go bw.handleError()
	})

	idx := atomic.AddInt64(&bw.workerIndex, 1) - 1
	bw.workerWg.Add(1)

	bw.pool.Do(func() error {
		err := w()
		bw.workerWg.Done()
		if err != nil {
			return &Error{
				Index: idx,
				Err:   err,
			}
		}
		return nil
	}, bw.errChan)
}

// Wait ...
func (bw *BatchWorker) Wait() ErrorList {
	bw.workerWg.Wait()

	close(bw.errChan)
	bw.handleErrWg.Wait()

	if bw.errs != nil && bw.SortErrors {
		sort.Slice(bw.errs, func(i, j int) bool {
			ie, ok := bw.errs[i].(*Error)
			if !ok {
				return false
			}
			je, ok := bw.errs[j].(*Error)
			if !ok {
				return false
			}
			return ie.Index < je.Index
		})
	}

	if len(bw.errs) != 0 {
		return bw.errs
	}
	return nil
}

func (bw *BatchWorker) handleError() {
	for {
		select {
		case err, ok := <-bw.errChan:
			if !ok {
				bw.handleErrWg.Done()
				return
			}
			bw.errs = append(bw.errs, err)
		}
	}
}
