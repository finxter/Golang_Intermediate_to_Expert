// Demo - Real use case of interface
// The io package in Go provides input-output primitives as well
// as interfaces to them. It is one of the most essential packages
// in Golang.

// The io package provides two very fundamental types the Reader and Writer.
// The reader provides a function that simply reads bytes from streams.
// The writer is writes bytes to the stream.
package main

import (
    "fmt"
    "io"
)

type pipe struct { 
    name string 
}
//implement the Read() interface for pipe
func (pipe) Read(b []byte) (int, error) {
    s := `ls -la` 
    copy(b, s) 
    return len(s), nil 
}

type file struct { 
    name string
}

//implement the Read() interface for file
func (file) Read(b []byte) (int, error) { 
    s := "<title>Learning Go at Finxter</title>"
    copy(b, s) 
    return len(s), nil 
}

// retrieve can read any device - file, pipe and process the data.
// Pass an io.Reader interface to the function. 
func retrieve(r io.Reader) error { 
    data := make([]byte, 100) 
    len, err := r.Read(data) 
    if err != nil { return err } 
    fmt.Println(string(data[:len])) 
    return nil 
}

func main() {
    p := pipe{"cfg_service"}
    f := file{"output.txt"}

    retrieve(p) // calls Read() of pipe
    retrieve(f) // calls Read() of file

    // Interface helps in implementing the polymorphism feature of 
    // OOP.
}
