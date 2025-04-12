// MEDIUM 2130 maximum twin sum, ELO 1318

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {

    ar:=[]int{}

    cur:=head
    
    for cur !=nil  {
        ar = append(ar,cur.Val)
        cur = cur.Next
    }

    l:=0
    r:=len(ar)-1
    max:=-99999
    for l<r {
        if ar[l]+ar[r] > max {
            max=ar[l]+ar[r]
        }
        l++
        r--
    }

    //fmt.Println(ar)
    return max
    
}
