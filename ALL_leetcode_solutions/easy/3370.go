// east 3370 , ELO 1198
// smallest number with all set bits

func smallestNumber(n int) int {

    // basically the answer to 2^N-1 
    // for any given number, we add 1
    // then we keep diving by 2
    // if we ever get a reminder, then that's not it
    // after diving forever, if we get to 1, then we have it

   // lazy programming

   if n == 1 {
    return 1
   }
   if n >1 && n <= 3 {
    return 3
   }
   if n >3 && n <= 7 {
    return 7
   }
   if n >7 && n <=15 {
    return 15
   }
   if n >15 && n <=31 {
    return 31
   }
   if n >31 && n <=63 {
    return 63
   }
   if n >63 && n <=127 {
    return 127
   }
   if n >127 && n <=255 {
    return 255
   }
   if n >255 && n <=511 {
    return 511
   }
   if n >511 && n <=1023 {
    return 1023
   }
   

    return 9999
    
}

