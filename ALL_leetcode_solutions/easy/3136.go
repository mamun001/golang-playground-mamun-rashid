func is_digit(b byte) bool {
    if b > 47 && b < 58 {
        return true
    }

    return false

}


func is_vowel(b byte) bool {
    if b == 65 || b == 69 || b == 73 || b == 79 || b == 85 {
        return true
    }

    if b == 97 || b == 101 || b == 105 || b == 111 || b == 117 {
        return true
    }
    
    return false

}

func is_valid_char(b byte) bool {
    if b > 47 && b < 58 {
        return true
    }

    if b > 64 && b < 91 {
        return true
    }

    if b > 96 && b < 123 {
        return true
    }

    return false

}

func is_conso(b byte) bool {

    if is_valid_char(b) && is_digit(b) == false && is_vowel(b) == false {
        return true
    }

    return false

}



func isValid(word string) bool {

    // A-Z 65-90
    // a-z 97-122
    // 0-9 48-57
    // a e i o u A E I O U
    // 65,69,73,79,85 97,101,105,111,117

    if len(word) < 3 {
        return false
    }


    vowel,conso := 0,0
    for i:=0;i<len(word);i++ {
      //fmt.Println(byte(word[i]))
      //fmt.Println(is_valid_char(word[i]))

      if is_valid_char(word[i]) == false {
        return false
      }
      if is_vowel(word[i]) {
        vowel++
      }
      if is_conso(word[i]) {
        conso++
      }
    }

    if vowel < 1 {
        return false
    }

    if conso < 1 {
        return false
    }

    return true
    
}
