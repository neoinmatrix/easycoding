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
}
type Tree struct {
	Root      *Node
	VisitType int
}

// 建树
func (t *Tree) Create(node *Node, data int) {
	tmp := Node{
		Color: 0,
		Left:  nil,
		Right: nil,
		Value: data,
	}
	if t.Root == nil {
		t.Root = &tmp
		if t.VisitType == 0 {
			t.VisitType = 1
		}
		return
	}
	if data < node.Value {
		if node.Left == nil {
			node.Left = &tmp
			return
		} else {
			t.Create(node.Left, data)
		}
	} else if data > node.Value {
		if node.Right == nil {
			node.Right = &tmp
			return
		} else {
			t.Create(node.Right, data)
		}
	} else { // data= t.value
		return
	}

}

// 遍历树
func (t *Tree) Visit(node *Node, f func(*Node)) {
	if node == nil {
		return
	}
	if t.VisitType == 1 {
		f(node)
		t.Visit(node.Left, f)
		t.Visit(node.Right, f)
	} else if t.VisitType == 2 {
		t.Visit(node.Left, f)
		f(node)
		t.Visit(node.Right, f)
	} else {
		t.Visit(node.Left, f)
		t.Visit(node.Right, f)
		f(node)
	}

}

func print(n *Node) {
	fmt.Println(n.Value)
}

func main() {
	var treedata = []int{
		5, 10, 13, 8, 4, 2, 3, 9,
	}

	var mytree = Tree{
		Root: nil,
	}
	for _, v := range treedata {
		mytree.Create(mytree.Root, v)
	}
	mytree.VisitType = 3
	mytree.Visit(mytree.Root, print)

}
