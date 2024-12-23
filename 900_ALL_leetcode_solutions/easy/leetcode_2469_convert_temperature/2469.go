// 2469 Convert the Temperature 
// Acceptance rate = 89.9
// DEC 22, 2024




func convertTemperature(celsius float64) []float64 {

    // initialize the ans slice with dummy values
    ans := []float64{1.1,2.2}

    // use the given formula to calculate K and F
    ans[0] = celsius + 273.15
    ans[1] = celsius * 1.80 + 32.00
    return ans
}
