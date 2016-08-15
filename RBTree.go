// Red and Black tree project main.go
package main

import (
	"fmt"
)

type Node struct {
	Color  int
	Left   *Node
	Right  *Node
	Parent *Node
	Value  int
}
type Tree struct {
	Root      *Node   //树
	VisitType int     //深度优先级
	Queue     []*Node //层次遍历队列
	Height    int     //树高
}

const TR_LEFT = 0  // tree rotate to left
const TR_RIGHT = 1 // tree roate to right
const RED = -1
const BLACK = 1
const T_LEFT = -1
const T_RIGHT = 1

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
func (t *Tree) CreateRBTree(node *Node, data int) (*Node, bool) {
	if node == nil {
		tmp := Node{
			Color:  RED,
			Left:   nil,
			Right:  nil,
			Parent: nil,
			Value:  data,
		}
		if t.Root == nil { // 根节点
			tmp.Color = BLACK
		}
		return &tmp, true
	}
	var change bool = false
	var towards int = T_LEFT
	if data < node.Value {
		if node.Left == nil {
			node.Left, change = t.CreateRBTree(node.Left, data)
			node.Left.Parent = node
		} else {
			node, change = t.CreateRBTree(node.Left, data)
		}
		towards = T_LEFT
	} else if data > node.Value {
		if node.Left == nil {
			node.Right, change = t.CreateRBTree(node.Right, data)
			node.Right.Parent = node
		} else {
			node, change = t.CreateRBTree(node.Right, data)
		}

		towards = T_RIGHT
	} else {
		return node, false
	}

	if node.Color == RED && node.Parent == nil { //根节点一定为黑色
		node.Color = BLACK
	}

	if change == true && node.Color == RED && node.Parent != nil && node.Parent.Left == node {
		//新节点 为父节点 右节点
		if towards == TR_RIGHT {
			node = t.RotateTree(node, TR_LEFT)
		}
		//新节点 为父节点 左节点
		if node.Parent.Right == nil { // 叔父为黑色
			node.Parent.Color = RED
			node.Color = BLACK
			node.Parent = t.RotateTree(node.Parent, TR_RIGHT)
		} else { // 叔父为红色
			node.Color = BLACK
			node.Parent.Color = RED
			node.Parent.Right.Color = BLACK
			if node.Parent == t.Root {
				node.Parent.Color = BLACK
			}
		}
	}

	//新节点 为父节点 左节点
	if change == true && node.Color == RED && node.Parent != nil && node.Parent.Right == node {
		if towards == TR_LEFT {
			node = t.RotateTree(node, TR_RIGHT)
		}
		if node.Parent.Left == nil { // 叔父为黑色
			node.Parent.Color = RED
			node.Color = BLACK
			node.Parent = t.RotateTree(node.Parent, TR_LEFT)
		} else { // 叔父为红色
			node.Color = BLACK
			node.Parent.Color = RED
			node.Parent.Left.Color = BLACK
			if node.Parent == t.Root {
				node.Parent.Color = BLACK
			}
		}
	}

	return node, false
}

//旋转平衡树
func (t *Tree) RotateTree(node *Node, towards int) *Node {
	if towards == TR_LEFT { //左旋
		rchild := node.Right
		rchild.Parent = node.Parent
		node.Parent = rchild
		node.Right = rchild.Left
		rchild.Left = node
		return rchild
	} else {
		lchild := node.Left
		lchild.Parent = node.Parent
		node.Parent = lchild
		node.Left = lchild.Right
		lchild.Right = node
		return lchild
	}
}

//主程序
func main() {
	var treedata = []int{
		//		8, 3, 10, 2, 5, 4,
		8, 3, 9, // 2, 1,
		//		8, 3, 9, 2, 5, 10, 4,
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
		mytree.Root, _ = mytree.CreateRBTree(mytree.Root, value)
	}
	//mytree.Root = mytree.RotateTree(mytree.Root, TR_LEFT)
	//深度遍历找树高
	//	mytree.DFS(mytree.Root, howdepth)
	//	fmt.Println(depth)
	//广度遍历整棵树
	mytree.BFS(mytree.Root, print)
	//	mytree.Root, _ = mytree.DeleteNode(mytree.Root, 2)
	//	mytree.BFS(mytree.Root, print)

}

//输出函数
func print(n *Node) {
	fmt.Print(n.Value, " ")
	fmt.Print("[", n.Color, "] ")
}
