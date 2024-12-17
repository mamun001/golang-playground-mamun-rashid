

// leetcode #21 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func insert_at_tail_inner(original_head *ListNode, changing_head *ListNode, Value int) *ListNode {
    
    // This now works even if the list is empty!!
    
        if original_head == nil {
                original_head = &ListNode{Val: Value, Next: nil}
                return original_head
        }

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



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    var answer *ListNode
    
    ptr1 := l1
    ptr2 := l2
    
    for ptr1 !=nil || ptr2 != nil {
        
        
        if ptr1 == nil {
            answer=insert_at_tail(answer,ptr2.Val)
            fmt.Println(ptr2.Val)
            ptr2 = ptr2.Next
            continue
        }
        
        if ptr2 == nil {
            answer=insert_at_tail(answer,ptr1.Val)
            ptr1 = ptr1.Next
            continue
        }
        
        if ptr1.Val <= ptr2.Val {
            answer=insert_at_tail(answer,ptr1.Val)
            ptr1 = ptr1.Next
        } else {
            answer=insert_at_tail(answer,ptr2.Val)
            ptr2 = ptr2.Next
        }
    }
    
    
    return answer
    
}
