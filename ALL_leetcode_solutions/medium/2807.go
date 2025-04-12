// MEDIUM! 2807 , ELO 1279, LL
// Insert Greatest Common Divisors in Linked List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func gcf (n1 , n2 int) int {
    if n1 == 0 || n2 == 0 {
        return 0
    }
    max:=1
    for i:=2;i<=min(n1,n2);i++{
        if n1 % i == 0 && n2 % i == 0 {
            max = i
        }
    }
    return max
}


func insertGreatestCommonDivisors(head *ListNode) *ListNode {

    //fmt.Println(gcf(999,1000))

    //dummy:=&ListNode{}
    //cur := head
    prev := head
    cur := head.Next

    for cur != nil {
        //fmt.Println(gcf(prev.Val,cur.Val))
        newval := gcf(prev.Val,cur.Val)
        newnode:=&ListNode{newval,cur}
        prev.Next = newnode
        prev = cur
        cur = cur.Next
    }
    

    return head
    
}
