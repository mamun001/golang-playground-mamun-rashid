package main

import (
	"fmt"
)

// Mamun Rashid
// Tested in Golang

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insert_first(LL *ListNode, Value int) *ListNode {

	LL = &ListNode{Val: Value, Next: nil}
	return LL
}

func insert_at_end(LL *ListNode, Value int) *ListNode {

	pointer := LL.Next
	if pointer == nil {
		LL.Next = &ListNode{Val: Value, Next: nil}
	}
	return LL
}

func main() {

	var foo_list *ListNode

	foo_list = insert_first(foo_list, 300)
	foo_list = insert_at_end(foo_list, 500)

	fmt.Println(foo_list.Val, foo_list.Next.Val)

}
