package main

import "fmt"

func loyal (nums ...int) int {
  return len(nums)
}

func main() {
 fmt.Println(loyal(1,2,3,4))
 fmt.Println(loyal(1))
}

