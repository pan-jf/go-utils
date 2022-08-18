package batchworker

import (
	"fmt"
	"testing"
	"time"
)

func wk(i int64) error {
	if i%2 == 0 {
		return fmt.Errorf("hello")
	}
	return nil
}

func TestBatchWorker(t *testing.T) {
	pool := NewWorkerPool(2, 256)

	bw := NewBatchWorker(pool)
	for i := int64(0); i < 1; i++ {
		idx := i
		bw.Do(func() error {
			return wk(idx)
		})
	}
	errs := bw.Wait()

	for _, e := range errs {
		ee := e.(*Error)
		fmt.Printf("%#v\n", ee)
	}
	fmt.Println("error -------->", errs)

	pool.Stop()

	time.Sleep(1 * time.Second)
}

func TestError(t *testing.T) {
	pool := NewWorkerPool(2, 256)

	bw := NewBatchWorker(pool)
	bw.Do(func() error {
		return nil
	})
	err := bw.Wait()
	if err != nil {
		t.Log("err:", err)
	}
}
