
// leetcode easy
// linked list, length, traverse, two_to_the_power_function
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */



func length(head *ListNode) int {
    //answer := 0
    if head.Next == nil {
        return 1
    } else {
        return 1 + length(head.Next)
    }
}


// THIS IS THE MAIN FUNCTION
func get_answer(head *ListNode, P int) int {
    
    // get the sum of all binary digits' values
    
    if head.Next == nil {
        return head.Val
    } else {
        return ( two_to_the_power(P) * (head.Val) + get_answer(head.Next, P-1))
    }
}

func two_to_the_power (n int) int {
    // 2^n
    if n == 0 {
        return 1
    } else {
        return 2 * two_to_the_power (n-1)
    }
}


func getDecimalValue(head *ListNode) int {
    
    l := length(head) // How long the 
    
    sum := get_answer(head,l-1)
    
    return sum
    
}
