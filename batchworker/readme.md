### golang 协程池

#### 使用样例

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	
	"github.com/pan-jf/go-utils/batchworker"
)

type MyTest struct {
	m   sync.RWMutex
	bw  *batchworker.BatchWorker
	num int
}

func (t *MyTest) AddNum() {
	// 加个写锁避免并发问题
	t.m.Lock()
	defer t.m.Unlock()
	t.num++
}

func (t *MyTest) GetNum() int {
	return t.num
}

func main() {
	fmt.Println("start, goroutine count:", runtime.NumGoroutine())
	var runner MyTest

	// 创建最大并发数=2，单次处理长度为100的协程池
	pool := batchworker.NewWorkerPool(2, 100)

	runner.bw = batchworker.NewBatchWorker(pool)

	// 模拟10任务并发，每个任务累加100000次
	for i := 0; i < 10; i++ {
		// 打印是4，其中1个是监听错误的协程，1个是主协程，2个是协程池开辟的
		fmt.Println("current goroutine count:", runtime.NumGoroutine())
		for j := 0; j < 100000; j++ {
			runner.bw.Do(func() error {
				runner.AddNum()
				return nil
			})
		}
	}

	// 接收错误
	errs := runner.bw.Wait()
	for _, err := range errs {
		fmt.Println(err)
	}

	pool.Stop()

	fmt.Println("result:", runner.GetNum())
}

```