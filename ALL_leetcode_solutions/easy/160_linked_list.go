// easy 160 ELO = unknown
// acc rate = 59%, so not easy!
// my alg = CPU > 5%, RAM > 59%

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

    ans := headA
    curA := headA
    curB := headB
    found := false

    if curA == curB {
        return curA
    }
    

    for curA != nil {
        //fmt.Println(" ")
        //fmt.Println(curA.Val)
        //curA = curA.Next
        curB = headB

        for curB != nil {
            //fmt.Println("     ", curB.Val)
            if curA == curB {
                //fmt.Println("hello")
                found = true
                return curA
            }
            curB = curB.Next
        }
        curA = curA.Next
    }

    if found == false {
        return nil
    }

    return ans
    
}

