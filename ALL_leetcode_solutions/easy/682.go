// 682 , easy, baseball game
// part of Neetcode 250
// using slice as stack

import "strconv"

func calPoints(operations []string) int {

    rec :=[]int{}

    for i:=0;i<len(operations);i++ {
        //fmt.Println(operations[i])
        if operations[i] == "C" {
            rec = rec[:len(rec)-1]
        } else if operations[i] == "D" {
            //fmt.Println("i",i)
            //fmt.Println("len of rec",len(rec))
            rec = append(rec,rec[len(rec)-1]*2)
        } else if operations[i] == "+" {
            last := len(rec)-1
            second_last := len(rec)-2
            rec = append(rec,rec[last]+rec[second_last])
        } else {
            num,_ := strconv.Atoi(operations[i])
            rec = append(rec,num)
        }
        //fmt.Println(rec)
    }

    //fmt.Println(rec)
    sum := 0
    for i:=0;i<len(rec);i++ {
        sum = sum + rec[i]
    }

    return sum
    
}
