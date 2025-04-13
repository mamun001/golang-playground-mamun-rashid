// easy 83, Linkedlist
// Remove Duplicates from Sorted List
// ELO unknown
// acc rate: Only 54%
// first time: earlier
// 2nd time: March 17 without help

// 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

    map1 := make(map[int]int)

    prev:=head
    prev=nil
    cur := head

    for cur != nil {
        fmt.Println(cur.Val)
        if map1[cur.Val] == 1 {
           prev.Next = cur.Next
           map1[cur.Val] = 1
           cur = cur.Next
        } else {
           map1[cur.Val] = 1
           cur = cur.Next
           if prev == nil {
              prev = head 
           } else {
              prev = prev.Next
           }
        }
    }

    //fmt.Println(map1)

    return head
    
}
