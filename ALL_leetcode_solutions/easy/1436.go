// easy 1436, ELO 1192
// destination city

// NOT a graph problem
// solved using arry and hashmap: CPU > 100% RAM> 94%!

func is_in (a []string, s string) bool {
    ans := false
    for i:=0 ; i< len(a) ; i++ {
        if s == a[i] {
            return true
        }
    }

    return ans
}

func destCity(paths [][]string) string {

    ans := "a"

    if len(paths) == 1 {
        return paths[0][1]
    }

    starts := []string {}

    for i:=0 ; i< len(paths) ; i++ {
        starts = append (starts,paths[i][0])
    }


    for i:=0 ; i< len(paths) ; i++ {
        //fmt.Println(path)
        if is_in(starts, paths[i][1]) == false {
            fmt.Println("found")
            return paths[i][1]
        }
    }

    return ans
}

