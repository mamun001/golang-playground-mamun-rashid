// 1337
// I don't know why is labelled easy. it should be labbelled medium
// finally done the 3rd time w/o help
// 4th time (March 22) again done w/o help, again it took almost 1 hour to debug


import "sort"

func strength(ar []int) int {
    ans:=0
    for i:=0;i<len(ar);i++{
        if ar[i]==0 {
            return ans
        } else if ar[i] == 1 {
            ans++
        }
    }
    return ans
}


// only get uniq elements form a sorted array
func uniq(ar []int) []int{
    ans:=[]int{}
    ans=append(ans,ar[0])
    for i:=1;i<len(ar);i++{
        if ar[i] != ar[i-1] {
            ans=append(ans,ar[i])
        }
    }
    return ans
}

func kWeakestRows(mat [][]int, k int) []int {

    ans:=[]int{}
    map_str_ind:=make(map[int][]int) // map strength to index
    ar_str:=[]int{} // array of just the strength, so later we can pick only kth weakest ones

    // fill the map and array defined above
    for i:=0;i<len(mat);i++ {
        //fmt.Println(strength(mat[i]))
        str:=strength(mat[i])
        map_str_ind[str]=append(map_str_ind[str],i)
        ar_str=append(ar_str,str)
    }
    //fmt.Println(map_str_ind)
    // sort the array of strengths
    sort.Ints(ar_str)
    //fmt.Println("sorted",ar_str)
    // only get the uniq elements from that array
    ar_str=uniq(ar_str)
    //fmt.Println("uniq",ar_str)
    // picked = how many have we picked off so far for the answer
    picked:=0
    for i:=0;picked<=k && i<len(ar_str);i++{
        //fmt.Println("in map, of arrays ith index",i,map_str_ind[ar_str[i]])
        for j:=0;j<len(map_str_ind[ar_str[i]]);j++ {
           this:=map_str_ind[ar_str[i]][j] // picking this index for the answer
           //fmt.Println("this",this)
           ans=append(ans,this)
           picked++
        }
    }
    //fmt.Println(ans)
    // sometimes ans ends up more than k elements
    ans=ans[0:k]
    
    return ans
    
}
