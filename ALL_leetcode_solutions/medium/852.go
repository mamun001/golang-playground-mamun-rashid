// medium 852 Peak Index in a mountain array
// ELO 1178 or so

// JAN 19, 2025

// my code finaly worked
// CPU beats 100%
// RAM beats 24%

// I forgeting the low = low+1 and high = high -1 in the key part of simple binary search
// and so keep reinventing the wheel

import "fmt"

func peakIndexInMountainArray(arr []int) int {
    L := len(arr)
    pos := len(arr)/2 
    last_pos := 0
    found := false 

    for found == false {
        fmt.Println(pos,arr[pos])
        if arr[pos] > arr[pos-1] && arr[pos] > arr[pos+1] { // decreasing in both directions i.e. peak
            found = true
            return pos
        } else if arr[pos] < arr[pos+1]  { // numbers are increasing to the right
            last_pos = pos
            pos = (pos + L) / 2 // move pos half the distance to the right
        } else if arr[pos] > arr[pos+1] {// numbers are increasing to the left
            pos = (last_pos + pos) / 2 // move pos half the distance to the left
        } // else
    } // for

    return pos
    
}
