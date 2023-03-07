// Exercise new vs make

// Use prints wherever necessary
/*
    1.) For a given struct use new() to allocate the struct. Initialise the string to 
    "hello gopher" and int to 100. Try using make() on the struct .It should not compile.

    2.) Use make() to initialise a slice of int with length 10. Then update the elements
    with new values (say 1,2,3).Try using new() for slice and it should return an address 
    that points to an empty or nil slice. Confirm this using if statement.

*/



package main

import "fmt"

type Bar struct {
    s string
    i int
}
func main(){

    // Write your code below.

    b := new(Bar)
    fmt.Println("before:",b)
    (*b).s = "hello gopher"
    (*b).i = 100
    fmt.Println("after:",b)

    //m := make(Bar)

    s := make([]int, 10)
    s[0] = 1
    s[1] = 2
    s[2] = 3
    fmt.Println(s)

    np := new([]int)
    fmt.Println(np)
    if *np == nil{
        fmt.Println("np is nil")
    }
}