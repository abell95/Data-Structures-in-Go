package main

import (
	"fmt"
)

//Node is main element of BST
type Node struct {
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
	Value      int
}

//Root points at first node
type Root struct {
	Root *Node
}

func initializeBSTree(val int) Root {
	newTree := Root{nil}
	firstNode := new(Node)
	firstNode.Value = val
	firstNode.Parent = nil
	firstNode.LeftChild, firstNode.RightChild = nil, nil
	newTree.Root = firstNode
	return newTree
}

// func (r *Root) addNode(val int, loc *Node) {
// 	newNode := new(Node)
// 	newNode.Parent, newNode.LeftChild, newNode.RightChild = nil, nil, nil
// 	newNode.Value = val
// 	if r.Root == nil {
// 		r.Root = newNode
// 		return
// 	}
// 	adder := r.Root
// 	if adder.Value > val {
// 		r.addNode(adder.RightChild, adder)
// 	} else if adder.Value <= val {

// 	}
// 	if adder == nil {
// 		adder = newNode
// 	}
// }

func (r *Root) foundElement(val int) (bool, int) {
	traverser := r.Root
	for traverser != nil {
		if traverser.Value >= val {

		}
	}
	return false, 0
}

func main() {
	fmt.Println("hello")
	tree := initializeBSTree(25)
	fmt.Println(tree.Root.Value)
	//tree.addNode(5)
	//fmt.Println(tree)
	//tree.foundElement(5)
}
