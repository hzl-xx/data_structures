package main

import (
	"fmt"
	"数据结构/list"
)

func main() {
	var n list.Node
	sl := list.SingleList{}
	sl.InitList()

	n1 := n.SetNode(1)
	n2 := n.SetNode(2)
	n3 := n.SetNode(3)
	n4 := n.SetNode(4)
	n5 := n.SetNode(5)
	sl.ListInsertHead(n1)
	sl.ListInsertHead(n2)
	sl.ListInsertTail(n3)
	sl.ListInsertTail(n4)
	sl.ListInsert(4, n5)
	//sl.ListDelete(4)
	sl.ListTraverse()

	res := sl.GetElement(1)
	res.PrintNode()

	res = sl.PrvElement(n3)
	res.PrintNode()
	res = sl.NextElement(n3)
	res.PrintNode()
	fmt.Println(sl.LocateElement(5))
}
