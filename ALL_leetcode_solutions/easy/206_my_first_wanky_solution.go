
// Leetcode easy: reverse a Linked List
// 206
// This was my first solution
// BUT, there is MUCH better solution that exists that's only about 8 lines

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Convert a Linked List to a Stack in the form of a slice
func convert_to_slice(head *ListNode) []int{
	if head == nil {
        mySlice1 := make([]int, 0) // emptyslice
		return mySlice1
	} else {
        // append Val to what we get back from rest(recursion)
        return append(convert_to_slice(head.Next), head.Val)
	}
}

// Because of the structure of our overall code, first insert into a Linked List has a special case function
func init_insert(LL *ListNode, Value int) *ListNode {
	LL = &ListNode{Val: Value, Next: nil}
	return LL
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
    //a := insert_at_tail_inner(head, head, Value)
    //fmt.Println("traverse--a")
    //traverse(a)
	return insert_at_tail_inner(head, head, Value)
}



func reverseList(head *ListNode) *ListNode {  
    var answer *ListNode = nil
    
    // stack1 is a basically a STACK in form of a slice: reverse-order of what's in Linkedin list
    // if Linkedin List is 1,2,3,4,5 , then
    // stack1 is 5,4,3,2,1
    stack1:= make([]int,0)
    stack1 = convert_to_slice(head)
    //fmt.Println("--stack1--content")
    //fmt.Printf("%v", stack)
    
    
    // So, now we un-pop at a time and insert into into new Linkedlist called answer
    //  SO that answer will 5,4,3,2,1 just like the stack (5 is at the stop of the stack)
    
    // because of our code, first insert has to be called with different function
    if len(stack1) > 0 {
       answer=init_insert(answer,stack1[0])
    }
    // now we POP and FILL answer Linkedinlist
    for i:=1; i < len(stack1); i++ {
        answer=insert_at_tail(answer,stack1[i])
    }
    
    return answer
    
}
