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


// insert at pos 1 regardless
func init_insert(LL *ListNode, Value int) *ListNode {

	LL = &ListNode{Val: Value, Next: nil}
	return LL
}

// get the LAST VALUE
func get_tail(head *ListNode) int {
	if head.Next == nil {
		return head.Val
	} else {
		return (get_tail(head.Next))
	}
}

//we need this because we need to save the original head
//This is common solution to a pattern that comes up in recursion, when the orignal arg value need to saved when recursion ends

func insert_at_tail_inner(original_head *ListNode, changing_head *ListNode, Value int) *ListNode {

	if changing_head.Next == nil {
		changing_head.Next = &ListNode{Val: Value, Next: nil}
		return original_head
	} else {
		return (insert_at_tail_inner(original_head, changing_head.Next, Value))
	}
}

//wrapper function that calls the inner function while passing the original head which remains unchanged
func insert_at_tail(head *ListNode, Value int) *ListNode {

	return insert_at_tail_inner(head, head, Value)
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

func length(head *ListNode) int {
	if head.Next == nil {
		return 1
	} else {
		return length(head.Next) + 1
	}
}

func main() {

	var foo_list *ListNode

	foo_list = init_insert(foo_list, 1)
	foo_list = insert_at_2nd(foo_list, 2)
	foo_list = insert_at_2nd(foo_list, 22)
	foo_list = insert_at_2nd(foo_list, 222)

	fmt.Println("--length--")
	fmt.Println(length(foo_list))

	fmt.Println("--traverse--")
	traverse(foo_list)

	fmt.Println("--tail--")
	fmt.Println(get_tail(foo_list))

	foo_list = insert_at_tail(foo_list, 999)

	fmt.Println("--traverse-after-insert-at-atil")
	traverse(foo_list)

	fmt.Println("--tail--")
	fmt.Println(get_tail(foo_list))

}
