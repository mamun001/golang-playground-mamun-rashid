// easy 997, ELO 1202
// find the Town judge


func findJudge(n int, trust [][]int) int {

   m := len(trust)

   if n == 1 {
     return 1
   }
   
   if m < n-1 { // not enough trusts that everyone can trust one person
    fmt.Println("not enough trust to go around")
    return -1
   }

   map_trusted := make(map[int]int) // count of trust for a given person
   map_truster := make(map[int]int) // count of how many people each person trusts
   for i:=0;i<m;i++{
     map_trusted[trust[i][1]]++
     map_truster[trust[i][0]]++
   }

   //fmt.Println(map_trusted)


   // see who is trusted the most and how many people trust him/her
   max_trusted_id:=1001
   max_trust:=0
   for k,v := range map_trusted {
      if v > max_trust {
        max_trusted_id = k
        max_trust = v
      }
   }
   //fmt.Println("max_trust",max_trust)
   //fmt.Println("max_trusted_id",max_trusted_id)

   // if max_trust is not n-1, not everyone trusts 1 person
   if max_trust < n-1 {
      //fmt.Println("not everyone trusts exactly 1 person")
      return -1
   }

   // find out how many people have the max trust count
   max_trusted_count := 0
   for _,v := range map_trusted {
      if v == max_trust {
        max_trusted_count++
      }
   }
   fmt.Println("how many people have max trust count",max_trusted_count)

   // more than 1 person has max trust
   if max_trusted_count > 1 {
      //fmt.Println("more than 1 possible judge")
      return -1
   }

   // the possible judge can't trust anyone
   if map_truster[max_trusted_id] > 0 {
      //fmt.Println("postential judge can't trusted anyone")
      return -1
   }

   
   if max_trusted_count == 1 {
      return max_trusted_id
   }

   

    return -1
    
}

