// 2011 Final Value of Variable After Performing Operations
// 89.5 acceptance rate

// Problem definition looks scary, but the problem is super simple
// Dec 22, 2024


func finalValueAfterOperations(operations []string) int {
    fmt.Println(operations[0])
    x := 0

    for _,v := range operations{
        
        switch v {
            case "--X":
            x--
            case "X--":
            x--
            case "X++":
            x++
            case "++X":
            x++
        }
    }
    return x
}
