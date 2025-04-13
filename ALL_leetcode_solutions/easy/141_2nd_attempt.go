// easy 141, no ELO
// part of neetcode 250
// first time: earlier. needed help
// second time: March 17. no help needed


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func hasCycle(head *ListNode) bool {

    if head == nil || head.Next == nil {
        return false
    }

    slow := head.Next
    fast := head.Next.Next

    for fast != nil && fast.Next != nil {
        if slow == fast {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next

    }

    return false
    
}
