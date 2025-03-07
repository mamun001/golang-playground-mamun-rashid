// MEDIUM 19, Linked List
// Since max node size 30, performance was no concern
// so, bit of programming by adding nodes to a slice


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

    ar := []ListNode{}

    ptr := head 
    

    // convert list to an array
    for ptr != nil {
        //fmt.Println(ptr.Val)
        ar = append(ar,*ptr)
        ptr = ptr.Next
    }

    // figure out what n is from the start
    //fmt.Println(ar)
   from_the_head := len(ar)+1-n
   //fmt.Println("from_the_head",from_the_head)

   // special case 1
   if len(ar) == 1 {
    return nil
   }


   // special case 2
   ptr = head
   if from_the_head == 1 {
    // remove the head
    ptr = ptr.Next
    return ptr
   }

   // rest
   for i:=1;i<from_the_head-1;i++ {
       //fmt.Println(ptr.Val)
       ptr = ptr.Next
   }
   ptr.Next = ptr.Next.Next

    return head
    
}
