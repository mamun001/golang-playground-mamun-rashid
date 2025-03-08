// easy 1184, ELO 1234
// distance between bus stops

func distanceBetweenBusStops(distance []int, start int, destination int) int {

    total := 0 

    for i:=0;i<len(distance);i++ {
        total = total + distance[i]
    }

    front_distance := 0

    if start > destination {
        start,destination = destination,start
    }

    for i:=start;i<destination;i++ {
        front_distance = front_distance + distance[i]
    }

    back_distance := total - front_distance

    //fmt.Println(front_distance,back_distance)

    return min(front_distance,back_distance)
    
}
