package main

import "fmt"

// 掌握数组的初步使用

func main() {
	// 数组的定义，因为数组是在内存开辟一段连续空间的，所以需要确定长度
	var n1 = [5]int{1, 2, 3, 4, 5}
	println(n1[0])

	n2 := [5]int{1, 2, 3, 4, 5}
	println(n2[0])

	n3 := [...]int{1, 2, 3, 4, 5} // 会自动推断出是五个元素
	println(n3[0])

	var n4 [5]int // 这会初始化五个元素都是 0
	println(n4[0])
	println(n4[1])
	println(n4[4])

	var n5 = [...]int{1: 3, 6: 6} // 能够指定下标对应位置来初始化，其他位置赋值0，数组大小按初始化推断
	println(n5[0])
	println(n5[1])
	println(n5[6])

	// 访问
	println(n5[6])
	n2[1] = 4
	for _, v := range n2 {
		println(v)
	}

	// 多维数组
	t := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(t)

}
