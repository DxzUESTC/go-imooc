package main

import "fmt"

// 方法对应的就是java等面向对象语言的对象的方法
// 但是结构体里没法写方法
// 写出来带有接受者的函数就是绑定了结构体的方法

// 方法与函数的区别
// 函数是没有接收者，不和结构体绑定
// 函数和方法有类似的语法，参数都是值传递

// * 号和 & 号的使用
// *跟变量，解引用
// *跟变量定义或参数定义，指针定义
// & 只有取地址一个作用

// 全局的结构体
type myLesson struct {
	id   int
	name string
	rate float32
}

type myTextLesson struct {
	myLesson
	wordcount int
}

// go没有class的概念，但是又想实现一部分的oo特征
// 注意给哪个结构体写的方法这里receiver这里就要写上它
func (l *myLesson) updatename(newName string) { // 注意方法这里还是值传入，所以传入指针是更合理的实现
	l.name = newName
	//println(&l)
}

func main() {
	var l1 myLesson = myLesson{id: 123, name: "MySQL"}
	//println(&l1)
	l1.updatename("MySQL Plus") // 这里是go语言方法表达式的自动取地址的语法糖
	//(&l1).updatename("MySQL Plus") // 编译时自动转换成这行代码
	fmt.Println(l1.name)

	var l2 myTextLesson
	l2.updatename("MySQL Plus") // 结构体组合之后依然可以使用被组合的结构体的方法
}
