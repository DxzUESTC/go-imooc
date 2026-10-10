package main

// go语言的隐式接口
// 隐式接口更好还是显式接口更好?
// 像java，定义接口后，需要用implement显式说明实现
// 像go，写完接口，然后用结构体方法来实现，不用显式说明
// go隐式接口可以在不改变原有的业务代码的基础上实现更多的接口

//type wuFanBanLv struct{} //这里不应该使用struct来描述一个行为
//func (w wuFanBanLv) play() {}
//func eat(w wuFanBanLv) {
//	w.play()
//	println("eating...")
//}
// 简单逻辑使用 结构体 和 方法 能够满足需求，但是这里wuFanBanLv本质是一类事务的抽象
// 当业务复杂或者后续需要更多的补充时，对于原始代码修改会非常繁多

// 接口像struct结构体，但是关键字不是struct，是interface
// 然后，接口只能声明方法，是行为契约
// 接口不依赖结构体，但他脱离不了具体类型的实现
// 就理解为一组方法签名的集合

// 接口示例
type wuFanBanLv2 interface {
	play()
}

type myLesson struct {
	id   int
	name string
}

type myMovie struct {
	title  string
	actors string
}

// 上述两个结构体本身无关联，但是行为相似，都能进行播放
func (m myMovie) play() {
	println("playing movie", m.title)
}
func (m myLesson) play() {
	println("playing lesson")
}

// 接口；结构体；结构体方法

func main() {
	var w wuFanBanLv2 = myMovie{title: "FightClub"} //因为myMovie实现了wuFanBanLv2这个接口，所以可以这么写
	w.play()
}
