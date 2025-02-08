// easy 872, ELO 1288
// BT
// Leaf-Similar Trees

// my alg CPU > 100% , RAM > 39%
// SO, NOT BAD!

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


import "slices"

func get_leaves(root *TreeNode) []int{ 

    ans := []int{}

    if root == nil {
        return ans
    }

    if root.Left == nil && root.Right == nil {
        return append(ans,root.Val)
    } 

    ans = append(ans,get_leaves(root.Left)...) 
    ans = append(ans,get_leaves(root.Right)...) 

    return ans


}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

    ans := false

    //fmt.Println(get_leaves(root1))
    //fmt.Println(get_leaves(root2))

    if slices.Equal(get_leaves(root1),get_leaves(root2)) {
        ans = true
    }

    return ans
    
}

