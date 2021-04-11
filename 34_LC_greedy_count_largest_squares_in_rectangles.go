func countGoodRectangles(rectangles [][]int) int {
    
    var shorter_side int = 1001
    shorter_sides := make([]int, 0)
    
    for i:=0;i<len(rectangles); i++ {
        
        if rectangles[i][0] <= rectangles [i][1] {
            shorter_side = rectangles[i][0]
        } else {
            shorter_side = rectangles[i][1]
        }
        
        shorter_sides = append(shorter_sides,shorter_side)
        
     } // FOR
    
    largest_shorter_side := 0
    for i:=0; i<len(shorter_sides); i++ {
        if shorter_sides[i] > largest_shorter_side {
            largest_shorter_side = shorter_sides[i]
        }
    } // FOR
        
    count_largest := 0
    for i:=0; i<len(shorter_sides); i++ {
        if shorter_sides[i] == largest_shorter_side {
            count_largest++
        }
    } 
    
    
    return count_largest
    
    } // FUNCTION
    
    
