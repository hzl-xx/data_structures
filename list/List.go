package list

import (
	"errors"
	"fmt"
)

/**
线性表(链表)
线性表是n个数据元素的有限序列

		 ——	顺序表（数组）
		｜
		｜
线性表  -
		｜		 ——	静态链表
		｜		｜	单链表
		 —— 链表	-
				｜	循环链表
				 ——	双向链表

链表应用场景：通讯录，一元多项式

 */

/*
顺序表
3 5 7 2 9 1 8
前驱:指定元素前面的元素 后继：指定元素后面的元素
 */

type SortList struct {
	m_pList map[int]int
	m_iSize int		// 容量
	m_iLength int	// 长度
}

func (sl *SortList) InitList(size int) bool {
	sl.m_iSize = size
	sl.m_pList = make(map[int]int)
	sl.m_iLength = 0
	return true
}

func (sl *SortList) DesList() {
	sl.m_pList = nil
}

func (sl *SortList) ClearList() {
	sl.m_pList = map[int]int{}
	sl.m_iLength = 0
}

func (sl *SortList) ListIsEmpty() bool {
	if sl.m_iLength == 0 {

		return true
	}
	return false
}

func (sl *SortList) ListLength() int {
	return sl.m_iLength
}

func (sl *SortList) GetElement(index int) (int, error) {
	if index <0 || index >= sl.m_iSize {
		return 0, errors.New("错误的下标")
	}
	return sl.m_pList[index], nil
}

// 寻找第一个满足e的数据元素的位置
func (sl *SortList) LocateEmelent(e int) int {
	for key, val := range sl.m_pList {
		if val == e {
			return key
		}
	}
	return -1
}

// 获取指定元素的前驱
func (sl *SortList) PrvElement(currentElement int) int {
	key := sl.LocateEmelent(currentElement)
	if key != -1 {
		if key == 0 {
			return -1
		} else {
			return sl.m_pList[key-1];
		}
	}
	return -1
}

// 获取指定元素的后继
func (sl *SortList) NextElement(currentElement int) int {
	key := sl.LocateEmelent(currentElement)
	if key != -1 {
		if key == sl.m_iLength-1 {
			return -1
		} else {
			return sl.m_pList[key+1];
		}
	}
	return -1
}

// 在指定位置插入元素
func (sl *SortList) ListInsert(i int, e int) bool {
	if i < 0 || i > sl.m_iLength {
		return false
	}
	for k := sl.m_iLength-1; k >= i; k-- {
		sl.m_pList[k+1] = sl.m_pList[k]
	}
	sl.m_pList [i] = e
	sl.m_iLength ++
	return true
}

// 在指定位置删除元素
func (sl *SortList) ListDelete(i int) bool {
	if i < 0 || i >= sl.m_iLength {
		return false
	}
	for k := i+1; k < sl.m_iLength; k++ {
		sl.m_pList[k-1] = sl.m_pList[k]
	}
	delete(sl.m_pList, sl.m_iLength - 1)
	sl.m_iLength --
	return true
}

func (sl *SortList) ListTraverse() {
	for key,val := range sl.m_pList {
		fmt.Println(key, val)
	}
}





type Node struct {
	data int
	next *Node
}

func (n *Node) PrintNode() {
	fmt.Println(n.data)
}

func (n *Node) SetNode(data int) Node {
	return Node{
		data:data,
		next:nil,
	}
}

// 单链表
type SingleList struct {
	list *Node
	length int
}

func (sl *SingleList) InitList() {
	n := Node{
		data:0,
		next:nil,
	}
	sl.list = &n
	sl.length = 0
}

func (sl *SingleList) ClearList() {
	sl.InitList()
}

func (sl *SingleList) ListEmpty() bool {
	if sl.length == 0 {
		return true
	} else {
		return false
	}
}

func (sl *SingleList) ListLength() int {
	return sl.length
}

func (sl *SingleList) GetElement(index int) *Node {
	if index < 0 || index >= sl.length {
		return nil
	}

	currentNode := sl.list
	for i := 0 ; i <= index; i ++ {
		currentNode = currentNode.next
	}
	return currentNode
}

func (sl *SingleList) LocateElement(data int) int {
	currentNode := sl.list
	count := 0
	for currentNode.next != nil {
		currentNode = currentNode.next

		if currentNode.data == data {
			return count
		}
		count ++
	}
	return -1
}

func (sl *SingleList) PrvElement(node Node) *Node {
	currentNode := sl.list
	for currentNode.next != nil {
		tempNode := currentNode
		currentNode = currentNode.next
		if currentNode.data == node.data {
			if tempNode == sl.list {
				return nil
			}
			return tempNode
		}
	}
	return nil
}

func (sl *SingleList) NextElement(node Node) *Node {
	currentNode := sl.list
	for currentNode.next != nil {
		currentNode = currentNode.next
		if currentNode.data == node.data {
			return currentNode.next
		}
	}
	return nil
}

func (sl *SingleList) ListTraverse() {
	currentNode := sl.list
	for currentNode.next != nil {
		currentNode = currentNode.next
		currentNode.PrintNode()
	}
}

func (sl *SingleList) ListInsert(index int, node Node) bool {
	if index < 0 || index > sl.length {
		return false
	}
	currentNode := sl.list
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return false
	}
	node.next = currentNode.next
	currentNode.next = &node
	sl.length ++
	return true
}

func (sl *SingleList) ListDelete(index int) (bool, Node) {
	if index < 0 || index >= sl.length {
		return false, Node{}
	}
	currentNode := sl.list
	beforeNode := &Node{}
	for i := 0; i <= index ; i++ {
		beforeNode = currentNode
		currentNode = currentNode.next
	}
	beforeNode.next = currentNode.next
	resNode := currentNode
	currentNode = nil
	sl.length --
	return true, *resNode
}

func (sl *SingleList) ListInsertHead(node Node) {
	temp := sl.list.next
	n := Node{
		data: node.data,
		next: nil,
	}
	sl.list.next = &n
	n.next = temp
	sl.length ++
}

func (sl *SingleList) ListInsertTail(node Node) {
	currentNode := sl.list
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &node
	sl.length ++
}

// 静态链表
type Node1 struct {
	data int
	curr int
}

type StaticList struct {
	list []*Node1
	freeList []*Node1 // 备用链表，存放空data的下标
}

func (s *StaticList) Init() {
	n := Node1{
		data: nil,
		curr: nil,
	}
	s.list[0] = &n
}

// 创建示例
func (s *StaticList) Create(length int) {
	for i := 1; i < length; i++ {
		if i == length -1 {
			s.list[i] = &Node1{
				data: i,
				curr: nil,
			}
		} else {
			s.list[i] = &Node1{
				data: i,
				curr: i+1,
			}
		}

	}
}

// 静态链表插入操作
func (s *StaticList) Insert(index int, data int) {
	// 创建节点

	i := s.FreeCurr()
	s.list[i] = &Node1{
		data: data,
		curr: i+1,
	}
	s.list[i-1].curr = i
}

// 获取空下标
func (s *StaticList) FreeCurr() int {
	i := s.freeList[0].curr
	if s.freeList[i].curr != nil {
		s.freeList[0].curr = s.freeList[i].curr
	}
	return i
}

// 删除
func (s *StaticList) Delete(index int) {
	// 创建节点

	s.list[index] = nil
	s.list[index-1].curr = index+1
}