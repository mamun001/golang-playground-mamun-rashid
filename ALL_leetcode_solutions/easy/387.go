// 62.9 pct

// easy 387

// First Unique Character in a String

//import "fmt"

func is_in_slice (sl []byte , b byte) bool {
    ans := false

    for _,v := range sl {
        if v == b {
            ans = true
            break
        }
    }

    return ans
}

func firstUniqChar(s string) int {

    ans := -1



    found_letters_so_far := []byte{} // letters we have found so far
    repeats := []byte{}

    
    //fmt.Println(is_in_slice(found_letters,'c'))
    

    // first pass to find repeats
    for _,v:= range s {
        //fmt.Println(byte(v))
        if is_in_slice(found_letters_so_far,byte(v)) { // already found earlier
            repeats = append (repeats,byte(v)) // if there are 3 or more, they all will be added and that's ok

        } else {
            found_letters_so_far = append(found_letters_so_far,byte(v))
        }
        
    }
    
    //fmt.Println("found_letters_so_far:", found_letters_so_far)
    //fmt.Println("repeats:", repeats)

    // 2nd pass to find the first letter that is not in REPEATS

    for i,v := range s {
        if is_in_slice(repeats, byte(v)) == false {
            ans = i
            break
        }
    }
    return ans
}
