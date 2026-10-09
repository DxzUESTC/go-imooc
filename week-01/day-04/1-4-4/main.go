package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何使用互斥锁同步Goroutines
// 1000个协程同时操作变量的并发问题
// 使用sync包的Mutex来定义一个互斥锁解决并发问题
var m = sync.Mutex{}

// 注意这里不是给具体的变量加锁，而是给操作加锁
func add(p *int) {
	m.Lock()
	defer m.Unlock() // 这个关键字会让后续语句在函数退出的时候执行
	*p++
	// 当业务逻辑比较长，且不巧的是出现了问题，就无法解锁，其他协程就会卡在自己的上锁这一步
	// m.Unlock()
}

func main() {
	i := 0
	for j := 0; j < 1000; j++ {
		go add(&i)
	}
	time.Sleep(3 * time.Second)
	// 也可以用waitGroup等待携程都跑完
	fmt.Println(i) //这里不加锁并不会输出1000，会有并发问题
}
