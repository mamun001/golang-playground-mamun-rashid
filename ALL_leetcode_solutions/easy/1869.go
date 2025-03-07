// easy 1869, ELO 1204

func checkZeroOnes(s string) bool {

    if s == "1" {
        return true
    }

    if s == "0" {
        return false
    }

    max_one_seg := 0
    max_zero_seg := 0
    this_one_seg := 0
    this_zero_seg := 0
    last := '9'
    this := '9'


    if s[0] == '0' {
        max_zero_seg,this_zero_seg = 1,1  
    }

    if s[0] == '1' {
        max_one_seg,this_one_seg = 1,1  
    }

    for i:=1;i<len(s);i++ {
      last = rune(s[i-1])
      this = rune(s[i])

      if last == '0' && this == '1' {
           this_one_seg = 1
           if this_one_seg > max_one_seg {
            max_one_seg = this_one_seg
           }
      } else if last == '0' && this == '0' {
           this_zero_seg++
           if this_zero_seg > max_zero_seg {
            max_zero_seg = this_zero_seg
           }
      } else if last == '1' && this == '0' {
           this_zero_seg = 1
           if this_zero_seg > max_zero_seg {
            max_zero_seg = this_zero_seg
           }
      } else if last == '1' && this == '1' {
           this_one_seg++
           if this_one_seg > max_one_seg {
            max_one_seg = this_one_seg
           }
      }
    }


    //fmt.Println(max_one_seg,max_zero_seg)

    return max_one_seg > max_zero_seg
    
}
