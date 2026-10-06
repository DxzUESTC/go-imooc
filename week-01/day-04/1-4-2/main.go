package main

import (
	"fmt"
	"time"
)

// go语言中的Channel
// 这里实现一个异步的字符串处理器
// 在main函数读取用户输入
// 处理输入，每个输入处理约10秒钟

func processString(c chan string) {
	for s := range c { // 可以用range来迭代传递管道中的数值
		fmt.Println("[processString] start process", s)
		time.Sleep(10 * time.Second)
		fmt.Println("[processString] finish process", s)
	}
}

func main() {
	// Channel的实现
	c := make(chan string, 10) // 如果channel没有缓存，即不写第二段的数字，数据会阻塞在这里无法传递
	c <- "moody"
	test := <-c
	fmt.Println(test)

	c1 := make(chan string, 10)
	var s string
	go processString(c1)
	fmt.Println("[main] start...")
	for true {
		fmt.Scanf("%s\n", &s)
		c1 <- s
	}

}

// Channel声明方法
// make(chan int) 无缓存
// make(chan int, 0) 无缓存
// make(chan int, 1) 有缓存

// Channel基本用法
// ch <- x     发送数据到管道ch里面
// x := <- ch  从管道ch里面取数据到x
// _ <- ch     从管道ch里面取数据并丢弃
// <- ch       从管道ch里面取数据并丢弃
