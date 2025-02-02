// 2465 easy, ELO 1250
// Number of distinct averages


import "sort"

func is_in(f []float64, f2 float64) bool {
    ans := false
    for i:=0;i<len(f);i++ {
        if f[i] == f2 {
            return true
        }
    }
    return ans
}

func distinctAverages(nums []int) int {

    //ans := -1
    sort.Ints(nums)
    n := len(nums)
    n_const := len(nums)
    ave := -100.0
    averages :=[]float64{}
    
    for i:=0;i<(n_const/2);i++ {
        //fmt.Println("_____________________")
        //fmt.Println(nums)
        ave = (float64(nums[0])+float64(nums[n-1]))/2.0
        //fmt.Println(ave)
        averages = append(averages,ave)
        nums=nums[1:n-1]
        n = len(nums)
    }
    //fmt.Println(averages)
    //sort.Float64s(averages)
    //fmt.Println(averages)

    uniq :=[]float64{}
    for i:=0;i<len(averages);i++{
        if is_in(uniq,averages[i]) == false {
            uniq = append(uniq,averages[i])
        }
    }
    //fmt.Println(uniq)

    return len(uniq)
    
}

