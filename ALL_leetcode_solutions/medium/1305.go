// medium 1305 ELO 1260
// did lazy programming
// beats CPU 8% and RAM 8%
// so, ya, barely scathed by


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


import "sort"



func to_slice(root *TreeNode) []int {
    ans :=[]int{}
    if root == nil {
        return ans
    }

    ans =append(ans,root.Val)
    ans = append(ans,to_slice(root.Left)...)
    ans = append(ans,to_slice(root.Right)...)

    return ans

}
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

    ans :=[]int{}

    t1 := to_slice(root1)
    t2 := to_slice(root2)
    ans = append(t1,t2...)
    sort.Ints(ans)
    
    return ans
    
}
