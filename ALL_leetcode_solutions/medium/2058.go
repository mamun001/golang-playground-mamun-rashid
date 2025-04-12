// 2058, ELO 1311
// find the minimum and maximum number of nodes between critical points

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {

    

    prev := head
    cur := head.Next
    criticals:=[]int{}
    index:=0

    for cur != nil && cur.Next != nil {
        if prev.Val < cur.Val && cur.Val > cur.Next.Val {
            //fmt.Println("max", cur.Val)
            criticals = append(criticals,index)
        } else if prev.Val > cur.Val && cur.Val < cur.Next.Val {
            //fmt.Println("min",cur.Val)
            criticals = append(criticals,index)                 
        }
        prev = prev.Next
        cur = cur.Next
        index++
    }

    if len(criticals) < 2 {
        return []int{-1,-1}
    }

    //fmt.Println(criticals)
    min:=999999999
    max:=criticals[len(criticals)-1] - criticals[0]

    for i:=1;i<len(criticals);i++ {
        if criticals[i]-criticals[i-1] < min {
            min = criticals[i]-criticals[i-1]
        }

    }

    
    return []int{min,max}
    
}
