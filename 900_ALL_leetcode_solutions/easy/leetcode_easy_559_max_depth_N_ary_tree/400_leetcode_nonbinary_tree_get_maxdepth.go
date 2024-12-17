/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// This code works for small trees, after times out after 10 seconds with large tress

func maxDepth(root *Node) int {
    
    answer := 0 // To start with 0 for any tree
    
    if root != nil {
        if len(root.Children) == 0 {
            // we are at a node and return 1
            answer = 1
        } else {
            //we are NOT at a node, let's calculate the MAX of depths of my children
            
            max := 0 // We start pretending there is no depth to any my children
            // Now we calculate the depth of each children and take the biggest
            for i:=0; i<len(root.Children); i++ {
                if maxDepth(root.Children[i]) > max {
                    max = maxDepth(root.Children[i])
                }
            }
            
            answer = 1 + max // return the max of depths of my children
        }
        
    } // root nil IF
    
    return answer
} // FUNC
    
   
