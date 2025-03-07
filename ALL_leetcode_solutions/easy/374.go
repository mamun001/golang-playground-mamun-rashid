// easy 374, part of needcode 250
// guess Number Higher or Lower

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return       -1 if num is higher than the picked number
 *                1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */



func guessNumber(n int) int {

    l:=1
    r:=n
    my_guess := (l+r)/2
    matched := false
    code := guess(my_guess)

    for matched == false {
        fmt.Println(l,r,my_guess,code)
        if code == 0 {
            matched = false
            return my_guess
        } else if code == 1 {
            l=my_guess+1
            my_guess = (l+r)/2
            code = guess(my_guess)
        } else if code == -1 {
            r=my_guess
            my_guess = (l+r)/2
            code = guess(my_guess)
        }
    }
    return my_guess
    
}
