
// Exercise : receiver functions or struct methods

// Use print whenever necessary

/*
    1. Given an embedded struct(struct within struct) called Address,
    inside the Employee. Write a receiver method of Address struct
    to return the address in the form of string as below
    "street:%s, city:%s, postalcode:%s"
    Create an Employee object and print it. Now call the struct 
    method defined above and print it

    Hint: Use fmt.Sprintf() to return a string in the receiver
    method

    2. Create another receiver function of struct Employee.
    The function should receive a new string to correct the name
    of an employee. First use pass by value and then pass by 
    reference in this receiver function. Observe the changes
    by calling this function with a new name and then printing 
    it in the main.
*/
package main

import "fmt"

type Address struct{
    street string
    city string
    postalcode string
}

type Employee struct{
    name string
    Address
}


func (a Address) addrString() string{
    return fmt.Sprintf("street:%s, city:%s, postalcode:%s",
                    a.street, 
                    a.city,
                    a.postalcode)
}

func (e *Employee) correctName(name string) {
    e.name = name 
}
func main(){
    // Your code below
    // Solution 1
    addr := Address{
            street: "MGRoad",
            city : "Bengaluru",
            postalcode: "560001",
    }
    emp := Employee{
            name: "Shiv",
            Address: addr,
    }
    fmt.Println(emp)
    fmt.Println(emp.addrString())

    // Solution 2
    emp.correctName("Shivaram")
    fmt.Println(emp)
}

