package main

import (
	"fmt"
	"数据结构/Queue"
)

/*
队列 先入先出
	分为：普通队列  环形队列

普通队列，前面走了之后后面不会上来，内存就空着，浪费了
环形队列：前面走了，后面来的往后排，排成一个圆，空间就不会浪费
环形队列：可以在固定大小的内存空间中反复使用

 */

func main() {
	var mq *Queue.MyQueue
	mq = &Queue.MyQueue{}
	mq.InitQueue(5)

	mq.InQueue(13)
	mq.InQueue(15)
	mq.InQueue(16)
	mq.InQueue(17)
	mq.InQueue(18)
	mq.QueueTraverse()

	fmt.Println("\n")
	_, elem := mq.DeQueue()
	fmt.Println(elem)
	fmt.Println("\n")
	_, elem = mq.DeQueue()
	mq.InQueue(30)
	_, elem = mq.DeQueue()
	mq.InQueue(40)
	_, elem = mq.DeQueue()
	_, elem = mq.DeQueue()
	mq.InQueue(50)
	mq.InQueue(60)
	_, elem = mq.DeQueue()
	fmt.Println(elem)

	fmt.Println("\n")
	mq.QueueTraverse()

	fmt.Println("\n")
	//mq.ClearQueue()
	//mq.InQueue(20)
	//mq.InQueue(300)
	//mq.QueueTraverse()
}
