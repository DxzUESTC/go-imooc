package main

import (
	"fmt"
	"time"
)

// 使用select语句等待多个channel
// 年会抢答
// 三个channel传回抢答
// c1 c2 c3 监听
// 打印最快的答案

func sender1(c chan string) {
	time.Sleep(7 * time.Second)
	c <- "i am sender 1"
}

func sender2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "i am sender 2"
}

func sender3(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "i am sender 3"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	fmt.Println("start")
	go sender1(c1)
	go sender2(c2)
	go sender3(c3)

	// 打印最快的答案
	//println(<-c1)
	//println(<-c2)
	//println(<-c3) // 这种只会阻塞然后顺序执行，无法实现打印最快答案的要求

	select { // select 能够做到哪个case先满足条件先进行操作，操作完毕即退出
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	case s3 := <-c3:
		fmt.Println(s3)
	default: // 如果没写default，select会一直等待case
		fmt.Println("no answer")
	}

}
// select 使用方法
// select：执行准备好的channel
select {
case <- channel1:
	//执行代码
case value := <- channel2:
	//执行代码
case channel3 <- value:
	//执行代码
default:
	//所有通道都没有准备好，执行的语句
}