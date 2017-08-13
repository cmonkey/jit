package main

import (
    "fmt"
    "time"
    "runtime"
)

func sleep(s string){
    var i int = 0
    for{
        i += 1
        fmt.Println(i)
        fmt.Println(s)
        time.Sleep(time.Second)
    }
}

func main(){
    fmt.Println(runtime.GOMAXPROCS(0))
    go sleep("hello")
    go sleep("world")

    time.Sleep(1000 * time.Second)
}
