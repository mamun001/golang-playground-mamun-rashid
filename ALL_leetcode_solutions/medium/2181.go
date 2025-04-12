// MEDIUM 2181, Linked List, ELO 1333

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {

    sum:=0
    cur := head.Next
    dummy:=&ListNode{-999999,nil}
    cur2:=dummy
    prev:=&ListNode{}

    for cur != nil {
        sum=sum+cur.Val
        if cur.Val == 0 {
            //fmt.Println(sum)
            cur2.Val = sum
            cur2.Next = &ListNode{}
            prev=cur2
            cur2=cur2.Next
            sum=0
            
        }
        cur = cur.Next
        
    }

    prev.Next=nil
    return dummy
    
}
