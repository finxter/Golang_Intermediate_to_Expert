package main

import "fmt"
import "os"
import "io"


func main(){
    file, err := os.Open("defer.txt") 
    if err != nil { 
        fmt.Println(err) 
        return 
    }

    defer file.Close() // defer the close() 

    // create a byte slice to hold the file contents
    data := make([]byte, 1024)
    for {
        n, err := file.Read(data)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("File reading error", err)
            return
        }
        fmt.Println("Read", n, "bytes:", string(data[:n]))
    }
}
