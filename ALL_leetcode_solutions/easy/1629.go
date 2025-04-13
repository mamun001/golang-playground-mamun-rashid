// easy 1629, ELO 1315
// slowest key
// March 24

func slowestKey(releaseTimes []int, keysPressed string) byte {

    duration := -1
    max := 0
    var winner byte

    for i:=0;i<len(releaseTimes);i++ {
        if i==0 {
            duration = releaseTimes[0]
        } else {
            duration = releaseTimes[i] - releaseTimes[i-1]
        }
        if duration > max {
            max = duration
            winner = byte(keysPressed[i])
        }
        if duration == max && byte(keysPressed[i]) > winner {
            max = duration
            winner = byte(keysPressed[i])
        }


    }


    return winner
    
}
