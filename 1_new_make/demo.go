package main

import "fmt"

// make vs new

func main()  {
    // Data types

    // make works only with maps, slices and channels
    m := make(map[int]int, 10) // map
    s := make([]int, 1)        // slice
    b := make(chan bool)       // chan
    // _:= make(int, 3)  --- not possible

    // new can be used with any data type
    mp := new(map[int]int)   // mp points to a newly allocated zero map value
    ip := new(int)           // ip points to a newly allocated zero int value

    fmt.Println(m, s, b, mp, ip)

    // Return types

    // make returns target type T and new returns an address = pointer
    var v map[int]int = make(map[int]int)
    var p *map[int]int = new(map[int]int)
    fmt.Println(v, p)


    // Memory allocation

    // make initialises internal data structures and its ready to use
    memMake := make(map[int]int) // map
    fmt.Println(memMake)
    memMake[0] = 1

    // new allocates, but zero's memory ie zero's a map pointer (*p == nil)
    // This is pretty much useless
    memNew := new(map[int]int)
    fmt.Printf("%p\n",memNew)
    if *memNew == nil{
        fmt.Println("i am nil")
    }
    //(*memNew)[0] = 1  // crashes



   // Conclusion: Some rules for make() and new():

   // For slices, maps and channels: use make
   // For arrays, structs and all value types: use new
}