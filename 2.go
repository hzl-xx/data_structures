package main

import (
	"数据结构/Stack"
)

func main() {

	var s = Stack.Stack{}
	s.InitStack(50)
	// 进制转换
	c := "0123456789ABCDEF";

	var (
		jinzhi = 16
	)

	var i = 458
	for {
		mod := i % jinzhi

		s.InStack(string(c[mod]))
		i = i / jinzhi
		if i == 0 {
			break
		}

	}

	s.StackTraverse()

}
