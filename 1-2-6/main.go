package main

import "fmt"

// 掌握go语言的map基础语法用法

func main() {
	// 创建
	// 使用make方法创建空map
	m1 := make(map[int]string)
	m1[461] = "RabbitMQ"
	m1[515] = "MySQL"
	m1[576] = "Golang"
	fmt.Println(m1)

	// 创建字面量
	m2 := map[int]string{
		461: "RabbitMQ",
		515: "MySQL",
		576: "Golang",
	}
	fmt.Println(m2[576])

	// 遍历
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	// 删除
	delete(m2, 515)
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	fmt.Println(m2[515]) // 注意这里还会打印出来空值
	fmt.Println("---")

	// 删除之后可以判断key对应的value是否存在
	s, ok := m2[515]
	fmt.Println(s, ok) // s是空的，但是ok对应的bool值为false

	// 使用len()判断长度
	fmt.Println(len(m2))
}
