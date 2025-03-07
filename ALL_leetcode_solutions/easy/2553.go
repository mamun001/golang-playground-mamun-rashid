// easy 2553, ELO 1217
// Separate the Digits in an array

func reverse(ar []int) []int{
    ans:= []int{}
    for i:=len(ar)-1;i>-1;i-- {
        ans = append(ans,ar[i])
    }
    return ans
}

func digits(n int) []int {
    ans := []int{}
    q := n / 10
    //fmt.Println("q",q)
    d := n % 10
    ans = append(ans,d)


    for q > 0 {
        //fmt.Println("here")
        d = q % 10
        //fmt.Println("d",d)
        q = q / 10
        //fmt.Println("q",q)
        ans=append(ans,d)
    }

    return reverse(ans)
}

func separateDigits(nums []int) []int {

    ans := []int{}

    //fmt.Println(digits(123))
    for i:=0;i<len(nums);i++ {
        //fmt.Println(digits(nums[i]))
        ans = append(ans,digits(nums[i])...)
    }

    return ans
    
}
