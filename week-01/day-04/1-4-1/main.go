package main

// Goroutines
// 并发的执行函数，如果不是函数要包装成匿名函数执行
// 不是协程，但是中文意会，就叫他协程
// Go中所有的逻辑均在 Goroutines 中执行，包括main()

import (
	"fmt"
	"time"
)

func remindMeLater(t time.Duration, notice string) {
	time.Sleep(t)
	println(notice)
}

func main() {
	go remindMeLater(5*time.Second, "学 Golang")
	go remindMeLater(10*time.Second, "学 Golang")
	go func() { // 使用匿名函数来并发执行
		println("Hello World")
	}() // 这个括号表示匿名函数的调用

	println("main() end")
	fmt.Scanf("%d")
}
