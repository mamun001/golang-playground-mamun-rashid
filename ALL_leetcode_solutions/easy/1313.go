

// leetcode easy level, 1313
// stack

type stack []int

func (s stack) Push(v int) stack {
    return append(s, v)
}


func decompressRLElist(nums []int) []int {
    l := len(nums)
    //fmt.Println(l)
    pairs := l/2
    //fmt.Println(pairs)
    s := make(stack,0)
    
    
    for i:=0; i < pairs; i++ {
        //fmt.Println("i",i)
        times:=nums[i*2]
        //fmt.Println("times",times)
        what_to_push:=nums[i*2+1]
        //fmt.Println("what",what_to_push)
        for freq:=1; freq <= times; freq++ {
            //fmt.Println("freq",freq)
            s = s.Push(what_to_push)
        }
        //fmt.Println()
        
    }
    
    return s
    
}
