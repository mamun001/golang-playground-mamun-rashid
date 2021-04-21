/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// There are much better solutions, but this is mine and it worked after 2-3 days of trying -Mamun Rashid


func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    
    // Create a fake default node with no children
    answer := TreeNode{Val: -1024, Right: nil, Left: nil}
    
    // if both trees' nodes are at void, return nil pointer
    if root1 == nil && root2 == nil {
        return nil
    }
    
    
    // If tree1's node is void, return value from tree2
    if root1 == nil && root2 != nil {
        answer.Val = root2.Val
        if root2.Left != nil { 
            answer.Left = root2.Left // root1 does not even exist
        } else {
            answer.Left = nil
        }
        
        if root2.Right != nil {
            answer.Right = root2.Right // root1 does not even exist
        } else {
            answer.Right = nil
        }
        return &answer
        
    }
    
    // If tree2's node is void, return value from tree1
    if root1 != nil && root2 == nil {
        answer.Val = root1.Val
        if root1.Left != nil {
           answer.Left = root1.Left
        } else {
            answer.Left = nil
        }
        
        if root1.Right != nil {
            answer.Right = root1.Right // root2 does not even exist
        } else {
            answer.Right = nil
        }
        return &answer
    }
    
    // if both input trees have values, return the sum
    if root1 != nil && root2 != nil {
        answer.Val = root1.Val + root2.Val
        answer.Left = mergeTrees(root1.Left , root2.Left) 
        answer.Right = mergeTrees(root1.Right , root2.Right) 
        return &answer
    }
        
    return &answer    
    
}
