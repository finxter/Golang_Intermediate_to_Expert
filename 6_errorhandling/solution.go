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
import "errors"

var passwordLen = errors.New("password length >= 6")

func errorHandler(){

    if r:= recover(); r!=nil{
        fmt.Println("Error Recovery:", r)
    }
}

// Solution 1
func passwordLengthChecker(_password string) bool{

    defer errorHandler()

    if len(_password) < 6{
        // write code here
        panic("password length must be >=6")
    }
    return true
}
// Solution 2
func passwordLengthChecker2(_password string) (bool, error){

    if len(_password) < 6{
        // write code here
        return false, passwordLen
    }
    return true, nil
}
func main(){

    // writing some test code
    fmt.Println("checking password length")
    // solution 1
    var ok bool =  passwordLengthChecker("1234") // fail
    fmt.Println(ok)

    ok = passwordLengthChecker("secret") // Ok
    fmt.Println(ok)

    // solution 2

    ok, err := passwordLengthChecker2("1234")
    if err != nil{
        fmt.Println("Error:", err)
    }else{
        fmt.Println(ok)
    }

}