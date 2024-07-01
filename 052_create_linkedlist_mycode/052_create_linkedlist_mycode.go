

package main

import (
	"fmt"
)

// Mamun Rashid

func main() {

	// Definition for singly-linked list.
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	var L *ListNode

	L = &ListNode{Val: 100, Next: nil}

	fmt.Println(L.Val)

}

