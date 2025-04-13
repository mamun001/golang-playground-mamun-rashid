// easy 2325, no ELO, acc 85%, but felt more difficult than that
// decode the message

import "fmt"

func decodeMessage(key string, message string) string {

    //"the quick brown fox jumps over the lazy dog"
    // a=97
    // b=98
    // c=99
    // d=100
    // e=101=99
    // h=104
    // k=107
    // n=110
    // q=113=d=100
    // t=116=97
    // w=119=
    // z=122

    decode_map :=make(map[byte]byte)
    decode_map[' ']=' '
    //fmt.Println(decode_map)

    pointer:='a'
    for i:=0;i<len(key);i++{
        _,exists := decode_map[key[i]]
        if exists == false {
            decode_map[key[i]]=byte(pointer)
            pointer=pointer+1
            //fmt.Println(rune(pointer))
        }
    }
    //fmt.Println(decode_map)

    ans:=""
    for i:=0;i<len(message);i++{
        ans=ans+string(decode_map[message[i]])
    }


    return ans
    
}
