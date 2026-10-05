package main

import "unsafe"

func main() {
	println(add(1, 2))
	println(swap(3, 2))
	func1(1, 2, 3)
	func1(4, 5)

	var f = add // f变成函数类型，并且可以被调用
	println(f(1, 3))

	var jj = 1
	println(jj)
	plus(jj)
	println(jj) // 还是1，go函数也是值传递
}

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func swap(i int, j int) (int, int) {
	return j, i
}

// 可变参数
func func1(ii ...int) {
	println(ii)
	println(unsafe.Sizeof(ii))
}

// 函数也可以作为另一个函数的入参或者出参
func math() func(i int, j int) int {

	return func(i int, j int) int {
		return i + j
	}
}

// 值传递
func plus(i int) {
	i++
}
