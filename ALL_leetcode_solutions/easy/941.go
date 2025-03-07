// easy 941, ELO 1208
// valid mountain array

func validMountainArray(arr []int) bool {

    ans := true
    fp := 10001
    fp_found := false

    if len(arr) <3 {
        return false
    }


    // edge case 
    if arr[1] < arr[0] {
        return false
    }

    for i:=1 ; i<len(arr)-1; i++ {
        if arr[i] == arr[i-1] || arr[i] == arr[i+1] {
            return false
        }
        if i>0 && i<len(arr)-1 && fp_found == false && arr[i-1] < arr[i] && arr[i] > arr[i+1] {
            fp = i // first_peak
            fp_found = true
        }
    }

    //fmt.Println(fp)
    if fp == 10001 {
        return false
    }

    // once we find a peak, every number after that better be strictly decreasing
    for i:=fp;i<len(arr)-1;i++ {
        if arr[i+1] >= arr[i] {
            return false
        }
    }




    return ans
    
}
