package main
import "fmt"
func tiger() (int,int,string){
  swigger := "this is cold blood"
  return 3,5,swigger
}

func bae() (string,int) {
  return "juju", 3
}

func main() {
  a,b,c := tiger()
  fmt.Printf("%d %d %s",b,a,c)
  f,_ := bae()

  fmt.Printf("\n %s \n",f)
}

