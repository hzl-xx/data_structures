package main

import (
	"fmt"
	"数据结构/Tree"
)

func main() {
	var t = Tree.Tree{}
	t.Init(9, 3)
	t.AddNode(0, 0, 5)
	t.AddNode(0, 1, 6)

	t.AddNode(1, 0, 4)
	t.AddNode(1, 1, 7)

	t.AddNode(2, 0, 9)
	t.TreeTraverse()
	node,_ := t.SearchNode(2)
	fmt.Println("\n")
	fmt.Println(node)
	fmt.Println("\n")
	node, _,_ = t.DeleteNode(5)
	fmt.Println(node)
	t.TreeTraverse()
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	t.DesTree()
}
