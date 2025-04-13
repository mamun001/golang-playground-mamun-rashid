import "sort"

func is_in(ar []int, n int) bool {
    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }
    return false
}

func to_set(ar[]int) []int{
    ans:=[]int{}
    for i:=0;i<len(ar);i++{
        if is_in(ans,ar[i]) == false {
            ans=append(ans,ar[i])
        }
    }
    return ans
}

func maxDivScore(nums []int, divisors []int) int {

    // we have convert divisors to a set, because wheh are are repeats, because I am using map to count score, scores go artificially (doubles, triples etc)
    divisor_set:=to_set(divisors)
    //fmt.Println(divisor_set)

    map_scores:=make(map[int]int)

    for i:=0;i<len(nums);i++ {
        for j:=0;j<len(divisor_set);j++{
            if nums[i] % divisor_set[j] == 0 {
                map_scores[divisor_set[j]]++
            }
        }
    }

    high:=0
    winner:=0
    for k,v := range map_scores {
        if v > high {
            high=v
            winner=k
            //fmt.Println(winner,high,"higher high")
        }
        if v == high && k < winner {
            winner=k
            //fmt.Println(winner,high,"match high")
        }
    }

    if high > 0 { 
        return winner
    }

    if high == 0 { // nothing is divisible 
       //fmt.Println("high is 0")
       sort.Ints(divisor_set)
       return divisor_set[0]
    }

    return winner


    
}
