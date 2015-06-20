package main

import "fmt"

type Person struct {
  name string
  age int
}


func main() {
  fmt.Println(Person{"swagger", 20})
  bitch_made := Person{name: "jordan", age: 23}
  fmt.Println(bitch_made.name)
  pointify := &bitch_made
  pointify.age = 55
  fmt.Println(bitch_made.age)
  fmt.Printf("this is the content in pointify struct %b \n", pointify.age)
}
