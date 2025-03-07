// easy 112, no ELO, binary tree, acc rate 52%
// did without help!

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {

    ptr := root


   if ptr == nil {
       //fmt.Println("ptr is nill, returning false")
       return false
   } else {
       //fmt.Println(ptr.Val,targetSum)
   }


   if ptr.Left == nil && ptr.Right == nil && ptr.Val == targetSum {
        //fmt.Println(ptr.Val,"returning true")
        return true
   } 

   if ptr.Left == nil && ptr.Right == nil && ptr.Val != targetSum {
        //fmt.Println(ptr.Val,"returning false")
        return false
   } 

   

   //fmt.Println(ptr.Val,"calling left and right")
   return hasPathSum(ptr.Left, targetSum - ptr.Val) || hasPathSum(ptr.Right, targetSum - ptr.Val)
   
   
   


    return false
    
}
