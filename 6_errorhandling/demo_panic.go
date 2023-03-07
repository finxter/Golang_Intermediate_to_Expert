// Error handling

/*
    1. panic / recover (either crash the program with panic or 
    use recover() to make a recovery i.e. avoid the crash).
*/

package main
import "fmt"

func errorHandler() { 
    // if r is nil then there is no error and user notices nothing,
    // otherwise if its not nil then we have an error. 
    if r := recover(); r != nil { 
        fmt.Println("Recovered ", r) 
        }
}

func Divide(nominator int, divider int) float32{
    defer errorHandler()
    if divider == 0 {
        panic("can't divide by 0") 
    } 
    return float32(nominator) / float32(divider) 
}

func main() {
    n := Divide(10, 0)
    fmt.Println(n)
}

