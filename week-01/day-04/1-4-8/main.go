package main

// 了解go语言中的泛型
// 1.18版本才有
// 有点缝合进来的意味

// 用一个比较大小，返回更小值的例子来认识go的泛型
// 最简单的实现就是在函数的名称和入参之间加一个中括号，把对应的基本数据类型添加进去
// func funcname[yourType baseType1 | baseType2 | ~baseType3](c1, c2 yourType){...}
// 这里例子中的~表示以baseType3为基础派生出来的各种类型 类似 type myInt int

// go中还有把中括号里的这堆东西抽象出来的方法
// 但是抽象的是，go用接口来实现

// go语言中有一个cmp.Ordered，也是接口，写好了可以比较的各种类型

// 简单泛型
func min1[myNumber int | float32 | float64](a, b myNumber) myNumber {
	if a < b {
		return a
	}
	return b
}

// 接口实现泛型示例
type myNumber2 interface {
	int | float32 | float64
}

func min2[T myNumber2](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(min2(2, 9))
}
