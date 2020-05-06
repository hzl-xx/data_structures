package Queue

import (
	"fmt"
)

/*
队列 先入先出
	分为：普通队列  环形队列

普通队列，前面走了之后后面不会上来，内存就空着，浪费了
环形队列：前面走了，后面来的往后排，排成一个圆，空间就不会浪费
环形队列：可以在固定大小的内存空间中反复使用

*/

// 实现环形队列
type MyQueue struct {
	capacity  int		// 容量
	queueList map[interface{}]interface{}	// 队列数组
	amount 	  int 		// 元素个数
	head 	  int 		//队头
	foot      int		// 队尾
}

// 环形队列 如果只有一个元素q[0] 既是队头，也是队尾

// 创建队列
func (mq *MyQueue) InitQueue(queueCapcity int) {
	mq.capacity = queueCapcity
	mq.ClearQueue()
	mq.queueList = map[interface{}]interface{}{}
}

// 销毁队列
func (mq *MyQueue) DesQueue() {
	mq.queueList = map[interface{}]interface{}{}
}

//清空队列
func (mq *MyQueue) ClearQueue() {
	mq.head = 0
	mq.foot = 0
	mq.amount = 0
}

//判断当前是否为空
func (mq *MyQueue) QueueEmpty() bool {
	if len(mq.queueList) == 0 {
		return true
	}
	return false
}

// 判断队列是否已满
func (mq *MyQueue) QueueFull() bool {
	if len(mq.queueList) == mq.capacity {
		return true
	}
	return false
}

//获取队列长度
func (mq *MyQueue) QueueLength() int {
	return len(mq.queueList)
}

// 入队
func (mq *MyQueue) InQueue(element int) bool {
	if mq.QueueFull() {
		return false
	}
	// 队尾等于 队尾模于容量
	mq.foot = mq.foot % mq.capacity
	mq.queueList[mq.foot] = element
	mq.foot ++
	mq.amount ++
	return true
}

// 出队
func (mq *MyQueue) DeQueue() (bool, interface{}) {
	if mq.QueueEmpty() {
		return false, nil
	}
	element := mq.queueList[mq.head]

	delete(mq.queueList, mq.head)
	mq.head ++
	mq.head = mq.head % mq.capacity
	mq.amount --
	return true, element
}

// 遍历队列
func (mq *MyQueue) QueueTraverse() {
	fmt.Println(mq.queueList, mq.head, mq.foot)
	for i := mq.head; i < mq.amount + mq.head; i++ {
		fmt.Println(i, mq.queueList[i % mq.capacity])
	}
}


/**
链式队列
 */
type QElement int
type QNode struct {
	data QElement
	next *QNode
}

type QList struct {
	// front.next 指向队头，rear指向队尾
	front, rear *QNode
}

func (q *QNode) Insert(ql *QList, e QElement) {
	q.data = e
	q.next = nil

	ql.rear.next = q
	ql.rear = q
}

func (q *QNode) Pop(ql *QList) {

	ql.front.next = ql.front.next.next

}

