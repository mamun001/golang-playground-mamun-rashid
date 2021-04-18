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

// I have chosen "value" at each NODE to be of type int32
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int32
}

// HERE "BinaryTree" is technically JUST THE root of bunch of "nodes"
type BinaryTree struct {
	root *BinaryNode
}

// So, is GOLANG SMART ENOUGH TO DISTINGUISH THE TWO INSERT FUNCTIONS, EVEN THOUGH ARE NAMED THE SAME????

// This is THE INSERT for the "tree" or "BinaryTree".
// This , in turn, calls the INSERT into "node"
func (t *BinaryTree) insert(value int32) *BinaryTree {

	// IF root is nil, that means we are below a leaf i.e. and empty location, not even a leaf
	if t.root == nil {
		t.root = &BinaryNode{value: value, left: nil, right: nil}
	} else {
		// we are NOT below a leaf, so insert here?
		t.root.insert(value)
	}
	return t
}

// INSERT AT NODE
// But, it is confirmed that we need both "insert" functions
func (n *BinaryNode) insert(value int32) {

	// insert based on value, keep going left or right until you find the right place for the value to be inserted

	if n == nil {
		// n == nil, means we are BELOW leaf , no need to insert at all?
		return
	} else if value <= n.value {
		// value_to_insert is LESS THAN OR EQUAL TO where we are

		if n.left == nil {
			// We are not below leaf, WE ARE AT A NODE with real value
			// IF left child is below-leaf (ie NIL) AND value is less than parent, insert here
			n.left = &BinaryNode{value: value, left: nil, right: nil}
		} else {
			// n.left is NOT NIL
			// That measn there are valid nodes to the left
			// We have to continute to compare (keep going down the left side
			n.left.insert(value)
		}
	} else {
		// value_to_insert is MORE THAN where we are

		if n.right == nil {
			// We are not below leaf, WE ARE AT A NODE with real value
			// IF right child is below-leaf (ie NIL) AND value is more than parent, insert here
			n.right = &BinaryNode{value: value, left: nil, right: nil}
		} else {
			// n.right is NOT NIL
			// That measn there are valid nodes to the right
			// We have to continute to compare (keep going down the right side
			n.right.insert(value)
		}
	}
}

// TRAVERSE AND PRINT THE WHOLE BINARY TREE
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {

	// Parameter Explanations:
	// Which tree to print: node (recursive is ok)
	// How much to leave on the left side of screen: ns
	// What to call "root" node: ch

	// This is "bottom" of recurison. This is where recursion bounces back to higher level
	if node == nil {
		// We are at BELOW-LEAF, ie void space, nothing to print
		return
	}

	// ADD "ns" count spaces at the beginning of each line
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	// In the following order, root gets printed first, then left, finish left, then right; Even while doing "Right half" , left comes first at the children level
	// Switching the order of the following also works as desired
	// Remember, regardless where we start "left" or "right", we still "finish children" before coming back up
	fmt.Fprintf(w, "%c:%v\n", ch, node.value) // PRINT "THIS" NODE's VALUE
	print(w, node.left, ns+2, 'L')            // Recursively print Left Side
	print(w, node.right, ns+2, 'R')           // Recursively print Right Side

}

// is_there function written by Mamun Rashid
// This function returns true if the given value exists in the tree and false if it does not
func is_there(value int32, thisnode *BinaryNode) bool {

	answer := false // default answer

	if thisnode == nil {
		// We are at BELOW-LEAF, ie void space, nothing found
		return false
	} else if thisnode.value == value {
		return true // found it!
	} else {
		// not found so far, now look at left children and below and look right children and below
		answer = is_there(value, thisnode.left) || is_there(value, thisnode.right)
	}

	return answer
}

// find_parent function written by Mamun Rashid
func find_parent(childvalue int32, thisnode *BinaryNode) int32 {

	var answer int32 = 0 // default answer

	if thisnode.left.value == childvalue {
		answer = thisnode.value
		return answer
	}

	if thisnode.right.value == childvalue {
		answer = thisnode.value
		return answer
	}

	return answer
}

// ENTRYPOINT
func main() {
	tree1 := &BinaryTree{}
	tree1.insert(100)
	tree1.insert(150)
	tree1.insert(50)
	tree1.insert(75)
	tree1.insert(125)
	tree1.insert(137)

	// Tree Looks Like this (if we insert 100,150,50,75,125,137 in this order)
	//         100
	//  50                 150
	//      75      125
	//                  137

	// which tree to print: tree1
	// How much to leave on the left side of screen: 10
	// What to call "root" node: 'M'
	fmt.Println("Print out of tree1")
	print(os.Stdout, tree1.root, 10, 'M')
	fmt.Println("  ----------------------- ")

	fmt.Println("6 Test cases of is_there function")
	fmt.Println(is_there(100, tree1.root))
	fmt.Println(is_there(137, tree1.root))
	fmt.Println(is_there(75, tree1.root))
	fmt.Println(is_there(1, tree1.root))
	fmt.Println(is_there(66, tree1.root))
	fmt.Println(is_there(300, tree1.root))
	fmt.Println("  ----------------------- ")
	
	
        fmt.Println("Test cases of find_parent function")
	fmt.Println(find_parent(150, tree1.root))

}
