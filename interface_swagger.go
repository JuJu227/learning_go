package main

import "fmt"

type Person struct {
  first_name string
  last_name  string
  age int
  family * Person
}

type Address struct {
  street_number int
  street_name string
  city string
  zip_code string
}

type Magic interface {
   jazz()
}

func main() {
    c := new(Person)
    c.first_name = "S"
    c.last_name = "jones"
    c.age = 33
    c.family = new(Person)
    go fmt.Println(c.first_name )
    go fmt.Println(c.last_name)
    go fmt.Println(c.age)
    c.family.first_name = "steph"
    fmt.Println(c.family.first_name)
    c.family.Talk()
 //fmt.Printf ("this should be a pointer %p \n", c)
}


func (p *Person) Talk() {
  fmt.Println(" what's there name", p.first_name)
}


