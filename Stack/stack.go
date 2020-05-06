package Stack

import (
	"fmt"
	"reflect"
)

/**
栈 先进后出
 */
type Stack struct {
	list []interface{}	// 栈列表
	size int	// 容量
	top int		// 栈顶，栈中元素个数
}

// 初始化栈
func (s *Stack) InitStack(size int) {
	s.size = size
	s.top = 0
	s.list = []interface{}{}
}

// 回收栈空间内存
func (s *Stack) DesStack() {
	s.list = append(s.list[:0], s.list[len(s.list):]...)
}

// 判断是否为空
func (s *Stack) StackEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

// 判断是否已满
func (s *Stack) StackFull() bool {
	if s.top == s.size {
		return true
	}
	return false
}

// 清空栈
func (s *Stack) ClearStack() {
	s.list = append(s.list[:0], s.list[len(s.list):]...)
}

// 获取长度
func (s *Stack) StackLength() int {
	return s.top
}

// 入栈 栈顶上升
func (s *Stack) InStack(element interface{}) bool {
	if s.StackFull() {
		return false
	}
	s.list = append(s.list, element)
	s.top ++
	return true
}

// 出栈 栈顶下降
func (s *Stack) DeStack() bool {
	if s.StackEmpty() {
		return false
	}
	s.list = append(s.list[:s.top-1], s.list[s.top:]...)
	s.top --
	return true
}

// 遍历栈中的元素
func (s *Stack) StackTraverse() {
	for i := s.top - 1; i >= 0; i -- {
		fmt.Println(reflect.ValueOf(s.list[i]))
	}
}

// 获取栈顶
func (s *Stack) TopStack() interface{} {
	return s.list[s.top-1]
}




/**
两栈共享空间
 */
type Stack1 struct {
	list []int
	top1, top2 int
}

// t 区分插入哪个栈
func (s *Stack1) Insert(element int, t int) {
	if s.top1 + 1 == s.top2 {
		return
	}
	if t == 0 {

		s.list[s.top1+1] = element
		s.top1++
	} else {
		s.list[s.top2-1] = element
		s.top2--
	}
}


/**
栈的链式存储
 */

type StackNode struct {
	data int
	next *StackNode
}

type StackList struct {
	topNode *StackNode
	count int
}

func (sl *StackList) Insert(node *StackNode) {
	node.next = sl.topNode
	sl.topNode = node
	sl.count++
}

func (sl *StackList) Pop() {
	sl.topNode = sl.topNode.next
	sl.count --
}