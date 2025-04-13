

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func len(head *ListNode) int {
   cur := head
   ans:=0
   for cur != nil {
      ans++
      cur = cur.Next
   }
   return ans
}

func powerof2(n int) int {
    ans:=1

    for i:=0;i<n;i++{
        ans=ans*2
    }
    return ans
}

func getDecimalValue(head *ListNode) int {

    //fmt.Println(len(head))
    //fmt.Println(powerof2(5))

    ans:=0
    cur := head
    i:=len(head)-1
    for cur != nil {
        ans=ans+cur.Val*powerof2(i)
        cur = cur.Next
        i--
    }


    return ans


    
}
