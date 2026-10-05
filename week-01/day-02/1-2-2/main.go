package main

// 掌握switch语法的使用

func main() {
	var grade rune = 'C'
	switch grade {
	case 'C':
		println("要加油")
		//break
		if false {
			break
		}
		println("...")
		fallthrough
	case 'B':
		println("还不错")
		//break
	case 'A':
		println("恭喜")
		//break
	default:
		println("要努力")
	}
}
