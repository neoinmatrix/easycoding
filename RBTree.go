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
		if data < parent.Value {
			parent = parent.Left
		} else if data > parent.Value {
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
	t.FixTree(insert)
	t.Root.Color = BLACK
}

//修正红黑树
func (t *Tree) FixTree(insert *Node) {
	for insert.Parent != nil && insert.Parent.Color == RED { //需要调整的情况 为红色一定有父节点
		if insert.Parent.Parent.Left == insert.Parent { //为左支
			if insert.Parent.Parent.Right == nil || insert.Parent.Parent.Right.Color == BLACK { //叔父黑色或空
				if insert.Parent.Right == insert {
					insert = insert.Parent
					t.RotateTree(insert, TR_LEFT)
				}
				insert.Parent.Parent.Color = RED
				insert.Parent.Color = BLACK
				t.RotateTree(insert.Parent.Parent, TR_RIGHT)
			} else { //叔父节点红色
				insert = insert.Parent.Parent
				insert.Left.Color = BLACK
				insert.Right.Color = BLACK
				insert.Color = RED
			}
		} else { //为右支
			if insert.Parent.Parent.Left == nil || insert.Parent.Parent.Left.Color == BLACK { //叔父黑色或空
				if insert.Parent.Left == insert {
					insert = insert.Parent
					t.RotateTree(insert, TR_RIGHT)
				}
				insert.Parent.Parent.Color = RED
				insert.Parent.Color = BLACK
				t.RotateTree(insert.Parent.Parent, TR_LEFT)
			} else { //叔父节点红色
				insert = insert.Parent.Parent
				insert.Right.Color = BLACK
				insert.Left.Color = BLACK
				insert.Color = RED
			}
		}
		//fmt.Println("")
	}
}

//旋转红黑树
func (t *Tree) RotateTree(node *Node, towards int) {
	if node == nil {
		return
	}
	var newnode *Node
	if towards == TR_LEFT { //左旋
		newnode = node.Right
		node.Right = newnode.Left
		if node.Right != nil {
			node.Right.Parent = node
		}
		newnode.Left = node
	} else { //右旋
		newnode = node.Left
		node.Left = newnode.Right
		if node.Left != nil {
			node.Left.Parent = node
		}
		newnode.Right = node
	}
	newnode.Parent = node.Parent
	node.Parent = newnode

	if newnode.Parent == nil { //原结点是根结点
		t.Root = newnode
	} else {
		if newnode.Parent.Left == node { //接在原结点的左支根
			newnode.Parent.Left = newnode
		} else { //接在原结点的右支根
			newnode.Parent.Right = newnode
		}
	}
}

//删除红黑树
func (t *Tree) DeleteNode(data int) {
	if t.Root == nil {
		return
	}
	//找到需要删除的点
	remove := t.Root
	for remove.Value != data {
		if remove.Value > data {
			remove = remove.Left
		} else {
			remove = remove.Right
		}
		if remove == nil {
			return
		}
	}
	//查找前驱或者后继
	var removenode *Node
	if remove.Left != nil { //前驱
		removenode = remove.Left
		for removenode.Right != nil {
			removenode = removenode.Right
		}
		remove.Value = removenode.Value
	} else if remove.Right != nil { //后继
		removenode = remove.Right
		for removenode.Left != nil {
			removenode = removenode.Left
		}
		remove.Value = removenode.Value
	} else {
		removenode = remove
	}
	if removenode.Parent == nil {
		t.Root = nil
		return
	}
	// 删除需要删除的点
	if removenode.Parent.Left == removenode {
		removenode.Parent.Left = removenode.Left
		if removenode.Left != nil {
			removenode.Left.Parent = removenode.Parent
		}
	} else {
		removenode.Parent.Right = removenode.Right
		if removenode.Left != nil {
			removenode.Right.Parent = removenode.Parent
		}
	}
	//调整树
	if removenode.Color == RED { //红色删除无碍
		return
	}
	t.FixTreeDelete(removenode)
}

//修正红黑树
func (t *Tree) FixTreeDelete(removenode *Node) {
	//新结点为红色 改为黑色 补充原来黑结点
	if removenode.Parent.Left == removenode.Left {
		if removenode.Left.Color == RED {
			removenode.Left.Color = BLACK
			return
		}
	} else {
		if removenode.Right.Color == RED {
			removenode.Right.Color = BLACK
			return
		}
	}

	for removenode.Color == BLACK && removenode.Parent != nil {
		if removenode.Parent.Left == removenode.Left { //左支 黑色结点一定存在兄弟结点
			uncle := removenode.Parent.Right
			if uncle.Parent.Color == BLACK && uncle.Color == BLACK &&
				(uncle.Left == nil || uncle.Left.Color == BLACK) &&
				(uncle.Right == nil || uncle.Right.Color == BLACK) { //周围结点全黑色
				uncle.Color = RED
				return
			}

			if uncle.Parent.Color == RED && uncle.Color == BLACK &&
				(uncle.Left == nil || uncle.Left.Color == BLACK) &&
				(uncle.Right == nil || uncle.Right.Color == BLACK) { //周围结点黑色 父红
				uncle.Color = RED
				uncle.Parent.Color = BLACK
				return
			}

			if uncle.Color == BLACK && uncle.Right != nil && uncle.Right.Color == RED { //周围结点黑色 堂兄右红
				tmpcolor := uncle.Color
				uncle.Color = uncle.Parent.Color
				uncle.Parent.Color = tmpcolor
				t.RotateTree(uncle.Parent, TR_LEFT)
				uncle.Right.Color = BLACK
				return
			}

			if uncle.Color == BLACK && uncle.Left != nil && uncle.Left.Color == RED { //周围结点黑色 堂兄右红
				t.RotateTree(uncle, TR_RIGHT)
				//uncle = uncle.Parent
				continue
			}

			if uncle.Color == RED { // 叔父节点为红色 周围全是黑色
				removenode.Parent.Color = RED
				removenode.Parent.Right.Color = BLACK
				t.RotateTree(removenode.Parent, TR_LEFT)
				continue
			}

		} else { //右支
			uncle := removenode.Parent.Left
			if uncle.Parent.Color == BLACK && uncle.Color == BLACK &&
				(uncle.Left == nil || uncle.Left.Color == BLACK) &&
				(uncle.Right == nil || uncle.Right.Color == BLACK) { //周围结点全黑色
				uncle.Color = RED
				return
			}

			if uncle.Parent.Color == RED && uncle.Color == BLACK &&
				(uncle.Left == nil || uncle.Left.Color == BLACK) &&
				(uncle.Right == nil || uncle.Right.Color == BLACK) { //周围结点黑色 父红
				uncle.Color = RED
				uncle.Parent.Color = BLACK
				return
			}

			if uncle.Color == BLACK && uncle.Left != nil && uncle.Left.Color == RED { //周围结点黑色 堂兄右红
				tmpcolor := uncle.Color
				uncle.Color = uncle.Parent.Color
				uncle.Parent.Color = tmpcolor
				t.RotateTree(uncle.Parent, TR_RIGHT)
				uncle.Left.Color = BLACK
				return
			}

			if uncle.Color == BLACK && uncle.Right != nil && uncle.Right.Color == RED { //周围结点黑色 堂兄右红
				t.RotateTree(uncle, TR_LEFT)
				//uncle = uncle.Parent
				continue
			}

			if uncle.Color == RED { // 叔父节点为红色 周围全是黑色
				removenode.Parent.Color = RED
				removenode.Parent.Left.Color = BLACK
				t.RotateTree(removenode.Parent, TR_LEFT)
				continue
			}
		}
	}
}

//主程序
func main() {
	var treedata = []int{
		12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13,
		10, 16, 6, 3, 8, 17,
		//		1, 2, 3, 5, 6, 7, 8,
	}

	var mytree = Tree{
		Root: nil,
		//		VisitType: 1,
		//		Height:    0,
	}

	for _, value := range treedata { //建立数
		mytree.CreateRBTree(value)
	}
	//mytree.RotateTree(mytree.Root, TR_LEFT)
	//mytree.Root = mytree.RotateTree(mytree.Root, TR_LEFT)
	//深度遍历找树高
	//	mytree.DFS(mytree.Root, howdepth)
	//	fmt.Println(depth)
	//广度遍历整棵树
	mytree.BFS(mytree.Root, print)
	mytree.DeleteNode(2)
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
