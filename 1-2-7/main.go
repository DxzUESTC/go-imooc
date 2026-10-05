package main

// 学习指针的基本知识
// 指针记录的是原变量在内存中的地址
// 指针本身也是个数字变量，其自己也有指针

func main() {
	i := 9            // 变量处在内存中
	var p *int = &i   // &i表示变量在内存中的位置，这里的p是指针，存的就是位置
	var pp **int = &p // pp表示为指针的指针

	println(i)
	println(*p) // *pointer表示取地址的内容
	println(&i)
	println(p) // pointer本身保存内存地址

	println(&p)
	println(pp)

	// 警惕值传递的陷阱
	add(i)
	println(i) // 这里并不会改变i这个变量的值，因为go中参数都是值传递

	addbypointer(p)
	println(i) // 这里传入的是指针，也就是传入了变量实际所在的内存地址

}

func add(j int) {
	j++
}

func addbypointer(p *int) {
	*p++
}
