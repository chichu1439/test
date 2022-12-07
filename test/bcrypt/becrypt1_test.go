package main

import (
	"sync"
)

func main() {
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
