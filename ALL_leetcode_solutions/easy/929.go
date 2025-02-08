// easy 929, ELO 1198
// Unique Email Addresses

import "strings"


func is_in(arr []string, s string) bool {
    ans := false
    for i:=0;i<len(arr);i++ {
        if arr[i] == s {
            return true
        }
    }
    return ans
}

func uniq(s string) string {
    two_parts := strings.Split(s,"@")
    //fmt.Println(two)
    user_part := strings.ReplaceAll(two_parts[0],".","")
    clean_user := strings.Split(user_part,"+")
    //fmt.Println(a)
    //fmt.Println(b[0])
    return clean_user[0]+"@"+two_parts[1]
}

func numUniqueEmails(emails []string) int {

    //ans := 0
    clean_emails :=[]string{}

    for i:=0;i<len(emails);i++ {
       s := uniq(emails[i])
       //fmt.Println(a)
       if is_in(clean_emails,s) == false {
          clean_emails = append(clean_emails,s)
       }
    }
    return len(clean_emails)
    
}

