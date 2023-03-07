// Exercise : interface

/*
    In this exercise we will see how interfaces 
    can be used to perform dependency injection (DI).
    What is DI ?
    Is an excellent way to decouple your code. It is
    a design pattern used in software engineering that
    helps to achieve loose coupling between components 
    by allowing components to receive their dependencies 
    from external sources, rather than creating them internally.

    Problem: 
    step1: You are given two structs sqlDatastore and consoleLogger and their
    struct methods (receiver functions). You job is to create two
    interfaces with methods (choose appropriate names) which represent these
    structs methods. 

    step2: Create another struct (choose appropriate name) and which
    will include these interface types that were defined in step1 as
    its field. Implement a struct method using reference/pointer type
    of this struct. You should now be able to call the struct methods of
    sqlDatastore and consoleLogger from this struct you created.

    Make a small conclusion as to how DI was implemented in this case?

    Hint: With interface the struct in step2 will not have to know the 
    details of struct methods of step1. i.e. no dependency 

*/
package main
import "fmt"

type sqlDatastore struct{}

func (sq *sqlDatastore) Save(data []byte) error {
    // Save data to MySQL database
    // Do not write anything in this function
    fmt.Println("Saving done:", string(data) )
    return nil
}

type consoleLogger struct{}

func (cl *consoleLogger) Log(message string) {
    fmt.Println(message)
}

// Write your code below

func main(){

    // Create instances of dependencies
    datastore := &sqlDatastore{}
    logger := &consoleLogger{}

    // Write your code below

}