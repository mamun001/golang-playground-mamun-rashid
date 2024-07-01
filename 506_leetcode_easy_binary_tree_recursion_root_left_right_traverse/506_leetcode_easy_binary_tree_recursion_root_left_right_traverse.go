

// leetcode easy
// 617. Merge Two Binary Trees (Leetcode)

// This is solution for:  https://leetcode.com/problems/merge-two-binary-trees/
// However , I did not write this. A leetcode user named Xing98. Credit belongs to him/her

// I only added comments and changed variables names so it makes more sense to me (while learning trees in Golnag)

// & = RAM ADDRESS
// * = variable that holds that address and resolves it
// & * = are complementary
// e.g.
// myString := "Hi"
// fmt.Println(*&myString)  prints "Hi"


// When you put * in front of TYPE, e.g. *int, you are really saying:
// FOO *int = variable that holds RAM ADDRESS (pointer) to int



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {   
    //  t1 and t2 VARs that holds ADDRESS of type "treenode" which is STRUCT already defined by leetcode

    // if one of t1 and t2 is nil, return the other
    if t1 == nil {
        //   NO RAM ADDRESS
        return t2
    }
    if t2 == nil { 
    //  NO RAM ADDRESS
        return t1
    }

    // ADD Values of t1 and t2 FOR THIS NODE ONLY
    ANSWER := &TreeNode{Val: t1.Val + t2.Val} 
    // ANSWER IS THE RAM ADDRESS that holds TREENODE TYPE and its "Val" = sum of the VALs of the two nodes of the two trees
    
    // NOW recurse!  Left and Right
    ANSWER.Left = mergeTrees(t1.Left, t2.Left)
    ANSWER.Right = mergeTrees(t1.Right, t2.Right)
    return ANSWER
}
