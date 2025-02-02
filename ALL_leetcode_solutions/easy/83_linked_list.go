// easy 83, Linkedlist
// Remove Duplicates from Sorted List
// ELO unknown
// acc rate: Only 54%

// 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

    cur := head
    //prev_val := -101
    
    for cur != nil && cur.Next != nil {
        //fmt.Println(cur.Val)
        if cur.Val == cur.Next.Val {
            //fmt.Println("this and next same")
            cur.Next = cur.Next.Next
        } else {
           cur = cur.Next
        }
        //deleteDuplicates(cur)
    }

    return head
    
}

