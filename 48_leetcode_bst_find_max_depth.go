/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func larger(a int, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
    
}
    
func maxDepth(root *TreeNode) int {
    
    answer := 0
    
    
    
    if root == nil {
        return 0
    } else {
        return 1 + larger(maxDepth(root.Left),maxDepth(root.Right))
    }
    
    return answer
    
}
