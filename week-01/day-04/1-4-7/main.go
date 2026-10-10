package main

import (
	"fmt"
	"reflect"
)

// go中的反射
// 用的不多 要能看懂
// 反射：用数据来表示元数据，来拿到数据的属性
// 基本上都是在reflect包里的方法
type myLesson struct {
	id     int
	name   string
	rating float32
}

func (ml myLesson) PrintName() {
	fmt.Println(ml.name)
}

func main() {
	m := myLesson{515, "MySQL", 10.0}

	// 对m这个类型里的字段，方法进行反射
	//用到TypeOf，TypeOf是一个函数
	typeOfm := reflect.TypeOf(m)
	// TypeOf 返回的是一个接口，这个接口叫 Type
	// 比如想看有没有rating这个字段，可以用这个接口的一个函数FieldByName
	// 这个函数返回两个东西，一个是包含有这个字段信息的结构体，一个是是否存在的bool值
	// 可以按住ctrl点进这个函数看实现
	fieldByName, has := typeOfm.FieldByName("rating")
	fmt.Println(fieldByName.PkgPath, has)

	// MethodByName返回两个东西，一个是返回方法的信息，一个是是否存在的信息
	methodByName, has := typeOfm.MethodByName("PrintName")
	fmt.Println(methodByName, has)

	// 对m这个值进行反射
	// 用到了reflect包中的ValueOf
	valueOfm := reflect.ValueOf(&m).Elem()
	// ValueOf返回的仍然是一个接口
	// 这个接口有很多的方法
	// 下面这个方法将m对应的字段置为初始，int字段为0，string字段为空字符串
	valueOfm.SetZero()
	fmt.Println(m)
}
