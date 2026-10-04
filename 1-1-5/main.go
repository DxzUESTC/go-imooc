package main

import "unsafe"

func main() {
	// 布尔
	var bo bool
	println(bo)

	// 数字
	var i1 = 42 // int 8字节，64位
	println(i1)

	var i2 int8 = 127 // -128~127
	println(i2)

	println(unsafe.Sizeof(i1))
	println(unsafe.Sizeof(i2))

	var i3 uint
	println(unsafe.Sizeof(i3))

	var f32 float32 = 0.8
	println(f32)
	println(unsafe.Sizeof(f32))

	var f64 float64 = 0.8
	println(f64)
	println(unsafe.Sizeof(f64))

	var c1 complex64 = 1 + 2i
	println(c1)
	println(unsafe.Sizeof(c1))

	var c2 complex128 = 1 + 2i
	println(c2)
	println(unsafe.Sizeof(c2))

	// 字符串
	var s1 string = "ellho"
	println(s1)
	println(unsafe.Sizeof(s1))

}
