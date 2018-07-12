/*
  Problem: Golang program to check whether a number is palindrome or not
*/

package main
import "fmt"

func main() {
  var number, remainder int
  var reverse int = 0
  var negative bool

  fmt.Print("Enter any integer Number : ")
  fmt.Scan(&number)
  temp := number
  if number < 0 {
    number = number * -1
    negative = true
  }

  for {
    remainder = number % 10
    reverse = reverse * 10 + remainder
    number /= 10

    if(number==0){
        break
    }
  }
  if negative {
    reverse *= -1
  }
  if reverse == temp {
    fmt.Printf("%d is a Palindrome\n",temp)
  } else {
    fmt.Printf("%d is not a Palindrome\n",temp)
  }
}

/******************************************

OUTPUT:

$ go run palindrome_02.go

Enter any integer Number : -121
-121 is a Palindrome

$ go run palindrome_02.go

Enter any integer Number : 231
231 is not a Palindrome

******************************************/