package main

import "fmt"

func main() {
	result := add()
	for i := 0; i < 5; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Printf("sumOld=%d \n", result(i))
	}
	result2 := add()
	fmt.Println("下面是一个新实例")
	for i := 5; i < 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Printf("sumNew=%d \n", result2(i))
	}
	//这里重新调用一下result实例，发现返回的sumOld值还是接续上面的
	fmt.Println("重新调用result实例")
	fmt.Printf("i=10 \t")
	fmt.Printf("sumOld=%d \n", result(10))
}

//通过闭包来访问其中的变量
func add() func(num int) int {
	sum := 0
	addResult := func(num int) int {
		sum += num
		return sum
	}
	return addResult
}
