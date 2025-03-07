// easy 

// 1693 Design Parking System

// 88.0 pct acc rate

// DEC 30, 2024

// classes, constructurs , methods and state in GO!


type ParkingSystem struct {
    big_count int
    medium_count int
    small_count int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{big_count: big, medium_count: medium, small_count: small}
}


func (this *ParkingSystem) AddCar(carType int) bool {

    ans := true

    if carType == 1 {
       if this.big_count < 1 {
            ans = false 
       } else {
            ans = true
            this.big_count--
       } // else  
    } // 1 , big

    if carType == 2 {
       if this.medium_count < 1 {
            ans = false 
       } else {
           ans = true
           this.medium_count--
       } // else  
    } // 1 , medium


    if carType == 3 {
       if this.small_count < 1 {
            ans = false 
       } else {
            ans = true
            this.small_count--
       } // else  
    } // 1 , small

   return ans

} // func



/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */


