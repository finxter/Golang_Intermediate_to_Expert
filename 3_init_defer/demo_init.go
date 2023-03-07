
package main

import "fmt"

func init(){
    fmt.Println("init 2")
}

func main(){
    fmt.Println("main")
}

func init(){
    fmt.Println("init 3")
}
func init(){
    fmt.Println("init 1")
}


