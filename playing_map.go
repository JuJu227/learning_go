package main

import  "fmt"

func main() {
  map_jazz := make(map[string]int)

  map_jazz["juju"] = 3
  map_jazz["malia"] = 9
  map_jazz["josef"] = 11

  fmt.Println("map:", map_jazz)
  fmt.Println(len(map_jazz))
  fmt.Println(map_jazz["josef"])

  map_jazz["steph"] = 33
  //map_jazz["cj"] = 331

  delete(map_jazz,"cj")

  fmt.Println(map_jazz)

  for k,v :=  range map_jazz{
    fmt.Printf("%s -> %b \n", k, v)
  }

}
