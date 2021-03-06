package main

import (
	"fmt"
)

// Mamun Rashid
// Tested at Go Playground

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insert_first(LL *ListNode, Value int) *ListNode {

	LL = &ListNode{Val: Value, Next: nil}
	return LL
}

func insert_at_2nd(LL *ListNode, Value int) *ListNode {

	LL.Next = &ListNode{Val: Value, Next: nil}
	return LL
}

func traverse(head *ListNode) {
	if head.Next == nil {
		fmt.Println(head.Val)
		return
	} else {
		fmt.Println(head.Val)
		traverse(head.Next)
	}
}

func main() {

	var foo_list *ListNode

	foo_list = insert_first(foo_list, 300)
	foo_list = insert_at_2nd(foo_list, 500)
	foo_list = insert_at_2nd(foo_list, 700)
	foo_list = insert_at_2nd(foo_list, 900)

	//fmt.Println(foo_list.Val, foo_list.Next.Val)

	traverse(foo_list)

}
