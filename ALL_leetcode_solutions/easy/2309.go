// easy 2309, ELO 1243
// Greatest English Letter in Upper and Lower Case

func greatestLetter(s string) string {

    // a=97
    // z=122
    // A=65
    // E=69
    // Z=90
    // diff = 97-65=32

    // save frequencies of each letter
    map_letters := make(map[byte]int)

    for i:=0;i<len(s);i++ {
            map_letters[s[i]]++
    }

    //fmt.Println(map_letters)


    // start at Z and see if both lower case and upper case exists, then  return that
    for i:=90;i>64;i--{
        if map_letters[byte(i)] > 0 && map_letters[byte(i+32)] > 0 {
            return string(i)
        }
    }

    // if there is no such pair
    return ""
    
}
