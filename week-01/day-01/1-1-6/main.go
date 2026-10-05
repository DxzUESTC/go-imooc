package main

func main() {
	//const a int = 0
	//const b int = 1 // 常量必须赋初始值
	//const c = "imooc"

	const (
		a = 0
		b = 1
		c = "imooc"
	)

	//iota go里的自增常量
	const (
		a2 = iota
		b2 = iota
		c2 = iota
		d2

		e2 = 1 << iota // 表示 1 左移 4 位，这时候 iota 已经自增成4，<< 左移运算符，右侧是移动位数
	)

	println(a2, b2, c2)

	const (
		i1= iota
		i2         = 1
		a3, b3, c3 = 1, false, "tryit"
	)
}
