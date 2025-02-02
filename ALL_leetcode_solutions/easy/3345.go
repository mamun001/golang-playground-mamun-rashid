// easy 3345 ELO 1235 smallest divisible digit
// this code is mostly lazy, max n is so small

func digits(n int) []int {
    ans :=[]int{}
    if n >= 100 {    
        ans = append(ans,n/100)  
        ans = append(ans,n/10)
        ans = append(ans,n%10)
    }
    if n >9 && n<100 {
        ans = append(ans,n/10)
        ans = append(ans,n%10)
    }
    if n < 10 {
        ans = append(ans,n)
    }

    return ans
}

func smallestNumber(n int, t int) int {

   //fmt.Println(digits(100))

   ans := n
   //found := false
   //for i:=n;found == false;i++ {
   for i:=n;i<1000;i++ { // chose 1000 as an impossible number, since product will break when ans is found
       product := 1
       for j:=0;j<len(digits(i));j++{
          product = product * digits(i)[j]
       }
       //fmt.Println(product)
       if product % t == 0 {
          return i
       }
   }

    return ans
    
}

