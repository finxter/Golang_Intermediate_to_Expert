// Scope of Variables

// Local
// Unexported
// Exported

package main

import "fmt"
import "pack1/pkg1"

var myvar = 4 // unexported variable

func main(){

    // Unexported
    fmt.Println(myvar)

    // Local or Block scope
    {
        myvar := 1
        fmt.Println(myvar)
    }

    // Unexported
    fmt.Println(myvar)

    //fmt.Println(pkg1.name)

    // Exported

    fmt.Println(pkg1.Name)

}