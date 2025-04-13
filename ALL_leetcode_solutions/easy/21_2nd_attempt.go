// easy 21, ELO unknown
// merge two sorted lists

// did first time earlier with help
// did 2nd time March 17 again with lots of help


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

    // edge cases
    if list1 == nil && list2 == nil {
        return list1
    }

    if list1 == nil && list2 != nil {
        return list2
    }

    if list1 != nil && list2 == nil {
        return list1
    }


    // using the dummy node makes life easire in so many LL problems

    dummy_head := &ListNode{-101,nil} 

    // cur3 is the pointer in the answer array
    // starts with the dummy node
    // this alos means that , before retrning, we have to cut off the dummy node
    cur3 := dummy_head


    // pointers within list1 and list2 as we move through
    cur1 := list1
    cur2 := list2


    // as long BOTH lists are going
    for cur1 != nil && cur2 != nil {
        if cur1.Val <=cur2.Val {
            cur3.Next=cur1 // next node of the answer is the node with the smaller value
            cur1 = cur1.Next
        } else {
            cur3.Next=cur2 // next node of the answer is the node with the smaller value
            cur2 = cur2.Next
        }
        cur3 = cur3.Next // We must move the "answer" LL pointer to next , so that next smallest node can be added there
    }


    // whichever list (list1/list2) is longer, we append the rest of it at the tail end of the "answer" LL
    if cur1 != nil {
        cur3.Next = cur1
    }

    if cur2 != nil {
        cur3.Next = cur2
    }



    return dummy_head.Next // Note we are cutting off the dummy node
    
}
