
// Exercise: scope (unexported, exported)
// Use prints when necessary

/*
    1. In this file, import the animals package using (import pack1/animals).
    Create an object of type Lion. Try printing the object. See if this compiles.
    If no, try to understand why ? What changes will you make to
    compile it successful?

    2. Create an object of type Cat from the animals package.
    The main program can’t compile because we are trying to
    access the unexported type animal from inside the composite
    literal. How can you make this compilable and print the object.

    Hint: Creating an object of type Cat can be a bit tricky because
    of struct within struct
*/
package main

import "fmt"

import "pack1/animals"

func main(){
    // Write code below

    // solution 1
    lion := animals.Lion{
            Name:         "Scotch",
            RoarStrength: 10,
            Age:          5,
    }
    fmt.Println(lion)

    // solution 2
    cat := animals.Cat{
        Animal: animals.Animal{
            Name: "rege",
            Age: 2,
        },
        MeowStrength : 3,
    }
    fmt.Println(cat)
}