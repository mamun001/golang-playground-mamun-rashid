// easy 3158 Find the XOR of Numbers Which Appear Twice
// ELO 1172

// JAN 18, 2025

// There is fancier solution knowing that XOR has a trick where it can find uniq numbers
// But, I do not have that in my head
// so, I used a brute force method to find the dups


import "sort"

func duplicateNumbersXOR(nums []int) int {

    dups :=[]int{}
    sort.Ints(nums) // sorting makes finding dups easier

    // find ups
    for i:=1 ; i < len(nums) ; i++ {
      if nums[i-1] == nums[i] {
         dups = append(dups,nums[i])
      }
    }
    
    //fmt.Println(dups)

    if len(dups) == 0 {
        return 0
    } 
    if len(dups) == 1 {
        return dups[0]
    }

    // now find the XOR with the rest
    ans := dups[0] ^ dups[1]
    for i:=2 ; i<len(dups) ; i++ {
      ans = ans ^ dups[i]
    }

    return ans
}
