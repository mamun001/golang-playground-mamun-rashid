// MEDIUM 1282, ELO 1267
// group the people given the group size ....

func groupThePeople(groupSizes []int) [][]int {

    ans := [][]int{}

    // map all the people ids to group sizes
    map_size_persons := make(map[int][]int)

    for i:=0;i<len(groupSizes);i++{
        map_size_persons[groupSizes[i]] = append(map_size_persons[groupSizes[i]],i)
    }

    //fmt.Println(map_size_persons)


    // for some groups, you will have N multiples of group sizes in value of the map
    // e.g. grou size of 3
    // map lmay look like
    // key=3 value=[2,7,9,6,3,5,10,12,11] (9 people , groups of 3)
    /*
    0,1,2 i=0 k=3  k*i+0 k*i+1 K*2
    3,4,5
    6,7,8
    */

    for k,v := range(map_size_persons) {
        //if k == len(v) {
           //ans = append(ans,v)
        //} 
        many := len(v)/k
        for i:=0;i<many;i++{ // make temp slices of exactly size k
            temp_slice := v[k*i:k*(i+1)]
            ans=append(ans,temp_slice)
        }
    }


    return ans
    
}
