package main

import ("fmt"; "time"; "reflect")

func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    for {
      c1 <- "is Malia a meanie"
      time.Sleep(time.Second * 2)
    }
  }()

  go func() {
    for {
      c2 <- "Yes!!!"
      time.Sleep(time.Second * 3)
    }
  }()

  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      }
    }
  }()

  var input string
  chan_swagger := make(chan  string)
  go MaliaIsMean(chan_swagger)
  juju_is_bad := <- chan_swagger
  fmt.Println(juju_is_bad)
  fmt.Scanln(&input)
  name_swagger :=  make(chan struct{})
  go tacos(name_swagger)
  nana := <- name_swagger
  fmt.Println(nana)
}

func tacos(beans chan Jazzy){
  beans.nana <- "which me wip"
}

type Jazzy struct {
  nana string
}

func MaliaIsMean(really chan string) {
  for {
     fmt.Println(reflect.TypeOf(1))
    really <-  "yes, really!!!"
  }
}

