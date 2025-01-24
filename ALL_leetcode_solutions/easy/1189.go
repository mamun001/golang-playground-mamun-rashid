// easy 1189 maximum number of baloons
// ELO 1181
// Jan 23, 2025

func min (arr []int) int {
    min := 10001
    for i:=0; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }
    return min
}

func maxNumberOfBalloons(text string) int {
    freq := make([]int,5) // b a l o n

    for i:=0; i<len(text);i++ {
        if text[i] == 'b' {
            freq[0]++
        } else if text[i] == 'a' {
            freq[1]++
        } else if text[i] == 'l' {
            freq[2]++
        } else if text[i] == 'o' {
            freq[3]++
        } else if text[i] == 'n' {
            freq[4]++
        }
    }

    // normalize
    freq[2] = freq[2] / 2
    freq[3] = freq[3] / 2

    fmt.Println(freq)

    return min(freq)
    
    
  
    
}

