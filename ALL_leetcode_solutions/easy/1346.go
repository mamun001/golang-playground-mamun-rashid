// easy 1346, ELO 1225
// Check if N and Its Double Exist

func checkIfExist(arr []int) bool {

    

    for i:=0;i<len(arr)-1;i++ {
      for j:=i+1;j<len(arr);j++ {
        if arr[i] == arr[j] * 2 || arr[i] * 2 == arr[j] {
            return true
        }
      }
    }

    return false
    
}
