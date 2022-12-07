package test

import (
	"fmt"
	"sync"
	"time"
)
var mutex sync.Mutex
func Abc(a string)  {
	mutex.Lock()
	fmt.Println("the lock is locked")
	channe :=make([]chan int,4)
	for i, _ := range channe {
		channe[i] = make(chan int)
		go func(i int,c chan int) {
			fmt.Println("Not lock:",i)
			mutex.Lock()
			fmt.Println("Locked:",i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock:",i)
			mutex.Unlock()
			c<-i
		}(i,channe[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, ints := range channe {
		<-ints
	}
}
