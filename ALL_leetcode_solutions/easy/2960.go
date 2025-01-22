// easy 2960, ELO 1169
// Count Tested Devices
// JAN 20, 2025


func countTestedDevices(batteryPercentages []int) int {

    ans := 0
    n := len(batteryPercentages)

    for i:=0; i<n; i++ {
        if batteryPercentages[i] > 0 {
            ans++
            for j:=i+1; j<n; j++ {
                if batteryPercentages[j] > 0 {
                   batteryPercentages[j] = batteryPercentages[j] - 1
                }
            }
        }
    }

    return ans
    
}

