// easy 965, ELO 1177
// univalued binary tree

// Needed youtube help to get it done



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {

    ans := true

    if root != nil {
        fmt.Println(root.Val)
    }

    if root == nil {
        return true
    } else if root.Left != nil && root.Left.Val != root.Val {
        return false
    } else if root.Right != nil && root.Right.Val != root.Val {
        return false
    } else {
        return isUnivalTree(root.Left) && isUnivalTree(root.Right)
    }


    return ans
    
}
