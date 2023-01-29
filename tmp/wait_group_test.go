package tmp

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var ops uint64 = 0
	arr := [1000]string{}
	for i := 0; i < 1000; i++ {
		arr[i] = fmt.Sprintf("%d", i)
	}
	for _, s := range arr {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			tt(s)
			//time.Sleep(10 * time.Second)
			//atomic.AddUint64(&ops, 1)
			//fmt.Println(i)
		}(s)
	}
	wg.Wait()
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

func tt(s string) {
	fmt.Println(s)
}
