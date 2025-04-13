// easy 83, Linkedlist
// Remove Duplicates from Sorted List
// ELO unknown
// acc rate: Only 54%
// first time: earlier
// 2nd time: March 17 without help
// 3rd time: March 18 without help

// 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

    dummy := &ListNode{-101,head}
    prev := dummy
    cur := head
    seen_map := make(map[int]int)

    
    for cur != nil {
        //fmt.Println("val",cur.Val)
        if seen_map[cur.Val] != 1 {
            //fmt.Println("val does not exist in map")
            seen_map[cur.Val] = 1
            //fmt.Println("map after inserting")
            //fmt.Println(seen_map)
            prev = prev.Next
        } else {
            //fmt.Println("val exists in map")
            prev.Next = cur.Next
        }
        cur = cur.Next
    }

    return head
    
}
