// MEDIUM 973, ELO 1231
// K closest points to Origin

// my alg CPU > 5% , RAM > 83%, so kind of slow
// I used map, heap would have been better

import "sort"

// get square of distance
func distsq (n1 int, n2 int) int {
    return n1*n1 + n2*n2
}

func kClosest(points [][]int, k int) [][]int {

    ans :=[][]int{}
    n := len(points)

    // don't use cpu cycles for simplest cases
    if k == len(points) {
        return points
    }

    // make a map of index in poiters ===> distance
    map1 := make(map[int] int)
    //fmt.Println(map1)

    // make a map of index in poiters ===> distance
    for i:=0 ; i<n ; i++ {
        map1[i] = distsq(points[i][0],points[i][1])
    }

    // find out the K shorest distances
    k_shorts := []int{}
    for i:=0; i<k;i++ {
        k_shorts = append(k_shorts,200000000)
    }
    //fmt.Println(k_shorts)
    d := 0
    threshold := 200000000
    for i:=0;i<n;i++ { // loop through all points in points
      d = distsq(points[i][0],points[i][1])
      if d < k_shorts[k-1] {
         //ans = append(ans,[]int{points[i][0],points[i][1]})
         k_shorts[k-1] = d // k-1 th has the largest k distance
         sort.Ints(k_shorts)
      }
      //fmt.Println(d)
    }
    //fmt.Println("k_shorts", k_shorts)
    //fmt.Println("maps", map1)

    // threshold = largest distance that qualifies
    threshold = k_shorts[k-1] // k-1 th has the largest k distance
    //fmt.Println("threshold",threshold)

    // loop thru pointers and insert to ans iF the point qualifies
    for k,v := range map1 {
        //fmt.Println(k,v)
        if v <= threshold {
            //fmt.Println(k,v)
            ans = append(ans,[]int{points[k][0],points[k][1]})
        }
    }


    //fmt.Println(map1[k_shorts[0]])
    //j := map1[k_shorts[0]]
    //fmt.Println(points[j][0],points[j][1])

    
    return ans
    
}
