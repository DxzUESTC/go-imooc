package main

import "fmt"

// panic机制是比error更加强硬的机制

// error机制不够强硬
// 使用error不会对调用者产生限制，调用者仍然可以无视错误，强行拿到返回结果
// panic理论上会在处罚的位置结束程序运行
// 但是又可以通过recover恢复，保证后续逻辑的执行完毕
// defer 加一个匿名函数func(){}()，将recover恢复操作放到里面用
// 这种叫用defer注册一个匿名延迟函数
// 这样做的原因是recover是函数层面的，而且必须在defer注册的函数中直接调用才会生效
// 这样建立了一个清晰的恢复边界

func divide(x, y float32) float32 {
	if y == 0 {
		panic("cannot divide by zero")
	}
	return x / y
}

func devideByZero() float32 {
	defer func() { //defer注册的函数无论是正常退出还是一场退出，都会执行
		// 用一个if判断，再加一个变量接住recover，能接到东西表示是panic了，然后再去恢复
		if r := recover(); r != nil {
			println("recovered from panic")
			println(r)
			// 异常记录，异常处理流程
		}
	}()
	res := divide(10, 0)
	fmt.Println("divide called")
	return res
}

func main() {
	// 调用上面创建的divide函数
	res := devideByZero()
	fmt.Println("devideByZero called")
	fmt.Println(res)
}
