
// Written by Mamun Rashid

func traverse(head *ListNode) {
    if head.Next == nil {
        fmt.Println(head.Val)
        return
    } else {
        fmt.Println(head.Val)
        traverse(head.Next)
    }
}
