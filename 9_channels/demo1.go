
// demo - goroutines (simple example)

// Now modified with channels so that we can
// now remove the sleep and instead use a channel
// to send a message
package main

import (
    "fmt"
    "time"
)

func broadcastInfo(info, name string, signal chan  string){
    fmt.Println("Broadcasting", name,":", info)
    signal <- "Done"  // send msg on the chan
}

func main(){
    start := time.Now()
    defer func ()  {
        fmt.Println(time.Since(start))
    }()

    // start a new goroutine using go keyword

    broadcasters := []string{"CNN", "FirstPost", "NYTimes" }

    broadcastSignal := make(chan string)// chan [message type]

    for _, broadcaster := range broadcasters {
        go broadcastInfo("I am learning Go", broadcaster, broadcastSignal )
        fmt.Println(<- broadcastSignal) // receive msg on the chan

    }
}
