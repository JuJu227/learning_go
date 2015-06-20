package main

import("os/exec")

func main(){
    cmd := exec.Command("say my name is golang!")
    cmd.Run()

}
