// tree project main.go
package main

import (
	"fmt"
)

type Node struct {
	Color   int
	Left    *Node
	Right   *Node
	Value   int
	Depth   int
	Balance int
}
type Tree struct {
	Root      *Node   //树
	VisitType int     //深度优先级
	Queue     []*Node //层次遍历队列
	Height    int     //树高
}

const TR_LEFT = 0  // tree rotate to left
const TR_RIGHT = 1 // tree roate to right
const TB_LRE = 0
const TB_LEFT = -1
const TB_RIGHT = 1

// 建普通树
func (t *Tree) CreateCommon(node *Node, data int, depth int) *Node {
	if node == nil {
		tmp := Node{
			Color: 0,
			Left:  nil,
			Right: nil,
			Value: data,
			Depth: depth,
		}
		return &tmp
	}
	if data < node.Value {
		node.Left = t.CreateCommon(node.Left, data, depth+1)
	} else if data > node.Value {
		node.Right = t.CreateCommon(node.Right, data, depth+1)
	}
	return node
}

//深度优先历树
func (t *Tree) DFS(node *Node, f func(*Node)) {
	if node == nil {
		return
	}
	if t.VisitType == 1 {
		f(node)
		t.DFS(node.Left, f)
		t.DFS(node.Right, f)
	} else if t.VisitType == 2 {
		t.DFS(node.Left, f)
		f(node)
		t.DFS(node.Right, f)
	} else {
		t.DFS(node.Left, f)
		t.DFS(node.Right, f)
		f(node)
	}

}

//广度优先遍历
func (t *Tree) BFS(node *Node, f func(*Node)) {
	//队列接口 一出两入
	if node == nil {
		return
	}
	t.Queue = append(t.Queue, node)

	var depth uint = 0
	var now = 0
	var tmp *Node
	var end = false
	for len(t.Queue) > 0 && !end {
		tmp = t.Queue[0]
		t.Queue = t.Queue[1:]

		if tmp == nil {
			t.Queue = append(t.Queue, nil)
			t.Queue = append(t.Queue, nil)
			fmt.Print("* ")
		} else {
			f(tmp)
			if tmp.Left != nil {
				t.Queue = append(t.Queue, tmp.Left)
			} else {
				t.Queue = append(t.Queue, nil)
			}
			if tmp.Right != nil {
				t.Queue = append(t.Queue, tmp.Right)
			} else {
				t.Queue = append(t.Queue, nil)
			}
		}
		now++
		if now == (1 << depth) { // 回车换层次 输出占位符
			for _, n := range t.Queue {
				if n != nil { // 换层次时 判断队列是否全为空 到底
					end = false
					break
				}
				end = true
			}
			fmt.Print("\n")
			now = 0
			depth++
		}
	}
	fmt.Println("")
	t.Queue = nil

}

//建平衡二叉树
func (t *Tree) CreateBalance(node *Node, data int) (*Node, bool) {
	if node == nil {
		tmp := Node{
			Color:   0,
			Left:    nil,
			Right:   nil,
			Value:   data,
			Balance: TB_LRE,
		}
		return &tmp, true
	}
	var taller bool = false
	var towards int = TB_LRE
	if data < node.Value {
		node.Left, taller = t.CreateBalance(node.Left, data)
		towards = TB_LEFT
	} else if data > node.Value {
		node.Right, taller = t.CreateBalance(node.Right, data)
		towards = TB_RIGHT
	} else {
		return node, false
	}

	if taller == true && towards == TB_LEFT { //由原先情况是否调整
		if node.Balance == TB_LEFT {
			node, taller = t.BalanceTree(node, TB_LEFT) //left不平衡 需调整
		} else if node.Balance == TB_RIGHT {
			node.Balance = TB_LRE
			taller = false
		} else {
			node.Balance = TB_LEFT
			taller = true
		}
	}

	if taller == true && towards == TB_RIGHT {
		if node.Balance == TB_RIGHT {
			node, taller = t.BalanceTree(node, TB_RIGHT) //right不平衡 需调整
		} else if node.Balance == TB_LEFT {
			node.Balance = TB_LRE
			taller = false
		} else {
			node.Balance = TB_RIGHT
			taller = true
		}
	}
	return node, taller
}

//左右平衡调整
func (t *Tree) BalanceTree(node *Node, towards int) (*Node, bool) {
	if towards == TB_LEFT { //左调整
		var left = node.Left
		if left.Balance == TB_LEFT { //左左不平衡
			node.Balance = TB_LRE
			left.Balance = TB_LRE
			node = t.RotateTree(node, TR_RIGHT)
		} else if left.Balance == TB_RIGHT { //左右不平衡
			lr := left.Right
			if lr.Balance == TB_LEFT {
				node.Balance = TB_LRE
				left.Balance = TB_RIGHT
			} else if lr.Balance == TB_RIGHT {
				node.Balance = TB_LEFT
				left.Balance = TB_LRE
			} else {
				node.Balance = TB_LRE
				left.Balance = TB_LRE
			}
			node.Left = t.RotateTree(left, TR_LEFT)
			node = t.RotateTree(node, TR_RIGHT)
		}
	} else { //右调整
		var right = node.Right
		if right.Balance == TB_RIGHT {
			node.Balance = TB_LRE
			right.Balance = TB_LRE
			node = t.RotateTree(node, TR_LEFT)
		} else if right.Balance == TB_LEFT {
			rl := right.Left
			if rl.Balance == TB_LEFT {
				node.Balance = TB_RIGHT
				right.Balance = TB_LRE
			} else if rl.Balance == TB_RIGHT {
				node.Balance = TB_LRE
				right.Balance = TB_LEFT
			} else {
				node.Balance = TB_LRE
				right.Balance = TB_LRE
			}
			node.Right = t.RotateTree(right, TR_RIGHT)
			node = t.RotateTree(node, TR_LEFT)
		}
	}
	if node.Balance == TB_LRE { //调整后的根节点 和 是否树增高
		return node, false
	} else {
		return node, true
	}
}

//旋转平衡树
func (t *Tree) RotateTree(node *Node, towards int) *Node {
	if towards == TR_LEFT { //左旋
		rchild := node.Right
		node.Right = rchild.Left
		rchild.Left = node
		return rchild
	} else {
		lchild := node.Left
		node.Left = lchild.Right
		lchild.Right = node
		return lchild
	}
}

//删除结点
func (t *Tree) DeleteNode(node *Node, data int) (*Node, bool) {
	var smaller = false
	var towards = TB_LRE
	if node == nil {
		return nil, smaller
	} else if node.Value > data {
		node.Left, smaller = t.DeleteNode(node.Left, data)
		towards = TB_LEFT
	} else if node.Value < data {
		node.Right, smaller = t.DeleteNode(node.Right, data)
		towards = TB_RIGHT
	} else {
		if node.Left != nil && node.Right == nil {
			node.Value = node.Left.Value
			node.Right = node.Left.Right
			node.Left, smaller = t.DeleteNode(node.Left, node.Left.Value)
			if smaller {
				node.Balance = TB_LRE
			}
		} else if node.Right != nil && node.Left == nil {
			node.Value = node.Right.Value
			node.Left = node.Right.Left
			node.Right, smaller = t.DeleteNode(node.Right, node.Right.Value)
			if smaller {
				node.Balance = TB_LRE
			}
		} else if node.Left != nil && node.Right != nil {
			var tmp *Node
			if node.Balance == TB_LEFT {
				tmp = node.Left
				for tmp.Right != nil {
					tmp = tmp.Right
				}
				node.Value = tmp.Value
				node.Left, smaller = t.DeleteNode(node.Left, tmp.Value)
				if smaller {
					node.Balance = TB_LRE
				}
			} else if node.Balance == TB_RIGHT {
				tmp = node.Right
				for tmp.Left != nil {
					tmp = tmp.Left
				}
				node.Value = tmp.Value
				node.Right, smaller = t.DeleteNode(node.Right, tmp.Value)
				if smaller {
					node.Balance = TB_LRE
				}
			} else {
				tmp = node.Right
				for tmp.Left != nil {
					tmp = tmp.Left
				}
				node.Value = tmp.Value
				node.Right, smaller = t.DeleteNode(node.Right, tmp.Value)
				if smaller {
					node.Balance = TB_LEFT
				}
			}
			return node, smaller
		} else {
			return nil, true
		}
	}
	//平衡分支
	if smaller == true && towards == TB_LEFT {
		if node.Balance == TB_LEFT {
			node.Balance = TB_LRE
			smaller = true
		} else if node.Balance == TB_RIGHT {
			node, smaller = t.BalanceTree(node, TB_RIGHT)
		} else {
			node.Balance = TB_RIGHT
			smaller = false
		}
	}

	if smaller == true && towards == TB_RIGHT {
		if node.Balance == TB_RIGHT {
			node.Balance = TB_LRE
			smaller = true
		} else if node.Balance == TB_LEFT {
			node, smaller = t.BalanceTree(node, TB_LEFT)
		} else {
			node.Balance = TB_LEFT
			smaller = false
		}
	}
	return node, smaller

}

//主程序
func main() {
	var treedata = []int{
		//		8, 3, 10, 2, 5, 4,
		8, 3, 9, 2, 5, 10, 4,
		//9, 10,
		//9, 1, 5, 8, 3, 7, 6, 0, 2, 4,
		//		10, 16, 9, 18, 13, 14,
		//		10, 16, 9, 18, 13, 11,
		//		15, 16, 10, 9, 13, 14,
		//		15, 16, 10, 9, 13, 11,
		//		10, 16, 9, 13, 14,
		//		10, 16, 9, 13, 11,
		//		15, 16, 10, 13, 14,
		//		15, 16, 10, 13, 11,
		//		11, 15, 10,
		//		13, 10, 11,
	}
	//5, 10, 13, 8, 4, 2, 3, 9,
	var mytree = Tree{
		Root:      nil,
		VisitType: 1,
		Height:    0,
	}
	for _, value := range treedata { //建立数
		//		mytree.Root = mytree.CreateCommon(mytree.Root, v, 1)
		mytree.Root, _ = mytree.CreateBalance(mytree.Root, value)
	}
	//mytree.Root = mytree.RotateTree(mytree.Root, TR_LEFT)
	//深度遍历找树高
	//	mytree.DFS(mytree.Root, howdepth)
	//	fmt.Println(depth)
	//广度遍历整棵树
	mytree.BFS(mytree.Root, print)
	mytree.Root, _ = mytree.DeleteNode(mytree.Root, 2)
	mytree.BFS(mytree.Root, print)

}

//测试
func test() {
	//	var ns []*Node
	//	ns = append(ns, mytree.Root)
	//	ns = append(ns, nil)
	//	fmt.Println(ns)
}

//输出函数
func print(n *Node) {
	fmt.Print(n.Value, " ")
	fmt.Print("[", n.Balance, "] ")
}

//查找查找深度
//func (t *Tree) howdepth(n *Node) {
//	if n.Depth > depth {
//		depth = n.Depth
//	}
//}
