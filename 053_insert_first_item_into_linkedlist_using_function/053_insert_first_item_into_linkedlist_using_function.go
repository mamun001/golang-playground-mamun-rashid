package main

import (
	"fmt"
)

// Mamun Rashid
// Tested in go playground

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insert_first(LL *ListNode, Value int) *ListNode {
	
	LL = &ListNode {Val: Value, Next: nil}
	return LL
}

func main() {

	var foo_list *ListNode

	foo_list = insert_first(foo_list, 300)

	fmt.Println(foo_list.Val, foo_list.Next)

}

