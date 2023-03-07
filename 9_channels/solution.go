
// Exercise: channel
// Use print wherever necessary

/*
    1. In continuation of the search files example from goroutines.
    You need to search files on a disk. The function fileSearch()
    is already given to search for the files. It sleeps for 1 sec if
    the search is successful. 

    You should remove all the sleep() and instead introduce a
    channel to read/write the search result from main(). The search result
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

func fileSearch(dir string, fileName string, ch chan<- string ) {

    fmt.Println("[SEARCHING] ", dir) 
    files, err := ioutil.ReadDir(dir) 
    if err != nil {
        fmt.Println("error:", err)
        return
    } 
    for _, file := range files {
        if file.Name() == fileName {
            ch <- "[FOUND]" + filepath.Join(dir, file.Name())
            return
        }
    } 
    ch <- "[NOT FOUND]" + dir
}

func main()  {

    start  := time.Now()
    defer func ()  {
       fmt.Println(time.Since(start))
    }()

    ch := make(chan string)

    go fileSearch("../9_channels", "demo1.go", ch) // found

    fmt.Println(<-ch)

    go fileSearch("../8_goroutines", "demo1.go", ch) // not found

    fmt.Println(<-ch)
}
