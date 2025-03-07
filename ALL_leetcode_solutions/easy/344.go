// easy 344, first one on NC list for 2 pointers, no ELO
// reverse string

func reverseString(s []byte)  {

    l := 0
    r := len(s)-1
    //fmt.Println(r)

    for l<=r {
        //fmt.Println(l,r)
        s[l],s[r] = s[r],s[l]
        l++
        r--
    }
}
