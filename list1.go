package main

import (
	"fmt"
	"数据结构/list"
)

func main() {
	// 3 5 7 2 9 1 8
	var l = list.SortList{}
	l.InitList(10)
	fmt.Println("表长度:", l.ListLength())
	l.ListInsert(0, 3)
	l.ListInsert(1, 5)
	l.ListInsert(2, 7)
	l.ListInsert(3, 2)
	l.ListInsert(4, 9)
	l.ListInsert(5, 1)
	l.ListInsert(3, 8)
	l.ListTraverse()
	e,_ := l.GetElement(5)
	fmt.Println("获取指定元素:", e)
	fmt.Println("获取指定元素下标:",l.LocateEmelent(8))
	fmt.Println("表长度:", l.ListLength())
	l.ListDelete(4)
	fmt.Println("表长度:", l.ListLength())
	if !l.ListIsEmpty() {
		l.ClearList()
	}

}
