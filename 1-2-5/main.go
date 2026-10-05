package main

import (
	"fmt"
)

// 了解切片，本质是一个结构体，长度可变，底层引用数组，组成是（指针+len+cap）
// []t 就是切片，赋值拷贝只复制切片头，共享底层数组
// [n]t 就是数组，拷贝会完整复制全部元素

func main() {

	// 切片的创建
	r := [5]int{1, 2, 3, 4, 5}
	s := r[1:3]
	//println(s[0])  // 这个会打印 ASCII
	fmt.Println(len(s), cap(s))
	// 切片 cap：从切片起始位置，到底层数组末尾总元素数量。所以这里s的cap是4
	// 但是这里也不能用s[3] = 0来赋值
	// len是用来索引的，cap只表示底层数组容量

	s1 := s[1:2]
	s2 := s[:2]
	s3 := s[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// r2 := [5]int{1,2,3,4,5}
	// s4 := r2[:]
	s4 := []int{1, 2, 3, 4, 5} // 数组个数的位置什么都不写时变成定义切片
	fmt.Println(s4)

	s5 := make([]int, 10) // make方法很少使用，也是用来创建切片的
	fmt.Println(s5)

	// 切片的访问
	// 直接使用下标
	// 遍历
	for _, v := range s {
		fmt.Println(v)
	}
	// len()可以打印长度
	fmt.Println(len(s5))

	// 切片的追加
	// s[0] ~ s[3]
	s = append(s, 0)
	fmt.Println(s)
}
