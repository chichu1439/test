package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	numCPUs := runtime.NumCPU()
	fmt.Println(numCPUs)
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte
		time.Sleep(time.Second * 1)

		return result
	})
	defer pool.Close()
	for i := 0; i < 50; i++ {
		go pool.Process(i)
	}

	fmt.Println("cost", time.Now().Sub(start))
}
func Process() {

}
