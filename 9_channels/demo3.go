// Buffered channels


package main

import (
    "fmt"
)

func main() {
    // Create a buffered channel that can hold up to 2 values

    c := make(chan int, 2)
    // same as c := make(chan int, 0 )

    // Send some values to the channel
    c <- 1 // it waits for you to consume
    c <- 2
    // Read the values from the channel
    fmt.Println(<-c)
    fmt.Println(<-c)
    fmt.Println(<-c)
}


/*

main()                channel                 another
goroutine                                     goroutine
==========================================================
receives     <-----    msg       <-----      sends   

                --->       [1, 2]
  [1, 2]              <---[]
  waiting to consume
*/

// In order to prevent any deadlocks
// 1. Choose appropriate size of the buffer for the channel
// 2. For a channel, number of receivers should not be > number of senders
// 3. For a channel, number of senders should not be > buffer size
