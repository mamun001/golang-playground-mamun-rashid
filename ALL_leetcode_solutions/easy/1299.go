// easy 1299 , ELO 1219
// replace elements with greatest

func replaceElements(arr []int) []int {

    next_right := arr[len(arr)-1]
    max_to_the_right := arr[len(arr)-1]

    for i:=len(arr)-1;i>-1;i-- {
        if i == len(arr)-1 {
            arr[i] = -1
        } else if max_to_the_right < next_right {
            max_to_the_right = next_right
            next_right = arr[i]
            arr[i] = max_to_the_right
        } else {
            next_right = arr[i]
            arr[i] = max_to_the_right
        }
    }
    

    return arr
    
}
