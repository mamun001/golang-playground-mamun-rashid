// easy 2236 , ELO unknown
// root equals sum of children
// easiest Binary tree problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {

    return root.Val == root.Left.Val + root.Right.Val
    
}
