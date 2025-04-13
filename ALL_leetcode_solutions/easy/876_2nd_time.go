// easy 876 ELO unknown
// Middle of the Linked List
// acc rate 79.9

// first time, complicated algorithm without help
// my alg: CPU > 100, RAM > 70%
// Jan 29

// second time, much better algorithm!  without help
// March 16


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

    slow := head
    fast := head

    for fast != nil && fast.Next != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }


    return slow
    
}
