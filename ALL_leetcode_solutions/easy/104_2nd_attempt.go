// easy 104
// depth of binary tree 
// 2nd attempt, March 18, done without help

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    if root.Left != nil && root.Right != nil  {
        return 1+ max(maxDepth(root.Left),maxDepth(root.Right))
    }

    if root.Left != nil && root.Right == nil  {
        return 1+ maxDepth(root.Left)
    }

    if root.Left == nil && root.Right != nil  {
        return 1+ maxDepth(root.Right)
    }

    return -99

    
}
