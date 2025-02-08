// 938 easy ELO 1335!
// Range Sum of BST
// did it a second time!


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func get_sum(root *TreeNode, low int, high int) int {

    ans := 0
    left_sum := 0
    right_sum := 0

    if root == nil {
       //fmt.Println(root.Val)
       return 0
    } 
    
    if root.Left != nil {
          //print_tree(root.Left)
           left_sum = get_sum(root.Left, low, high)
    } 
    
    if root.Right != nil {
          //print_tree(root.Right)
          right_sum = get_sum(root.Right, low , high)
    }

    if root.Val >= low && root.Val <= high {
        ans = root.Val + left_sum + right_sum
    } else {
        ans = left_sum + right_sum
    }

    return ans


}


func rangeSumBST(root *TreeNode, low int, high int) int {


    //ans := 0

   //print_tree(root)
   return get_sum(root, low, high)

return 0

    
}
