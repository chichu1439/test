package dy_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	a := NewConcurrencyLimiter(10)
	a.Get()
	a.Release()
	a.Reset(5)

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		a.Get()
		go func(j int) {
			wg.Done()
			defer a.Release()
			//time.Sleep(time.Second(1).)
		}(i)
	}
	wg.Wait()
	fmt.Println("cost", time.Now().Sub(start))
}
