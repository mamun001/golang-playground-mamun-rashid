// 3217 MEDIUM LL ELO 1342
// first time 3/29 : needed help

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func is_in (ar []int, n int) bool {
    
    for i:=0;i<len(ar);i++ {
        if n == ar[i] {
            return true
        }
    }
    return false
}


func modifiedList(nums []int, head *ListNode) *ListNode {

    map_found_already := make(map[int]bool)

    for i:=0;i<len(nums);i++ {
        map_found_already[nums[i]] = true
    }

    dummy := &ListNode{0,head}
    cur := dummy

    for cur.Next != nil {
        if map_found_already[cur.Next.Val] == true {
             cur.Next = cur.Next.Next
        } else if map_found_already[cur.Next.Val] == false {
             cur = cur.Next
        }
    }

    return dummy.Next
    
}
