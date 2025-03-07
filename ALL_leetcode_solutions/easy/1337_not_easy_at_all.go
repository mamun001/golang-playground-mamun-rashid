// easy 1337
// finally done the 3rd time
// spent several hours the first time
// spent another 3+ second time
// 3rd time was the charm!!
// was able to do this without using heap
// my alg CPU > 100% (1 sec), some of the submission even had 4 ms, mine was 1 ms
//        RAM > 53%
//So, not bad!

// I think is rated low (1224) , because Python has a heap module that needs no additional code

// MY overall alg:
// make a func for strength that takes into account the "index" , because for same "strength", index matters
// make a map of "strength" to index (row# of matrix)
// make a slice of only the keys of the map
// sort the slice and take only K values
// NOW look up the index (row value) in the map above for the top k values in the slice
// make a quick array out of those indexes

import "sort"

func strength(ar []int , f float64) float64 {
    var ans float64 = 0.0
    for i:=0;i<len(ar);i++ {
        if ar[i] == 1 {
            ans = ans + 1.00
        } else {
            return ans + f
        }
    }
    //fmt.Println(float32(ind))
    ans = ans + f
    return ans
}

func kWeakestRows(mat [][]int, k int) []int {

    ans:=[]int{}

    //a:=[]int{1,0}
    //fmt.Println(strength(a))

    //i := 1
    //fmt.Println(float32(i)/100.0)

    var f float64 = 0.0
    map1:=make(map[float64]int)
    for i:=0;i<len(mat);i++{
        f = float64(i)/100.0
        //fmt.Printf("%.2f\n", f)
        //fmt.Printf("%.2f\n", strength(mat[i],f))
        //fmt.Println(strength(mat[i],f))
        map1[strength(mat[i],f)] = i     
    }

    slice1 := []float64{}
    for k,_ := range map1 {
        slice1 = append(slice1,k)
    }

    sort.Float64s(slice1)
    slice1 = slice1[:k]
    //fmt.Println(slice1)

    for i:=0;i<len(slice1);i++ {
        //fmt.Println(map1[slice1[i]])
        ans = append(ans,map1[slice1[i]])
    }
    

    return ans
    
}
