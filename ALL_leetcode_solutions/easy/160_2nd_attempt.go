// easy 160 ELO = unknown
// acc rate = 59%, so not easy!
// did this first time = earlier
// did this 2nd time = March 17

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func getIntersectionNode(headA, headB *ListNode) *ListNode {

    map1:=make(map[*ListNode]int) 

    cur:=headA
    for cur != nil {
        map1[cur]=1
        cur=cur.Next
    }
    //fmt.Println(map1)
    cur=headB
    for cur != nil {
        _,exists := map1[cur]
        if exists {
            return cur
        } else {
            cur=cur.Next
        }
    }
     
    return nil
    
}
