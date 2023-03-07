// Exercise init and defer
// Use prints wherever necessary

/*
    1.For the given init() function and a global variable 
    shouldIlearnGo, what will the main() print ? Try to understand
    the flow. 

    2. Uncomment the import "pack1/pkg" line and line 34:
    which prints pkg1.Name. Run the program and now observe how the 
    order in which init functions get called
*/

package main

import "fmt"
//import "pack1/pkg1"

var shouldILearnGo = learnGo()

func learnGo() int{
    return 0
}
func init() {
    fmt.Println("I am init in main")
    shouldILearnGo = 1
}

func main(){
    // Write your code below
    if shouldILearnGo == 1{
        fmt.Println("I am learning Go")
    }
    //fmt.Println("I am learning Go at:", pkg1.Name)
}