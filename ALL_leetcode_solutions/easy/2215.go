// easy, 2215, ELO 1208
// find the difference between two arrays
// 

func is_in(ar []int, n int) bool {
    ans:= false 
    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }
    return ans
}

func findDifference(nums1 []int, nums2 []int) [][]int {

    ans:=[][]int{}
    ans1 := []int{}
    ans2 := []int{}
    

    

    for i:=0;i<len(nums1);i++ {
        if is_in(nums2, nums1[i]) == false && is_in(ans1, nums1[i]) == false {
            ans1 = append(ans1,nums1[i])
        }
    }
    //fmt.Println(ans1)

    for i:=0;i<len(nums2);i++ {
        if is_in(nums1, nums2[i]) == false && is_in(ans2, nums2[i]) == false {
            ans2 = append(ans2,nums2[i])
        }
    }
    //fmt.Println(ans2)
    ans = append(ans,ans1)
    ans = append(ans,ans2)


    return ans
    
}
