package main

import (
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	c := make(chan struct{}, 5)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			defer func() {
				<-c
			}()
			bcryptD()
		}(i)
	}
	wg.Wait()
}
