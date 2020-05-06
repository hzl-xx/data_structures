package Tree

import (
	"errors"
	"fmt"
)

/**
度：节点的子节点个数
二叉树所有度小于等于2
二叉树遍历分为：前序遍历、中序遍历、后续遍历  前中后顺序是相对于根来说的
前序遍历：根左右   中序遍历：左根右  后续遍历：左右根

树与数组之间的算法转换
[]int {3 58 26(9/0)7}  		父亲节点下标*2+1 该节点左
			3(0)
	5(1)			8(2)	父亲节点下标*2+2 该节点右
2(3)	6(4)	9(5)	7(6)
如果没有子节点 用0表示
 */

type Tree struct {
	m_pTree []int	// 树节点数组
	m_pSize int	// 大小
}

func (t *Tree) Init(size int, root int) {
	t.m_pSize = size
	t.m_pTree = make([]int, size, size)
	t.m_pTree[0] = root
}

func (t *Tree) DesTree() {
	t.m_pTree = []int{}
}

func (t *Tree) SearchNode(nodeIndex int) (int, error) {
	if nodeIndex < 0 || nodeIndex >= t.m_pSize {
		return 0, errors.New("非法的节点索引")
	}
	if nodeIndex == 0 {
		return 0, errors.New("无意义的节点下标")
	}
	return t.m_pTree[nodeIndex],nil
}

/**
nodeIndex  父节点下标
direction  插入的是左还是右	0 左	1 右
 */
func (t *Tree) AddNode(nodeIndex int, direction int, node int) (bool, error) {
	if nodeIndex < 0 || nodeIndex >= t.m_pSize {
		return false, errors.New("非法的节点索引")
	}
	//if nodeIndex == 0 {
	//	return false, errors.New("无意义的节点下标")
	//}

	if direction == 0 {
		if nodeIndex*2+1 >= t.m_pSize {
			return false, errors.New("超出树的最大容量")
		}
		if t.m_pTree[nodeIndex*2+1] != 0 {
			return false, errors.New("该节点已经存在数据")
		}
		t.m_pTree[nodeIndex*2+1] = node
	} else {
		if nodeIndex*2+2 >= t.m_pSize {
			return false, errors.New("超出树的最大容量")
		}
		if t.m_pTree[nodeIndex*2+2] != 0 {
			return false, errors.New("该节点已经存在数据")
		}
		t.m_pTree[nodeIndex*2+2] = node
	}
	return true, nil
}

func (t *Tree) DeleteNode(nodeIndex int) (int, bool, error) {
	if nodeIndex < 0 || nodeIndex >= t.m_pSize {
		return 0, false, errors.New("非法的节点索引")
	}
	if nodeIndex == 0 {
		return 0, false, errors.New("无意义的节点下标")
	}
	if t.m_pTree[nodeIndex] == 0 {
		return 0, false, errors.New("该节点不存在数据")
	}
	node := t.m_pTree[nodeIndex]
	t.m_pTree[nodeIndex] = 0
	return node, true, nil
}

func (t *Tree) TreeTraverse() {
	for _,val := range t.m_pTree {
		fmt.Println(val)
	}
}


/**
链表实现二叉树
 */
type Node struct {
	index int
	data int
	lp, rp, pp *Node
}


func (n *Node) Init() {
	n.index = 0
	n.data = 0
	n.lp, n.rp, n.pp = nil, nil, nil
}

func (n *Node) GetNode(index int, data int) *Node{
	n.index = index
	n.data = data
	n.lp, n.rp, n.pp = nil, nil, nil
	return n
}


func (n *Node) SearchNode(index int) *Node {
	if n.index == index {
		return n
	} else if n.lp != nil && n.lp.index == index {
		return n.lp
	} else if n.rp != nil && n.rp.index == index {
		return n.rp
	} else {
		return nil
	}
}

func (n *Node) DeleteNode() {
	if n.lp != nil {
		n.lp.DeleteNode()
	}

	if n.rp != nil {
		n.rp.DeleteNode()
	}

	if n.pp != nil {
		if n.pp.lp == n {
			n.pp.lp = nil
		}

		if n.pp.rp == n {
			n.pp.rp = nil
		}
	}
	n = nil
}

// 前序遍历
func (n *Node) BeforeTraversal() {
	fmt.Println(n.index, "-----------", n.data)
	if n.lp != nil {
		n.lp.BeforeTraversal()
	}
	if n.rp != nil {
		n.rp.BeforeTraversal()
	}
}

// 中序遍历
func (n *Node) InorderTraversal() {

	if n.lp != nil {
		n.lp.InorderTraversal()
	}
	fmt.Println(n.index, "-----------", n.data)
	if n.rp != nil {
		n.rp.InorderTraversal()
	}
}

// 后续遍历
func (n *Node) AfterTraversal() {
	fmt.Println(n.index, "-----------", n.data)
	if n.rp != nil {
		n.rp.AfterTraversal()
	}
	if n.lp != nil {
		n.lp.AfterTraversal()
	}

}

type Tree1 struct {
	node *Node
}

func (t Tree1) Init(node *Node) {
	t.node = node
}

func (t Tree1) DesTree() {
	t.node = nil
}

func (t Tree1) SearchNode(index int) *Node {
	return t.node.SearchNode(index)
}

func (t Tree1) AddNode(index int, dir int, node *Node) bool {
	temp := t.SearchNode(index)
	if temp == nil {
		return false
	}

	if dir == 0 {
		if temp.lp != nil {
			return false
		}
		temp.lp = node
	} else {
		if temp.rp != nil {
			return false
		}
		temp.rp = node
	}
	return true
}

func (t Tree1) DeleteNode(index int) bool {
	temp := t.SearchNode(index)
	if temp == nil {
		return false
	}

	temp.DeleteNode()
	return true
}

func (t Tree1) BeforeTraversal() {
	t.node.BeforeTraversal()
}

func (t Tree1) InorderTraversal() {
	t.node.InorderTraversal()
}

func (t Tree1) AfterTraversal() {
	t.node.AfterTraversal()
}





