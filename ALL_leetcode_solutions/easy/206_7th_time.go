// leetcode easy 206, part of neetcode 250
// Did this a seventh time on March 17 (1 day after sixth time)
// in 1 day, I forgot about 5%
// feb 23, did 4th time
// feb 24, did a 5th time
// march 16, did a 6th time
// march 17, did 7th time



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    prev:= head // force prev to be of type LL, lazy coding
    prev=nil
    cur:=head

    for cur!= nil {
        save:=cur.Next // because we will need this for cur at the end
        cur.Next = prev // "reverse"
        prev=cur // we have to do this BEFORE cur is changed
        cur=save // keep going forward
        
    }

    return prev
    
}
