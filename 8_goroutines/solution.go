// Exercise: goroutines
// Use print wherever necessary

/*
    1. You need to search files on disk. The function fileSearch()
    is already given to search for the files. It sleeps for 1 sec if
    the search is successful. 

    step1: First use the regular or normal function
    call to search a file in current directory and note the 
    time taken for the same.

    step2: Now use the goroutine call and note the time taken.
    Observe the difference and comment which one is faster ?

    Hint: You need to add time.Now() and time.Since() to get the time
    taken by main() and you can search for any file in the current
    directory.
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

    go fileSearch("../8_goroutines", "demo.go")

    // this is need when we use goroutine
     time.Sleep(time.Second * 2)
}


// With goroutine it takes 2 secs and normal call takes 3 secs.

// Go routines is faster