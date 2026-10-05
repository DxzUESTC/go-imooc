package main

import "fmt"

// 掌握 for 循环语句写法

func main() {
	// 经典写法
	// 使用 fori 快捷生成模板
	for j := 0; j < 5; j++ {
		println(j)
	}

	// for each 迭代语句
	// 使用 forr 快捷生成模板
	// 示例
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range arr {
		fmt.Printf("%d\t%d\n", i, v)
	}

	// for 还可以相当于其他语言的 while 循环语法
	i := 0
	for i < 5 {
		println(i)
		i = i * 2
		i++
	}

	// 这是一个死循环
	for {
		println(i)
		if i > 5 { // 这种死循环里也可以用条件判断跳出循环
			break
		}
	}
}
