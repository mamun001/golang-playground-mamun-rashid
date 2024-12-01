//
// Example use of fscanf

package main

import (
  "fmt"
)

func main() {

   var first_name string // we need this var to store what we read via fscanf

   fmt.Print("Enter First Name: ")
   fmt.Scanf("%s", &first_name) // read via fscanf STDIN and store . NOTE & is a pointer
   fmt.Print("Your first name is: ")
   fmt.Println(first_name)
}
