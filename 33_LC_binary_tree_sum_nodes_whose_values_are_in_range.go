/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func rangeSumBST(root *TreeNode, low int, high int) int {
    
    if root == nil {
		return 0
	}
 
    if root.Val <= high && root.Val >= low {
	   return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
    } else {
       return rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
    }
       
}
