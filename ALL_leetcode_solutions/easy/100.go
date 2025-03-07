// easy, 100, same tree
// part of NC 250

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

    ans:= true

    if p == nil  && q == nil{
        return true
    }

    if p == nil  && q != nil{
        return false
    }

    if p != nil  && q == nil{
        return false
    }

    if p.Val != q.Val {
        return false
    }

    if p.Left == nil && q.Left != nil {
        return false
    }

    if p.Right == nil && q.Right != nil {
        return false
    }

    
    l := isSameTree(p.Left , q.Left)
    r := isSameTree(p.Right , q.Right)
    

    if l == false {
        return false
    }
    
    if r == false {
        return false
    }

    return ans
}
