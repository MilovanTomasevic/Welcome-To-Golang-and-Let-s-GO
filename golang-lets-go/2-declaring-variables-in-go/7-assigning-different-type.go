package main

import "fmt"

func main() {
  var number int = 10
  var price float64 = 15.10

  fmt.Println(number, price)

  price = (float64)number
  fmt.Println(price)
}
