

// Mamun Rashid's code

package main

import (
	"fmt"
)


// Definition for singly-linked list.
// I don't know where I got this, but this is the simplest to the point of memorizing
//
// & = MEMORY ADDRESS
// * in front of a type =  OPPOSITE OF & = value inside that memory address
// 
// example from stackoverflow
// myString := "Hi"
// fmt.Println(*&myString)  // prints "Hi"
//
// I should have named this LINKEDLIST instead of LISTNODE
//
type ListNode struct {
	Val  int
	Next *ListNode
}
// ListNode is a struct , *ListNode IS A POINTER TO another struct



// Takes in a LINKEDLIST and a Value
func insert_first(LL *ListNode, Value int) *ListNode {
        // LL is what is sent to us to this function
        // we are changing what is sent to us and RETURNING THE CHANGED ITEM on the fly
        // this assumes list is empty. if that is not case, we are overwriting stuff.
	LL = &ListNode{Val: Value, Next: nil}
	return LL
}


//
func insert_at_2nd(LL *ListNode, Value int) *ListNode {

        // changing what is sent to us ON THE FLY, not a good practice, i think
	LL.Next = &ListNode{Val: Value, Next: nil}
	//pointer := LL.Next
	//if pointer == nil {
	//LL.Next = &ListNode{Val: Value, Next: nil}
	//}

	//for pointer != nil {
	//x := 0
	//x = x+1 // do nothing
	//fmt.Println(pointer.Val)
	//pointer = pointer.Next
	//}
	//pointer = &ListNode{Val: Value, Next: nil}

	return LL
}


//  TRAVERSE
func traverse(head *ListNode) {
        // head is what is sent to us
        // if next is nil, that means we are the last item, print that and be done with it
	if head.Next == nil {
		fmt.Println(head.Val)
		return
        // this is NOT the last item , so print and keep going with head.next
	} else {
		fmt.Println(head.Val)
		traverse(head.Next)
	}
}


// LENGTH
func length(head *ListNode) int {
	if head.Next == nil {
		return 1
	} else {
		return length(head.Next) + 1
	}
}

func main() {

	var foo_list *ListNode

	foo_list = insert_first(foo_list, 300)
	foo_list = insert_at_2nd(foo_list, 500)
	foo_list = insert_at_2nd(foo_list, 700)
	foo_list = insert_at_2nd(foo_list, 900)

	fmt.Println(length(foo_list))

	//fmt.Println(foo_list.Val, foo_list.Next.Val)

	traverse(foo_list)

}

