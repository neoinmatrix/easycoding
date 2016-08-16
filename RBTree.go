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

//建红黑树
func (t *Tree) CreateRBTree(data int) {
	var insert *Node = &Node{
		Color:  RED,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Value:  data,
	}
	if t.Root == nil {
		insert.Color = BLACK
		t.Root = insert
		return
	}
	parent := t.Root
	var ipoint *Node = nil
	for parent != nil {
		ipoint = parent
		if data < ipoint.Value {
			parent = parent.Left
		} else if data > ipoint.Value {
			parent = parent.Right
		} else {
			return
		}
	}
	insert.Parent = ipoint

	if data < ipoint.Value {
		ipoint.Left = insert
	} else {
		ipoint.Right = insert
	}
	//t.FixTree(insert)
	//	if ipoint.Color == BLACK { //父黑 正常完成
	//		return
	//	}

	//	//新节点 为父节点 左节点
	//	if ipoint.Parent.Left == ipoint {
	//		//新节点 为父节点 右节点
	//		if towards == T_RIGHT {
	//			pnode := ipoint.Parent
	//			pnode.Left = t.RotateTree(ipoint, TR_LEFT)
	//			ipoint = pnode.Left
	//		}
	//		//新节点 为父节点 左节点
	//		if ipoint.Parent.Right == nil { // 叔父为黑色
	//			ipoint.Parent.Color = RED
	//			ipoint.Color = BLACK
	//			if ipoint.Parent.Parent != nil {
	//				pnode := ipoint.Parent.Parent
	//				pnode.Left = t.RotateTree(ipoint.Parent, TR_RIGHT)
	//				ipoint = pnode.Left
	//			} else {
	//				t.Root = t.RotateTree(ipoint.Parent, TR_RIGHT)
	//				ipoint = t.Root.Left
	//			}
	//		} else { // 叔父为红色
	//			ipoint.Color = BLACK
	//			ipoint.Parent.Color = RED
	//			ipoint.Parent.Right.Color = BLACK

	//		}
	//	}

	//	//新节点 为父节点 右节点
	//	if ipoint.Parent.Right == ipoint {
	//		if towards == T_LEFT {
	//			pnode := ipoint.Parent
	//			pnode.Right = t.RotateTree(ipoint, TR_RIGHT)
	//			ipoint = pnode.Right
	//		}
	//		if ipoint.Parent.Left == nil { // 叔父为黑色
	//			ipoint.Parent.Color = RED
	//			ipoint.Color = BLACK
	//			if ipoint.Parent.Parent != nil {
	//				pnode := ipoint.Parent.Parent
	//				pnode.Right = t.RotateTree(ipoint.Parent, TR_LEFT)
	//				ipoint = pnode.Right
	//			} else {
	//				t.Root = t.RotateTree(ipoint.Parent, TR_LEFT)
	//				ipoint = t.Root.Left
	//			}
	//		} else { // 叔父为红色
	//			ipoint.Color = BLACK
	//			ipoint.Parent.Color = RED
	//			ipoint.Parent.Left.Color = BLACK
	//		}
	//	}

	//	if ipoint.Parent == t.Root && ipoint.Parent.Color == RED {
	//		ipoint.Parent.Color = BLACK
	//	}
}

func (t *Tree) FixTree(insert *Node, towards int) {
	for insert.Parent != nil && insert.Parent.Color == RED { //需要调整的情况
		if insert.Parent.Parent.Left == insert.Parent { //为左支
			if insert.Parent.Parent.Right.Color == RED { //叔父节点红色
				insert = insert.Parent.Parent
				insert.Left.Color = BLACK
				insert.Right.Color = BLACK
				insert.Color = RED
			} else { //叔父黑色或空
				if insert.Parent.Right == insert {
					insert = insert.Parent
					t.RotateTree(insert, TR_LEFT)
				}
				insert.Parent.Parent.Color = RED
				insert.Parent.Color = BLACK
				t.RotateTree(insert.Parent.Parent, TR_RIGHT)
			}
		} else { //为右支
			if insert.Parent.Parent.Left.Color == RED { //叔父节点红色
				insert = insert.Parent.Parent
				insert.Right.Color = BLACK
				insert.Left.Color = BLACK
				insert.Color = RED
			} else { //叔父黑色或空
				if insert.Parent.Left == insert {
					insert = insert.Parent
					t.RotateTree(insert, TR_RIGHT)
				}
				insert.Parent.Parent.Color = RED
				insert.Parent.Color = BLACK
				t.RotateTree(insert.Parent.Parent, TR_LEFT)
			}
		}
	}
}

//旋转红黑树
func (t *Tree) RotateTree(node *Node, towards int) {
	if node == nil {
		return
	}
	var tmproot = node.Parent
	var whichchild
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

//删除红黑树
func (t *Tree) DeleteNode(data int) {
	if data == t.Root.Value {

	}

}

//主程序
func main() {
	var treedata = []int{
		5, 3, 1,
		//12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, // 14, 13, 10, 16, 6, 3, 8, 17,
		//8, 3, 15, 18, 17,
		//		8, 9, 10,
		//		8, 3, 9, 2, 5, 10, 4,
		//		9, 10,

	}
	//5, 10, 13, 8, 4, 2, 3, 9,
	var mytree = Tree{
		Root: nil,
		//		VisitType: 1,
		//		Height:    0,
	}

	for _, value := range treedata { //建立数
		mytree.CreateRBTree(value)
	}
	mytree.RotateTree(mytree.Root, TR_RIGHT)
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
	if n.Color == RED {
		fmt.Print("[R] ")
	} else {
		fmt.Print("[B] ")
	}

}
