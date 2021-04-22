
func is_empty(LL *ListNode) bool {
    if LL == nil {
        return true
    } else {
        return false
    }
}

func insert_first(LL *ListNode, Value int) *ListNode {

	LL = &ListNode{Val: Value, Next: nil}
	return LL
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    var answer *ListNode
    
    fmt.Println(is_empty(answer))
    
    answer = insert_first(answer, 10)
    
    fmt.Println(is_empty(answer))
    
    
    
    return answer
    
}
