// easy 876 ELO unknown
// Middle of the Linked List
// acc rate 79.9

// my alg: CPU > 100, RAM > 70%
// Jan 29,2025


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

    cur := head
    ans := head
    L := -1
    M := -1

    for cur != nil {
        //fmt.Println(cur.Val)
        L++

        cur = cur.Next
    }
    //fmt.Println(L)
    if L % 2 == 0 {
        M = L/2 // 4th node
    } else {
        M = (L+1) / 2 // if 5, then 2 = 3rd node
    }
    fmt.Println(M)

    i:=0
    cur = head
    for cur != nil {
        //fmt.Println(i,cur.Val)
        if i == M {
            ans = cur
            break
        } else {
            cur = cur.Next
            i++
        }
    }


    return ans
    
}

