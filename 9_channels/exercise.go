
// Exercise: channel
// Use print wherever necessary

/*
    1. In continuation of the search files example from goroutines.
    You need to search files on a disk. The function fileSearch()
    is already given to search for the files. It sleeps for 1 sec if
    the search is successful. 

    You should remove all the sleep() and instead introduce a
    channel to read/write the search result from main().The search result
    must include "FOUND"/"NOT FOUND" + path of the file. Also adjust
    the print statements accordingly by moving it to main().


*/

package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "time"
)

func fileSearch(dir string, fileName string) {

    fmt.Println("[SEARCHING] ", dir) 
    files, err := ioutil.ReadDir(dir) 
    if err != nil {
        fmt.Println("error:", err)
        return
    } 
    for _, file := range files {
        if file.Name() == fileName {
            time.Sleep(time.Second)
            fmt.Println( "[FOUND] " + filepath.Join(dir, file.Name()) )
            return
        }
    } 
    fmt.Println("[NOT FOUND] " + dir)
}

func main()  {

    start  := time.Now()
    defer func ()  {
       fmt.Println(time.Since(start))
    }()

    go fileSearch("../9_goroutines", "demo1.go")

    // this is need when we use goroutine
     time.Sleep(time.Second * 2)
}
