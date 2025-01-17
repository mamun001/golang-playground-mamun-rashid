/ 2894 divisible and non divisible sums difference

// acc rate 89.0 %

// DEC 24, 2024



func sum_of_non_divisibles(num1 int, m int) int {
    

    if num1 == 1 && m != 1 { 
        return 1 // not divisible by m
    }

    if num1 == 1 && m == 1 { 
        return 0 // divisible by m
    }

    // not divisible by m
    if num1 % m != 0 {
        return num1 + sum_of_non_divisibles (num1-1, m) // keep recursing until you get to 1
    }

    // divisible, so dont add
        return sum_of_non_divisibles (num1-1, m) // keep recursing until you get to 1
    
}


func sum_of_divisibles(num1 int, m int) int {
    

    if num1 == 1 && m != 1 { 
        return 0 // not divisible by m
    }

    if num1 == 1 && m == 1 { 
        return 1 // divisible by m
    }

    // not divisible by m
    if num1 % m != 0 {
        return sum_of_divisibles (num1-1, m) // keep recursing until you get to 1
    }

    // divisible, so dont add
        return num1 + sum_of_divisibles (num1-1, m) // keep recursing until you get to 1
    
}



func differenceOfSums(n int, m int) int {
    //return sum_of_non_divisibles(10,3)
    // return sum_of_divisibles(5,6)
    return sum_of_non_divisibles(n,m) - sum_of_divisibles(n,m)
    
}

