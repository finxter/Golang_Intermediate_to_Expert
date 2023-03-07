
// Exercise error handling

// Use print statements wherever needed and use meaningful
// function, variable names

/*
    1. Given a passwordLengthChecker() function which accepts a string for 
    password. Write panic() to crash the program if the length 
    of the password is less than six.
    Now write a recover function to ensure that the program
    does not crash.
    Hint : use recover() to avoid crash.

    2. Use error pattern (i.e. create error object using errors
    package) to handle the error instead of panic/recover. No need
    to log to an external file, just print it to the console.
*/


package main

import "fmt"

func passwordLengthChecker(_password string) bool{
    if len(_password) < 6{
        // write code here
    }
    return true
}
func main(){
    fmt.Println("hello")
    // write code below to test

}