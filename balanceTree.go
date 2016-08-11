// tree project main.go
package main

import (
	"fmt"
)

type Node struct {
	Color int
	Left  *Node
	Right *Node
	Value int
	Depth int
}
type Tree struct {
	Root      *Node
	VisitType int
	Queue     []*Node
}

// 建树
func (t *Tree) Create(node *Node, data int, depth int) *Node {
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
		node.Left = t.Create(node.Left, data, depth+1)
	} else if data > node.Value {
		node.Right = t.Create(node.Right, data, depth+1)
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
	if node == nil {
		return
	}
	t.Queue = append(t.Queue, node)

	var depth uint = 0
	var now = 0
	var tmp *Node
	for len(t.Queue) > 0 {
		tmp = t.Queue[0]
		t.Queue = t.Queue[1:]

		if tmp == nil {
			//			t.Queue = append(t.Queue, nil)
			//			t.Queue = appe	nd(t.Queue, nil)
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
		if now == (1 << depth) {
			fmt.Print("\n")
			now = 0
			depth++
		}
	}
	fmt.Println("")

}

func print(n *Node) {
	fmt.Print(n.Value, " ")
}

func main() {
	var treedata = []int{
		5, 10, 13, 8, 4, 2, 3, 9,
	}

	var mytree = Tree{
		Root:      nil,
		VisitType: 1,
	}
	for _, v := range treedata {
		mytree.Root = mytree.Create(mytree.Root, v, 1)
	}
	//	var ns []*Node
	//	ns = append(ns, mytree.Root)
	//	ns = append(ns, nil)
	//	fmt.Println(ns)
	//mytree.DFS(mytree.Root, print)
	mytree.BFS(mytree.Root, print)

}
