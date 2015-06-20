package main

import "fmt"

func main(){
  var  x interface{}  = 7
  i :=  x.(x)

  type I  interface { m() }
  var y I
   s := y.(string)
   r :=  y.(io.Reader)

   fmt.Printf("this is what in s %s, r -> %s", s, r)

}
