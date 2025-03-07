// MEDIUM 167, Two Sum II
// did it w/o help!!
// my alg, beats cput 100% , beats RAM 72%

func twoSum(numbers []int, target int) []int {

    ans :=[]int{0,0}

    map1 := make(map[int]int)
    repeat_map := make(map[int]int)
    

    //fmt.Println(numbers[0])

    ptr:=0

    for ptr <len(numbers)  {
        need  := target - numbers[ptr]
        _, exists_map1 := map1[need]
        if exists_map1 == true {
            repeat_map[numbers[ptr]] = ptr
        } else {
            map1[numbers[ptr]] = ptr
        }

        //need  := target - numbers[ptr]
        _, exists_repeat := repeat_map[need]

        if need == numbers[ptr] && exists_repeat == true {
            ans[0] = map1[need]+1
            ans[1] = repeat_map[need]+1
            return ans
        }

        _, exists_map1 = map1[need] 
        if exists_map1 == true && map1[need] != ptr {
            ans[1] = ptr+1
            ans[0] = map1[need]+1
            return ans
        }

        ptr++
        
    }

    return ans
    
}
