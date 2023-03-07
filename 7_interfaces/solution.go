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

/* -------------- step 1 --------------------*/
// interface for sqlDatastore
type Datastore interface {
    Save(data []byte) error
}

// interface for consoleLogger
type Logger interface {
    Log(message string)
}

/* -------------- step 2 --------------------*/

// New struct now includes the interfaces that were defined
type StoreLogger struct {
    ds Datastore
    l  Logger
}

func (sl *StoreLogger) SaveSomething(data []byte) {
    if err := sl.ds.Save(data); err != nil {
        sl.l.Log("Failed to save data: " + err.Error())
        return
    }

    sl.l.Log("Data saved successfully.")
}

func main() {
    // Create instances of dependencies
    datastore := &sqlDatastore{}
    logger := &consoleLogger{}

    // Inject dependencies into the StoreLogger instance
    sl := &StoreLogger{
        ds: datastore,
        l:  logger,
    }

    // Use the StoreLogger instance to perform some task
    sl.SaveSomething([]byte("save with StoreLogger some data"))
}

// DI

// Interface package  
// package A
// package B
