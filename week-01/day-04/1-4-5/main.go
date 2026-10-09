package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 掌握go中的错误和异常的处理方式
// error是一个接口，实现了 Error() string方法
// 本质是通过返回值处理错误

// 自己写的函数怎么使用错误处理，自定义error
// 以除法函数为例
func divide(x, y float32) (float32, error) { // error是一个接口
	if y == 0 {
		return 0.0, errors.New("divide by zero")
	}
	return x / y, nil
}

func main() {
	// string 转 int "123"-> 123
	// sdk runtime
	//i, _ := strconv.Atoi("moody") // 输入改成非法的字符，但是i还是接收到了合法的0
	i, err := strconv.Atoi("123")
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(i)
	}
}
