package main

import (
	"fmt"
	"数据结构/Stack"
)

func main() {

	regStr := "[()][]"

	var cNeed int32 = 0

	var c = Stack.Stack{}
	c.InitStack(len(regStr))

	var p  = Stack.Stack{}
	p.InitStack(len(regStr))

	for _,val := range regStr {
		if val != cNeed {
			c.InStack(val)

			switch (string(val)) {
				case "[":
					cNeed = ']'
					p.InStack(cNeed)
				case "(":
					cNeed = ')'
					p.InStack(cNeed)
				default:
					fmt.Println("不匹配")
					//os.Exit(1)
			}
			p.StackTraverse()
			fmt.Println(string(cNeed))
			fmt.Println("--------")
		} else {
			c.DeStack()
			p.DeStack()

		}
	}

	if p.StackEmpty() {
		fmt.Println("符号符合规则")
	} else {
		fmt.Println("符号不符合规则")
	}
}
