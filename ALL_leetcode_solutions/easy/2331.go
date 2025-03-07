// easy 2331, ELO 1304
// Evaluate Boolean Tree 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {

    if root.Val == 0 {
        return false
    }

    if root.Val == 1 {
        return true
    }

    if root.Val == 2 {
        return evaluateTree(root.Left) ||  evaluateTree(root.Right)
    } 

    if root.Val == 3 {
        return evaluateTree(root.Left) && evaluateTree(root.Right)
    } 


    return true
    
}
