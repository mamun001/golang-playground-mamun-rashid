
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


