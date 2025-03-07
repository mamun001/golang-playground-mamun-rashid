// easy 141, no ELO
// part of neetcode 250

// Linked List
// You can have a map of nodes!

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func hasCycle(head *ListNode) bool {

    map1 := make(map[ListNode] int)


    ptr := head

    //map1[*ptr] = 1

    //fmt.Println(map1)
    //i := 0

    for ptr != nil  {
        //fmt.Println(ptr.Val)
        _, exists := map1[*ptr]
        if exists == true  {
            return true
        }
        map1[*ptr] = 1
        ptr = ptr.Next
    }

    if ptr == nil {
        return false
    }
    //fmt.Println(map1)


    return false
    
}
