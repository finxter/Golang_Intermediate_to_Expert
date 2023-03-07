// Typecasting
package main

import "fmt"

func main(){
    var i int
    var f float64

    i = int(f)
    f = float64(i)
    fmt.Printf("i=%d %T, f=%f %T\n", i, i, f, f)

    f = 45.66
    i = int(f)
    fmt.Printf("i=%d\n", i) // typecasting float to int will only assign the floor value

    // cannot convert between non convertible types
    // below will throw errors
    s := "typecasting"
    i= int(s)
    b := byte(s)

}