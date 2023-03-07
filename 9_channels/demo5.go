// demo: for range and close channel

// after close(ch) you can no longer write to the channel 
// but can read.

// use for ..range when 
// 1.) You're receiving a fixed number of values from a channel

// 2.) When you're waiting for a channel to close. If you're 
// using a channel to signal the completion of a process or 
// the termination of a goroutine, you can use a for loop with 
// range to wait for the channel to close

package main

import (
    "fmt"
)

func main() {
    // Create an unbuffered channel
    c := make(chan int)

    // Start a goroutine to send values to the channel
    go func() {
        for i := 1; i <= 5; i++ {
            c <- i
        }
        close(c) // close the channel when we're done sending values
    }()

    // Use a for loop and range to read values from the channel
    for v := range c {
        fmt.Println(v)
    }
}
