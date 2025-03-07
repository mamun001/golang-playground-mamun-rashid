// Medium 881, ELO 1521
// boats to save people
// my alg worked for most cases, BUT, failed on a long one, that could not troubleshoot
// had to lookup youtube. This is algorithm is almost same as mine, but , mine had
// too many if end else if, making it hard to debug
// this is simpler and works for all cases

import "sort"

func numRescueBoats(people []int, limit int) int {

    
    boats_used := 0
    sort.Ints(people)
    //fmt.Println(people)
    l:=0 // lightest left
    r:=len(people)-1 // heaviest left
    available_weight :=0 // weight remaining after putting the heaviest person on


    for l<=r {
        //fmt.Println("l,r:",l,r)
        boats_used++
        //fmt.Println("boats",boats_used)
        available_weight = limit -people[r]
        //fmt.Println("avaiable weight",available_weight)
        r-- // took the heaviest person
        //fmt.Println("r",r)
        if l<=r && people[l] <= available_weight {// at least 1 person left & we can fit 1 more
           l++
           //fmt.Println("l:",l)
        }     
    }

    return boats_used
    
}
