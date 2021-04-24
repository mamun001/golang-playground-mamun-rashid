/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* test cases I made and 3 hard case from leetcode
[1]
[1,2]
[1,2,2]
[1,2,3]
[1,2,3,4]
[1,2,2,3,null,null,3]
[1,2,2,1,2,2,1]
[1,2,2,1,2,1,1]
[1,2,2,1,2,3,4,2,1,4,3]
[1,2,2,1,2,3,4,4,3,2,1]
[1,2,2,1,2,3,4,2,1,4,3,5,6]
[1,2,2,1,2,3,4,2,1,4,3,5,6,null,null,6,5]
[1,2,2,2,null,2]
[1,2,2,3,4,4,3]
[1,2,2,null,3,null,3]
*/

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

func compare_slices (s1 []int, s2 []int) bool {
    answer := true // we start assuming its true and make it false down the road if it false
    l1 := len(s1)
    l2 := len(s2)
    
    if l1 != l2 {
        return false // quick turn around
    } else {
        for i:=0; i < l1; i++ {
            if s1[i] == s2[i] {
                continue
            } else {
                answer = false
                return answer
                break
            }
        }
    }
    
    return answer
}

func reverse(s1 []int) []int {
    answer:= make([]int,0)
    
    
    L := len(s1)
    for i:=L-1; i >=0 ; i-- {
        answer=append(answer,s1[i])
    }
    
    return answer
}



func isSymmetric(root *TreeNode) bool {
    
    answer := true
    
    left_slice := convert_to_slice(root.Left)
    right_slice := convert_to_slice(root.Right)
    
    // Magic of this algorith (found accidentally Mamun Rashid)
    // Convert either half of tree to slices
    // reverse one of the slices and compare with the other(unreversed)
    // if they are the same, then halves of the trees are mirrors
    
    // This algorith has a blind spot
    // if the tree's halves have non-null values ONLY on side and on the side for each half
    // slice representation "cuts out" this information and the we return true even theough real answer is fale.
    // e.g. [1,2,2,2,null,2]
    // we have a edge case for this and test for before we call the compare function
    
    
    // edge cases may not be covered by the Algorithm
    if root.Left == nil && root.Right == nil {
        return true
    }
    
    if root.Left == nil && root.Right != nil {
        return false
    }
    
    if root.Left != nil && root.Right != nil {
        if root.Left.Left == nil && root.Right.Right != nil {
          return false
        }
        if root.Left.Right == nil && root.Right.Left != nil {
           return false
        }
        if root.Left.Left != nil && root.Right.Right == nil {
          return false
        }
        if root.Left.Right != nil && root.Right.Left == nil {
           return false
        }
    }
    
                                     
    // Main part of the algorithm
    reverse_left_slice := reverse (left_slice)
    //fmt.Println("reverse s1:", reverse_left_slice)
    
    answer = compare_slices(reverse_left_slice,right_slice)
    
    
    
    return answer
    
}
