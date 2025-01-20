// MEDIUM 1441 Build an Array with stack Operations
// ELO 1180

// JAN 19, 2025
// My alg performance: CPU beats 100%, RAM beats 32%


//import "fmt"

func is_equal (sl1 []int, sl2 []int) bool {
    ans := true
    if len(sl1) != len(sl2) {
        return false
    }

    for i:=0; i < len(sl1); i++ {
        if sl1[i] != sl2[i] {
            return false
        }
    }
    return ans
}

func buildArray(target []int, n int) []string {

    ans := []string{}
    stack := []int{} // this is where push or pop
    target_pos := 0 // index at target
    stack_complete := false
    stream_val := 1 // index in stream, same as value
    //a := 0

    //testing my func
    //a := []int{1,3}
    //b := []int{1,3}
    //fmt.Println(is_equal(a,b))

    for stack_complete == false {
        //fmt.Println("stack at the top",stack)
        // pushing is a must
        stack = append(stack,stream_val)
        ans = append(ans,"Push")
        //fmt.Println("stream_val",stream_val)
        //fmt.Println("target value",target[target_pos])
        if stream_val == target[target_pos] { // push was the right thing to do
           //fmt.Println("just pushed")
           target_pos++
        } else { // wrong value, let's pop
            //fmt.Println("popping")
            ans = append(ans,"Pop")
            stack = stack[:len(stack)-1] // remove the last element
        }
        stream_val++
        //fmt.Println("stack at the bottom",stack)
        //fmt.Println("checking if stack is same as target")
        if is_equal(target,stack) {
            //fmt.Println("stack and target are equal")
            stack_complete = true
        }
        //fmt.Println("--------------------")
    } // for

    return ans
}
