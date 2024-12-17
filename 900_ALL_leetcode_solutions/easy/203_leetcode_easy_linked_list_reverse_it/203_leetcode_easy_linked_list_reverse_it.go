
// leetcode easy
// reverse linked list

// BEST SOLUTION from FoodMuncher  At Leetcode
// Mamun Rashid added the comments

func reverseList(head *ListNode) *ListNode {  // given by LC
    var ANSWER *ListNode   // define new blank LL
    
    for head != nil {  // LOOP THROUGH EACH POINTER BEGGING AT HEAD IE THE START
        ANSWER = &ListNode{Val: head.Val, Next: ANSWER}  // THIS LINE HAS TO BE MEMORIZED. THERE IS NO EASIER SOLUTION
                                                         // FILL UP THE STRUCT ANSWER NOW POINTS TO
                                                         // ANSWER IS REDEFINED LIKE a=a+1
                                                         // head is current pointer in the original list
                                                         // &ListNode means = ANSWER is POINTER (&) TO TYPE ListNode  : HARD PART
                                                         // BUT, listnode needs two values what are those? what's inside {} block
                                                         // Whsoe values are with { } block
                                                         // in the {} block Val: head.Val  = original list's current pointer's value
                                                         // ANSWER's NEXT is current ANSWER  : we just did the a=a+1 thing
        head = head.Next                                 // MOVE the porint pointing to origibal lost to next thing and LOOP TILL we are at void
     }
    
    return ANSWER
}
