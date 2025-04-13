// easy 203, Remove Linked List Elements
// acc 51%
// part of neetcode youtube playlist
// first time: done w/o help! took some debugging!
// second time (March 17) : had to get a LOT of help

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

    if head == nil {
        return head
    }


    dummy := &ListNode{51,head}
    prev := dummy
    cur := head

    for cur != nil {
        if cur.Val == val {
               prev.Next = cur.Next 
        } else {        
          prev = cur
        }
        cur = cur.Next
    }

    
    return dummy.Next
    
}
