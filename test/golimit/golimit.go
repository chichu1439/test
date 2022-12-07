package main

import (
	"fmt"
	"github.com/zh-five/golimit"
	"log"
	"time"
)

func main() {
	log.Println("开始测试...")
	start := time.Now()
	g := golimit.NewGoLimit(0) //max_num(最大允许并发数)设置为2

	for i := 0; i < 10; i++ {
		//并发计数加1.若 计数>=max_num, 则阻塞,直到 计数<max_num
		g.Add()
		go func(g *golimit.GoLimit, i int) {
			defer g.Done() //并发计数减1

			time.Sleep(time.Second * 1)
			log.Println(i, "done")
		}(g, i)
		if i > 5 {
			g.SetMax(0)
			break
		}
	}

	log.Println("循环结束")

	g.WaitZero() //阻塞, 直到所有并发都完成
	log.Println("测试结束")
	fmt.Println("cost", time.Now().Sub(start))

}
