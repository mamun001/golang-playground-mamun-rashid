// 832 easy
// flipping an image

// 82.6 pct acc rate

// JAN 5, 2025

import "fmt"

func flip (n int) int {

    ans := 0

    if n == 0 {
        return 1
    } else if n == 1 {
        return 0
    }

    return ans
}



func flipAndInvertImage(image [][]int) [][]int {

    n := len(image[0])
    original_row := []int{}

    for i :=0 ; i <n; i++ {
        // save original row data before half of gets written over
        for j :=0 ; j < n ; j++ {
            original_row = append(original_row,image[i][j])
        }
        fmt.Println(original_row)
        for j :=0 ; j<n; j++ {
            // FORMULA new j = len -old j - 1 and THEN FLIP
            image[i][j] = flip(original_row[n-j-1])
        }
        original_row = []int{} // reset to empty before doing the next row
    }
    
    return image
}
