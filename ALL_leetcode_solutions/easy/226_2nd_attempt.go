// easy 226
// invert binary tree
// 2nd attempt, no help

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

    if root == nil {
        return nil
    }

    invertTree(root.Left)
    invertTree(root.Right)
    temp:=root.Left
    root.Left=root.Right
    root.Right=temp
    

    return root

    
}
