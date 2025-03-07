// easy 108, BST, part of NC 250
// convert sorted array to binary search tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

    if len(nums) == 0 {
        return nil
    }

    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]} // This line I had to lookup in chatgpt


    
    left_ar := nums[:mid]
    right_ar := nums[mid+1:]

    left_child := sortedArrayToBST(left_ar)
    right_child := sortedArrayToBST(right_ar)
    
    root.Left = left_child
    root.Right = right_child



    return root
    
}
