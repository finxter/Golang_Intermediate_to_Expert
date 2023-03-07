// demo : more channel concepts

// demo2 :  Pass channel as params
// demo3 :  Buffered channels
// demo4 : Select keyword
// demo5 : for range and close on channels

package main

import "fmt"


// goroutine usually either writes or reads but not both

// here it writes (channel param as write)
func worker1(ch chan<- int){ 

    ch <- 1
}

// here it reads (channel param as read)
func worker2(ch <-chan int){

    fmt.Println("reading on channel:",<-ch)
}

func main ()  {

    c1 := make(chan int)
    c2 := make(chan int )
    go worker1(c1)
    fmt.Println(<-c1) // read on c1
    go worker2(c2)
    c2 <- 15   // write to c2
}

