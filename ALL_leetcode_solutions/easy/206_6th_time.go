// leetcode easy 206, part of neetcode 250
// Did this a sixth time on March 16 (20 days after fifth time)
// in the 21 days, I forgot about 20%
// feb 23, did 4th time
// feb 24, did a 5th time
// march 16, did a 6th time



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    prev := head // just to force of "prev" to be of type Listnode, because I am lazy!
    prev = nil
    cur := head

    for cur != nil {
        save := cur.Next
        cur.Next = prev

        prev = cur
        cur = save
    }

    return prev
    
}
