package main

import "fmt"

func plus(a int, b int) int {
 return a+b
}

func main() {
  res := plus(1,2)
  fmt.Println(res)
  fmt.Println(fib(10000,1,0))
  fmt.Println(fact(55))
}

func fib(k int,a int,b int) int {
 if k == 0 {
  return a
 }else if k == 1 {
  return b
 }else{
   return(fib(k-1, b, a+b))
 }
}

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
