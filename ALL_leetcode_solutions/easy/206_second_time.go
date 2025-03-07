// leetcode easy 206, part of neetcode 250
// Did this a second time, after a long time
// the four steps must be either memorized or practiced to death

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    prev := head // because Go complains about pre being untyped
    prev = nil
    curr := head
    saved_next := head // because Go complains about pre being untyped

    for curr != nil {
        saved_next = curr.Next
        curr.Next = prev
        prev = curr
        curr = saved_next     
    }

    return prev
    
}
