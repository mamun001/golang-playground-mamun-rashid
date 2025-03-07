// easy 110, Balanced Binary Tree
// part of NC 250!

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func abs (n int) int {
    if n >= 0 {
        return n
    } else {
        return n*(-1)
    }
}


func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    if root.Left != nil || root.Right != nil{
        return 1 + max(height(root.Left),height(root.Right))

    }


    return 0
}

func isBalanced(root *TreeNode) bool {

    //fmt.Println(height(root))

    if root == nil {
        return true
    }

    l := height(root.Left)
    r := height(root.Right)

    //fmt.Println(l,r)

    if abs(l-r) >1 {
        return false
    } else {
        return isBalanced(root.Left) && isBalanced(root.Right)
    }
    
}
