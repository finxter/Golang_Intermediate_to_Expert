// Error handling

//    2. error pattern ( create an error object).


package main

import ("fmt"
        "log"
        "runtime/debug"
        "errors"
        "os"
)

var noZero = errors.New("divider should not be zero")

func Divide(nominator int, divider int) (float32, error){

    if divider == 0 {
        log.Println(string(debug.Stack()))
        return 0, noZero
    } 
    return float32(nominator) / float32(divider), nil
}

func main() {
    // step1 : Set a log file by opening a file.
    // step2 : call the divide() function with 0
    // step3 : Implement Divide() function by handling 0 and logging stack trace
    // step4 : evaluate return values and log err if any to the file.

    f, err := os.OpenFile("err_logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil { 
        log.Println(err) 
    }
    log.SetOutput(f)

    fmt.Println("start")

    // error case
    n, err := Divide(10, 0)
    if err != nil{
        log.Println("ERROR:", err)
    }else{
        fmt.Println(n)
    } 

    // normal case
    n, err = Divide(10, 2)
    if err != nil{
        log.Println("ERROR:", err)
    }else{
        fmt.Println(n)
    } 

    fmt.Println("end")
}