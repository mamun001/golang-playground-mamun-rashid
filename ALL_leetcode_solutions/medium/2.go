// MEDIUM!! LINKED LIST!! problem 2
// no ELO
// acc rate 45%
// done March 17, first try, no help

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// get length of a LL
func len(head *ListNode) int {
     ans:=0
     cur:=head
     for cur != nil {
        ans++
        cur = cur.Next
     }
     return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    cur1 := l1
    cur2 := l2
    l3 := &ListNode{0,nil} // dummy node
    cur3 := l3

    
    //add up up to the point of the shorter LL
    carryover:=0
    for cur1 != nil && cur2 != nil {
        //fmt.Println(i)
        cur3.Val = cur1.Val + cur2.Val + carryover
        if cur3.Val > 9 {
            carryover = 1
            cur3.Val = cur3.Val % 10
        } else {
            carryover =0 
        }
        cur1 = cur1.Next
        cur2 = cur2.Next 
        cur3.Next = &ListNode{0,nil}
        cur3 = cur3.Next
    }

    // if list1 is longer, add the rest from list1
    for cur1 != nil {
        //fmt.Println(i)
        cur3.Val = cur1.Val + carryover
        if cur3.Val > 9 {
            carryover = 1
            cur3.Val = cur3.Val % 10
        } else {
            carryover =0 
        }
        cur1 = cur1.Next
        cur3.Next = &ListNode{0,nil}
        cur3 = cur3.Next
    }
    // if list2 is longer, add the rest from list2
    for cur2 != nil {
        //fmt.Println(i)
        cur3.Val = cur2.Val + carryover
        if cur3.Val > 9 {
            carryover = 1
            cur3.Val = cur3.Val % 10
        } else {
            carryover =0 
        }
        cur2 = cur2.Next
        cur3.Next = &ListNode{0,nil}
        cur3 = cur3.Next
    }

    // last carryover, if it exists
    if carryover == 1 {
        cur3.Val = 1
        cur3.Next = &ListNode{0,nil}
    }

    //fmt.Println(len(l1))
    //fmt.Println(len(l2))
    //fmt.Println(len(l3))

    // the way the alg works, I always end up with an extra node at the end
    // remove the last mode!
    cur3 = l3 
    for i:=1;i<len(l3)-1;i++{
        cur3 = cur3.Next
    }
    cur3.Next = nil

    return l3
    
}
