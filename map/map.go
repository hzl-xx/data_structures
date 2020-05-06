package _map

import "fmt"

/**
图：无向图 有向图
无向图：
	连通图 每一个顶点都有到其他顶点的路径
	完全图 每个顶点之间都有连线  边数=n*(n-1)/2
	生成树 仅有能够连接顶点的边组成  边数=n-1

图的存储结构：邻接矩阵、邻接表、十字链表、邻接多重表
 */


type Node struct {
	data string
	isVisited bool // 该节点是否被访问过
}

func (n *Node)Init(data string) {
	n.data = data
}

/**
         A
	   /   \
     B      D
   /  \    /  \
  C    F  G -  H
   \  /
	E

	A  B  C  D  E  F  G  H
A	   1     1
B   1     1        1
C      1        1
D   1                 1  1
E         1        1
F      1        1
G            1           1
H            1		  1
 */
type CMap struct {
	cap int // 图中最多可以容纳的顶点数
	nodeCount int // 已经添加的顶点（节点个数）
	nodeArray map[int]*Node // 用来存放顶点的数组
	matrix [][]int // 用来存放邻接矩阵
}

func (cm *CMap) Init(cap int) {
	cm.cap = cap
	cm.nodeCount = 0
	cm.nodeArray = map[int]*Node{}
	cm.matrix = make([][]int, cap, cap)
	for i := 0; i < cap; i ++ {
		cm.matrix[i] = make([]int, cap, cap)
	}
}

func (cm *CMap) AddNode(node *Node) bool {
	n := &Node{}
	n.data = node.data
	cm.nodeArray[cm.nodeCount] = n
	cm.nodeCount ++
	return true
}

// 重置顶点
func (cm *CMap) ResetNode()  {
	for _, val := range cm.nodeArray {
		val.isVisited = false
	}
}

// 为有向图设置邻接矩阵
func (cm *CMap) SetValToMatrixForDirectedGraph(row int, col int, val int) bool {
	if row < 0 || row > cm.cap {
		return false
	}
	if col < 0 || col > cm.cap {
		return false
	}
	cm.matrix[row][col] = val
	return true
}

// 为无向图设置邻接矩阵
func (cm *CMap) SetValToMatrixForUndirectedGraph(row int, col int, val int) bool {
	if row < 0 || row > cm.cap {
		return false
	}
	if col < 0 || col > cm.cap {
		return false
	}
	cm.matrix[row][col] = val
	cm.matrix[col][row] = val
	return true
}

// 打印矩阵
func (cm *CMap) PrintMatrix() {
	for _, val := range cm.matrix {
		for _, value := range val {
			fmt.Print(value, "   ")
		}
		fmt.Println("")
	}
}

// 从矩阵中获取值
func (cm *CMap) GetValueFromMatrix(row, col int) int {
	if row < 0 || row > cm.cap {
		return -1
	}
	if col < 0 || col > cm.cap {
		return -1
	}
	return cm.matrix[row][col]
}

// 深度优先搜索
func (cm *CMap) DepthFirstTraverse(index int) {
	val := 0
	fmt.Print(cm.nodeArray[index].data, "   ")
	cm.nodeArray[index].isVisited = true

	// 通过邻接矩阵判断是否与其他顶点有连接
	for i := 0; i < cm.cap; i++ {
		val = cm.GetValueFromMatrix(index, i)
		if val != 0 {	// 判断是否有弧连接连接其他顶点
			// 再判断该点是否被访问过
			if cm.nodeArray[i].isVisited {
				continue
			} else {
				cm.DepthFirstTraverse(i)
			}
		} else {	// 如果没有去向索引为i的顶点的弧，则循环继续
			continue
		}
	}
}

// 广度优先遍历
func (cm *CMap) BreadthFirstTraverse(index int) {
	fmt.Print(cm.nodeArray[index].data, "   ")
	cm.nodeArray[index].isVisited = true

	cm.BreadthFirstTraverseImpl(index)

}

func (cm *CMap) BreadthFirstTraverseImpl(preIndex int) {
	var (
		value int
		index int
	)
	if preIndex > 10 {
		return
	}
	for i := 0; i < preIndex+1; i++ {
		for j := 0; j < cm.cap; j++ {
			value = cm.GetValueFromMatrix(j, i)
			if value != 0 {
				if cm.nodeArray[j].isVisited {
					continue
				} else {
					 fmt.Print(cm.nodeArray[j].data, "   ")
					 cm.nodeArray[j].isVisited = true
					 index = j
				}
			}
		}
	}
	if index == 0 {
		return
	} else {
		cm.BreadthFirstTraverseImpl(index)
	}

}

/**

 获取最小生成树-普里姆算法
		A
    6/ 1| 5\
  B -2- F -9- E
  3\ 8/  4\ 2/
     C -7- D
A B C D E F
0 1 2 3 4 5

*/

// 定义边的属性
type Edge struct {
	                      // 边的权重
	nodeIndexA, nodeIndexB, weigthValue int
	isSelect bool // 是否被选择
}

func (e *Edge) Init(indexA, indexB, weight int) {
	e.nodeIndexA = indexA
	e.nodeIndexB = indexB
	e.weigthValue = weight
	e.isSelect = false
}

func (e *Edge) GetEdge(indexA, indexB, weight int) *Edge {
	var edge Edge
	edge.nodeIndexA = indexA
	edge.nodeIndexB = indexB
	edge.weigthValue = weight
	edge.isSelect = false
	return &edge
}



type CMap1 struct {
	CMap
	// 保存最小边
	edge []*Edge
}

func (cm *CMap1) Init(cap int) {
	cm.cap = cap
	cm.nodeCount = 0
	cm.nodeArray = map[int]*Node{}
	cm.matrix = make([][]int, cap, cap)
	for i := 0; i < cap; i ++ {
		cm.matrix[i] = make([]int, cap, cap)
	}
	cm.edge = make([]*Edge, cap-1, cap-1)
}

// 普里姆算法
func (cm *CMap1) PrimTree(index int) {
	var (
		value int
		edgeCount int
		nodeVec []int
		edge Edge
	)

	fmt.Println(cm.nodeArray[index].data)
	edgeVec := map[int]*Edge{}	// 存放备选边
	nodeVec = append(nodeVec, index)
	cm.nodeArray[index].isVisited = true
	// 获取顶点关联的所有边
	for edgeCount < cm.cap - 1 {
		temp := nodeVec[len(nodeVec)-1]
		for i := 0; i < cm.cap; i++ {
			value = cm.GetValueFromMatrix(temp, i)
			if value != 0 {
				if cm.nodeArray[i].isVisited {
					continue
				} else {

					edgeVec[i] = edge.GetEdge(temp, i, value)
				}
			}
		}
		// 从可选边集合中找出最小的边
		edgeIndex := cm.GetMinEdge(edgeVec)
		edgeVec[edgeIndex].isSelect = true

		fmt.Println(edgeVec[edgeIndex].nodeIndexA, "---", edgeVec[edgeIndex].nodeIndexB, "  ", edgeVec[edgeIndex].weigthValue)

		// 存放最小边
		cm.edge[edgeCount] = edgeVec[edgeIndex]
		edgeCount ++

		// 存放最小边的另一个点
		nextNodeIndex := edgeVec[edgeIndex].nodeIndexB
		nodeVec = append(nodeVec, nextNodeIndex)
		cm.nodeArray[nextNodeIndex].isVisited = true

		fmt.Println(cm.nodeArray[nextNodeIndex].data)
		// 获取连着边的顶点，从这个顶点开始获取最小边
	}


}

// 获取最小边
func (cm *CMap1) GetMinEdge(edgeVec map[int]*Edge) int {
	weight := 0
	index := 0
	for _,vals := range edgeVec {
		fmt.Println("(",vals.nodeIndexA, "--------", vals.nodeIndexB,"    ",vals.isSelect,")")
	}
	for key, val := range edgeVec {
		if !val.isSelect {
			weight = val.weigthValue
			index = key
			break
		}
	}

	if weight == 0 {
		return -1
	}

	for key1, value := range edgeVec {
		if !value.isSelect && weight > value.weigthValue {
			weight = value.weigthValue
			index = key1
		}
	}
	return index
}

//func (cm *CMap1) KruskalTree() {
//	value := 0
//	edgeCount := 0
//	nodeSet := [][]int{}
//	var (
//		edge Edge
//	)
//	k := 0
//	edgeVec := map[int]*Edge{}
//	for i := 0; i < cm.cap; i++ {
//		for j := i+1; j < cm.cap; j++ {
//			value = cm.GetValueFromMatrix(i, j)
//			if value != 0 {
//				edgeVec[k] = edge.GetEdge(i, j, value)
//				k++
//			}
//		}
//	}
//
//
//	for edgeCount < cm.cap - 1 {
//		// 找出最小边
//		minIndex := cm.GetMinEdge(edgeVec)
//		edgeVec[minIndex].isSelect = true
//
//		// 找出最小连接点
//		nodeAIndex := edgeVec[minIndex].nodeIndexA
//		nodeBIndex := edgeVec[minIndex].nodeIndexB
//
//		aIsInSet := false
//		bIsInSet := false
//
//		nodeAInSetLabel := -1
//		nodeBInSetLabel := -1
//
//		// 找出点所在的集合
//		for i := 0; i < len(nodeSet); i ++ {
//			aIsInSet = cm.IsInSet(nodeSet[i], nodeAIndex)
//			if aIsInSet {
//				nodeAInSetLabel = i
//			}
//		}
//
//		for i := 0; i < len(nodeSet); i ++ {
//			aIsInSet = cm.IsInSet(nodeSet[i], nodeBIndex)
//			if bIsInSet {
//				nodeBInSetLabel = i
//			}
//		}
//
//		if nodeAInSetLabel == -1 && nodeBInSetLabel == -1 {
//			vec := []int{}
//			vec = append(vec, nodeAIndex, nodeBIndex)
//			nodeSet = append(nodeSet, vec)
//		} else if nodeAInSetLabel == -1 && nodeBInSetLabel != -1 {
//			nodeSet[nodeBInSetLabel] = append(nodeSet[nodeBInSetLabel], nodeAIndex)
//		} else if nodeBInSetLabel == -1 && nodeAInSetLabel != -1 {
//			nodeSet[nodeAInSetLabel] = append(nodeSet[nodeAInSetLabel], nodeBIndex)
//		} else if nodeBInSetLabel != -1 && nodeAInSetLabel != -1 &&  nodeAInSetLabel != nodeBInSetLabel{
//			nodeSet[nodeAInSetLabel] = append(nodeSet[nodeAInSetLabel], nodeBIndex)
//			cm.MergeNodeSet(nodeSet[nodeAInSetLabel], nodeSet[nodeBInSetLabel])
//			for k := nodeBInSetLabel; k < len(nodeSet) - 1; k++ {
//				nodeSet
//			}
//		} else if nodeBInSetLabel != -1 && nodeAInSetLabel != -1 &&  nodeAInSetLabel != nodeBInSetLabel {
//			continue
//		}
//
//		cm.edge[edgeCount] = edgeVec[minIndex]
//		edgeCount
//	}
//}