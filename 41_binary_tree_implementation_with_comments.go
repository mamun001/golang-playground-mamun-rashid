package main

// Tested at Go Playground
// This is modified from https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
// Change names of variable and type
// Added comments so that code makes sense.

// io and os are only needed for printing

// TODO
// More comments
// More functions to add so that I get confortable using the code to build more things (e.g. find a value in a tree)

import (
	"fmt"
	"io"
	"os"
)

// I have chose "value" at each NODE to be of type int32
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int32
}

// HERE "BinaryTree" is technically JUST THE root of bunch of "nodes"
type BinaryTree struct {
	root *BinaryNode
}

// insert at "BinaryTree" i.e. at "root" ??
func (t *BinaryTree) insert(value int32) *BinaryTree {

	// IF root is nil, that means we are below a leaf i.e. and emty location, not even a leaf
	if t.root == nil {
		t.root = &BinaryNode{value: value, left: nil, right: nil}
	} else {
		// we are NOT below a leaf, so insert here?
		t.root.insert(value)
	}
	return t
}

// INSERT at "node"
func (n *BinaryNode) insert(value int32) {

	// insert based on value, keep going lef or right until you find the right place for the value to be inserted

	// n == nil, means we are BELOW leaf, go back up.
	if n == nil {
		return
	} else if value <= n.value {
		// We are not below leaf, we are at a node with real value
		// IF left child is below-leaf (ie empty) AND value is less than parent, insert here
		if n.left == nil {
			n.left = &BinaryNode{value: value, left: nil, right: nil}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{value: value, left: nil, right: nil}
		} else {
			n.right.insert(value)
		}
	}
}

// Traverse and print the whole binary tree
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	fmt.Println("Hello, playground")

	tree := &BinaryTree{}
	tree.insert(108)
	tree.insert(95)
	tree.insert(120)
	tree.insert(98)
	tree.insert(110)

	print(os.Stdout, tree.root, 0, 'M')
}

