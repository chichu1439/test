package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sync"
	"time"
)

func bcryptD() string {
	cost := bcrypt.DefaultCost
	tc := float64(0)
	hash := []byte("")
	e := error(nil)
	passwd := ".PinNo" + "abcdefghijk01234" + "11contId, 3, 16)"
	for tc <= 250 && cost <= bcrypt.MaxCost {
		startT := time.Now()
		hash, e = bcrypt.GenerateFromPassword([]byte(passwd), cost)
		if e != nil {
			fmt.Println(e)
		}
		tc = time.Since(startT).Seconds() * 1000
		//log.Debugsf("tc = %v", tc)
		cost += 1
		//log.Debugsf("iterations = %v", cost)
	}
	pinNo := base64.StdEncoding.EncodeToString(hash)
	return pinNo
}
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	c := make(chan struct{}, 3)

	for i := 0; i < 1000; i++ {
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
	fmt.Println("cost", time.Now().Sub(start))
}
