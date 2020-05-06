package main

import (
	_map "数据结构/map"
)

func main() {
	m := _map.CMap1{}
	m.Init(6)
	var n = _map.Node{}
	var n1 = _map.Node{}
	var n2 = _map.Node{}
	var n3 = _map.Node{}
	var n4 = _map.Node{}
	var n5 = _map.Node{}
	n.Init("A")
	m.AddNode(&n)
	n1.Init("B")
	m.AddNode(&n1)
	n2.Init("C")
	m.AddNode(&n2)
	n3.Init("D")
	m.AddNode(&n3)
	n4.Init("E")
	m.AddNode(&n4)
	n5.Init("F")
	m.AddNode(&n5)


	m.SetValToMatrixForUndirectedGraph(0, 1, 6)
	m.SetValToMatrixForUndirectedGraph(0, 4, 5)
	m.SetValToMatrixForUndirectedGraph(0, 5, 1)
	m.SetValToMatrixForUndirectedGraph(1, 2, 3)
	m.SetValToMatrixForUndirectedGraph(1, 5, 2)
	m.SetValToMatrixForUndirectedGraph(2, 5, 8)
	m.SetValToMatrixForUndirectedGraph(2, 3, 7)
	m.SetValToMatrixForUndirectedGraph(3, 5, 4)
	m.SetValToMatrixForUndirectedGraph(3, 4, 2)
	m.SetValToMatrixForUndirectedGraph(4, 5, 9)
	//m.PrintMatrix()
	//	//
	//	//m.ResetNode()
	//	//m.DepthFirstTraverse(0)
	//	//fmt.Println()
	//	//m.ResetNode()
	//	//m.BreadthFirstTraverse(0)
	m.PrimTree(0)
}
