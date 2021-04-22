/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Code : Mamun rashid
// Luckily, I cam up with this succinct code myself
// Ran faster than 100% of all Go Submissions


func invertTree(root *TreeNode) *TreeNode {
    
    answer := root
    
    if root == nil {
        return nil
    } else {
        answer.Val = root.Val
        hold_left := answer.Left 
        answer.Left = answer.Right
        answer.Right = hold_left
        invertTree(answer.Left)
        invertTree(answer.Right)
    }
    
    return answer
    
}
