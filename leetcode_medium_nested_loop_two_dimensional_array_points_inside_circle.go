// leetcode medium #1828 Queries on Number of Points Inside a Circle
// nested loops, 2-dimensional arrays, math

import (
    "fmt"
    "math"
)


// find out if the given point is inside a given circle
func isInside (px int,py int,cx int,cy int,cr int) bool {
    
    distance_from_center := math.Sqrt(math.Pow(float64(px-cx),2)+math.Pow(float64(py-cy),2))
    
    if distance_from_center <= float64(cr) {
        return true 
    } else {
        return false
    } // if
    
} // func


// for a given query (circle), how many given points are inside
func countPoints(points [][]int, queries [][]int) []int {
    
    answer := make([]int, 0)
    
    // why are queries called queries, they are circles
    
    // first example 4 POINTS & 3 CIRCLES(queries) green, red, blue
    
    // 3 of 4 points are in GREEN circle
    // 2 of 4 points are in RED
    // 2 of 4 points are in BLUE
    
    
    // take first point (P1) 1,3 , it falls is in GREEN because 
    // green center is 2,3
    // distance between P1 and Gcenter = sqrt(1+0) = 1 which <= Gradius which 1 = INSIDE
    
    // P2 = 3,3
    // green center = 2,3
    // distance = sqrt(1,1) =1 <= Green Radius (1) = INSIDE
    
    // troubleshooting
    //fmt.Println(isInside(px,py,cx,cy,cr))
    //fmt.Println(isInside(points[0][0],points[0][1],2,3,1))
    //fmt.Println(queries[0][0],queries[0][1],queries[0][2])
    //fmt.Println(isInside(points[0][0],points[0][1],queries[0][0],queries[0][1],queries[0][2]))
    //fmt.Println(len(points))   
    
    for q:=0 ; q<len(queries); q++ {
        count := 0
        for i:=0; i<len(points) ; i++ {
          if (isInside(points[i][0],points[i][1],queries[q][0],queries[q][1],queries[q][2])) == true {
              count++
          }
        }
        answer = append(answer,count)
    }
    
    return answer
    
}
