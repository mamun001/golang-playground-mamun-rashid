// easy 203, Remove Linked List Elements
// acc 51%
// part of neetcode youtube playlist
// done w/o help! took some debugging!

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

    ptr := head

    if ptr == nil {
        return ptr
    } else if ptr.Val == val {
        return removeElements(ptr.Next, val)
    } else if ptr.Next == nil {
        return ptr
    } else if ptr.Next != nil && ptr.Next.Val == val {
        ptr.Next = ptr.Next.Next
        removeElements(ptr,val)
    } else if ptr.Next != nil && ptr.Next.Val != val {
        removeElements(ptr.Next,val)
    }



    return head
    
}
