// MEDIUM 2095, LL
// ELO 1324

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {

    if head == nil {
        return head
    }

    if head.Next == nil {
        return nil
    }

    prev := &ListNode{}
    prev = nil
    slow := head
    fast := head 

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = prev.Next.Next

    return head
    
}
