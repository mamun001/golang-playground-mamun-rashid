/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Tested
// Mamun Rashid
// Convert a binary tree sent in the from of (example) [1,2,3,4,5,6,7] or [1,2,3,null,null,6,7] turn it back to a slice
// This may come in handy in caomparing two trees

func convert_to_slice(root *TreeNode) []int {
    
    answer := []int{} // This should never be used, we just define it so that GO lang does not complain about no return
    
    
    // if we are in void, nothing to return
    if root == nil {
        return []int{}
    }
    
    
    // if leave-node, rteurn its Value
    if root.Left == nil && root.Right == nil {
        slice_w_root := []int{root.Val}
        return slice_w_root
    }
    
    // if left side is dead, return value of current node's value appended by right side
    if root.Left == nil && root.Right != nil{
        slice_w_root := []int{root.Val}
        return append (slice_w_root,convert_to_slice(root.Right)...)
    }
    
    // if right side is dead, return left side appended by current node's value
    if root.Left != nil && root.Right == nil {
        return append (convert_to_slice(root.Left),root.Val)
    }
    
    // If we have both children alive
    // return = left sides + current node's value + right side's value
    if root.Left != nil && root.Right != nil {
        temp2 := append (convert_to_slice(root.Left),root.Val)
        temp3 := append (temp2,convert_to_slice(root.Right)...)
        return temp3
    }

    return answer 
}

