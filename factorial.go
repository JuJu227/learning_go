package main

import (
  "fmt"
)

func main(){
result := fact(90)
fmt.Printf("this is the result of the fact call %d", result)
}

func fact(n int64 ) int64 {
  if n == 1 {
    return n
  } else { 
    return(n * fact(n-1)) 
  }
}
