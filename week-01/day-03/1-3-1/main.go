package main

import "fmt"

// 学习go语言的结构体知识
// go是没有类的

func main() {
	// id int
	// name string
	// rating float
	l1 := struct {
		name string
	}{} // 第二个大括号是初始化用的，但是这样使用struct非常麻烦
	l1.name = "MySQL"

	// type 关键字
	// 可以起别名
	//type ii int
	//var i1 ii = 10

	type my_lession struct {
		id     int
		name   string
		rating float32
	}
	l2 := my_lession{id: 416, name: "MySQL", rating: 10.0}

	// 访问
	fmt.Println(l2.rating)

	// 继承，组合结构体，其实不是继承
	type my_text_lession struct {
		my_lession
		wordcount int
	}
	l3 := my_text_lession{id: 416, name: "MySQL", rating: 10.0, wordcount: 5}
	fmt.Println(l3.my_lession.id, l3.my_lession.name)
	fmt.Println(l3.rating) // go提供的一种语法糖
	fmt.Println(l3.wordcount)

}
