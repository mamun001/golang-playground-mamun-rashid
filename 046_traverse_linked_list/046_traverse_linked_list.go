//
// Written by Mamun Rashid
//
// ONLY THE FUNCTION , no main
//
// see 056_linked_list_length_traverse_insert_at_one_and_two.go for full code and explanations

func traverse(head *ListNode) {
    if head.Next == nil {
        fmt.Println(head.Val)
        return
    } else {
        fmt.Println(head.Val)
        traverse(head.Next)
    }
}
