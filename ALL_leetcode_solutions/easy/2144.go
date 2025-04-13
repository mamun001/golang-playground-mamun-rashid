// easy 2144, ELO 1261
// minimum cost of buying candies with discount

import "sort"

func reverse(ar []int) []int {
    ans:=[]int{}
    for i:=len(ar)-1;i>-1;i--{
        ans=append(ans,ar[i])
    }
    return ans
}

func minimumCost(cost []int) int {

    if len(cost) == 1 {
        return cost[0]
    }

    if len(cost) == 2 {
        return cost[0]+cost[1]
    }

    sort.Ints(cost)
    cost=reverse(cost)
    //fmt.Println(cost)

    ans:=0
    left_overs:=len(cost) % 3
    L:=len(cost)

    
    for i:=0;i<L-left_overs;{
        ans=ans+cost[i]+cost[i+1]
        i=i+3
    }
    

    if left_overs == 1 {
        ans=ans+cost[L-1]
    }

    if left_overs == 2 {
        ans=ans+cost[L-1]+cost[L-2]
    }


    return ans

    
}
