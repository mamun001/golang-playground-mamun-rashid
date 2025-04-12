// 1721, MEDIUM , LL, 1386
// swapping nodes in a Linked List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func len(h *ListNode) int {
   cur := h
   ans := 0
   for cur != nil {
      ans++
      cur = cur.Next
   }  
   return ans
}


func swapNodes(head *ListNode, k int) *ListNode {

    //fmt.Println(len(head))
    l := len(head)
    
    if k > l/2 {
        k = l - k + 1
    }
    rev_k := l - k + 1
    //fmt.Println(k,rev_k)

    dummy := &ListNode{-9999,nil}

    cur := head 
    i:=1
    for cur != nil {
        if i == k  {
            //fmt.Println(cur.Val)
            dummy = cur
            dummy.Val = cur.Val
        }
        if i == rev_k  {
            //fmt.Println(cur.Val)
            dummy.Val , cur.Val = cur.Val, dummy.Val
            return head
        }
        cur = cur.Next
        i++
    }

    return head
    
}
