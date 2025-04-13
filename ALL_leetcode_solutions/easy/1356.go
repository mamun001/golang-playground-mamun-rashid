// easy 1356, ELO 1258
// sort integers by the number of 1 bits
// Took me 90minutes to code and debug
// BUT performance: CPU > 47% and RAM > 5%, so ok+
// Youtube solutions 2 in Python and 1 java
// Python ones usedd heap (1) and fancy lambda func (1)

import "sort"

func ones(n int) int {
    q:=n
    ans:=0

    for q>0 {
        if q%2 == 1 {
            ans++
        }
        q=q/2
    }
    return ans
}


func sortByBits(arr []int) []int {

    
    // value(number) to count of ones MAP
    map_val_ones:=make(map[int]int) 
    for i:=0;i<len(arr);i++ {
        map_val_ones[arr[i]] = ones(arr[i])
    }
    //fmt.Println(map_val_ones)


    // we need track if any values come more than once
    map_val_count:=make(map[int]int) 
    for i:=0;i<len(arr);i++ {
        map_val_count[arr[i]]++
    }
    //fmt.Println(map_val_count)


    // make a map of number_ones_to ==> number is arr
    // "values" is an array because same count can be on multiples indexes
    map_count_vals:=map[int][]int{} 
    for k,v := range map_val_ones {
        map_count_vals[v]=append(map_count_vals[v],k)
    }
    //fmt.Println(map_count_vals)

    //sort each value(array) in the map above (ascending)
    for _,v := range map_count_vals{
        sort.Ints(v)
    }
    //fmt.Println(map_count_vals)


    // when we range through a map, there is no guarntee that k's will come at hand sorted
    // even though they are sorted when printed
    // so we have to put the Keys in a array and sort it
    keys:=[]int{}
    for k,_ := range map_count_vals{
        keys=append(keys,k)
    }
    sort.Ints(keys)
    //fmt.Println(keys)


    //now simply make a new array of all numbers looked up by sorted keys above
    ans:=[]int{}
    for i:=0;i<len(keys);i++ {
        for j:=0;j<len(map_count_vals[keys[i]]);j++ {
              for k:=0;k<map_val_count[map_count_vals[keys[i]][j]];k++ { // if there are repeats in arr
                 ans=append(ans,map_count_vals[keys[i]][j])
              }
        }
    }

    
    return ans
    
}
