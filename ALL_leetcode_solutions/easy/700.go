// easy 700, ELO unknown
// Search in a binary search tree

// kind of easy BT problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {

    if root == nil {
        return nil
    } else if root.Val == val {
        return root
    } else if root.Val < val {
        return searchBST(root.Right, val)
    } else {
        return searchBST(root.Left, val)
    }
    
    
}

