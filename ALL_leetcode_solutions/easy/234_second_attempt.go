// easy 234 , no ELO, acc 55%
// first attempt = took too long (using slices)
// second attempt = worked , but needed a lot of help ; March 17

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


// This func is same as Leetcode 206
func reverse(LL *ListNode) *ListNode {
    prev := LL
    prev = nil
    cur := LL

    for cur != nil {
        save:= cur.Next
        cur.Next=prev
        prev = cur
        cur = save
    }
    return prev
}


func isPalindrome(head *ListNode) bool {

    // edge cases
    if head == nil || head.Next == nil {
        return true
    }

    // two pointers technique : slow and fast
    slow:=head
    fast:=head

    // FIND MIDDLE
    // if len of LL is even , that middle is considered the pointer to n/2+1
    // e.g. if 1,2,3,4 ==> pointer to 3 is the "middle"
    // note that we don't need check for fast.Next.Next to be nil
    // because if fast.Next is NOT nil, that will not break the code because fast.Next.Next is nil and NOT OOM
    for fast !=nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // now slow points to the middle

    // reverse the 2nd half
    reversed_2nd_half := reverse(slow)

    cur1:=head
    cur2:=reversed_2nd_half

    for cur2 != nil { // cur1 will go double the length, so we must stop at end of cur2
        if cur1.Val != cur2.Val {
            return false
        } else {
            cur2 = cur2.Next
            cur1 = cur1.Next
        }
    }

    

    return true

    
}
