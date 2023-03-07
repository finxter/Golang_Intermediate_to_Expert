// demo - goroutines (simple example)
package main

import (
    "fmt"
    "time"
)

func broadcastInfo(info, name string){
    time.Sleep(time.Second)
    fmt.Println("Broadcasting", name,":", info)
}

func main(){
    start := time.Now()
    defer func ()  {
        fmt.Println(time.Since(start))
    }()

    // start a new goroutine using go keyword

    broadcasters := []string{"CNN", "FirstPost", "NYTimes" }
    for _, broadcaster := range broadcasters {
        go broadcastInfo("I am learning Go", broadcaster)
    }
    time.Sleep(time.Second * 2) // delay or sleep for 2 secs
}
