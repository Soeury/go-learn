package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer fmt.Println("test-----defer")
	runtime.Goexit() // 注意这里和 return 不一样
	// return              // return 会打印执行 goroutine 后面的打印
	fmt.Println("test") // 这里不会执行
}

func main() {

	// goroutine  -  runtime.Goexit()

	// runtime.Goexit() 退出当前 goroutine 协程 ，
	// return 退出当前函数 ， goroutine 继续执行

	go func() {
		fmt.Println("before")
		test()
		fmt.Println("after")
	}()

	time.Sleep(time.Second * 5)
}
