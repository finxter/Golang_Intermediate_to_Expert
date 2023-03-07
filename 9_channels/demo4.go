// demo for using select when we have multiple goroutines and
// receive multiple values from goroutines.
package main

import (
    "fmt"
    "time"
)

func worker1(c1 chan string) {
    for {
        time.Sleep(time.Second * 1)
        c1 <- "Message from worker1" // send msg on chan 1
    }
}

func worker2(c2 chan string) {
    for {
        time.Sleep(time.Second * 2)
        c2 <- "Message from worker2" // send msg on chan 2
    }
}

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go worker1(c1)
    go worker2(c2)

    for {
        select {
        case msg1 := <-c1: // receive on chan1
            fmt.Println(msg1)
        case msg2 := <-c2:   // receive on chan2
            fmt.Println(msg2)  
        }
    }
}
