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

func initializeTree() Root {
	newTree := Root{nil}
	return newTree
}

func (r *Root) addNode(val int) {
	newNode := new(Node)
	newNode.Parent, newNode.LeftChild, newNode.RightChild = nil, nil, nil
	newNode.Value = val
	if r.Root == nil {
		r.Root = newNode
		return
	}
	adder := r.Root
	if adder.Value > val {
		r.addNode(adder.RightChild)
	} else if adder.Value <= val {

	}
	if adder == nil {
		adder = newNode
	}
}

func (r *Root) foundElement(val int) (bool, int) {
	traverser := r.Root
	for traverser != nil {
		if traverser.Value >= val {

		}
	}
}

func main() {
	fmt.Println("hello")
	tree := initializeTree()
	tree.addNode(5)
	fmt.Println(tree)
	tree.foundElement(5)
}
