// This demo will explain the importance and advantages
// of having an interface with a simple example

package main

import "fmt"
import "reflect"


// Without Interfaces

type Car struct{ speed int }
type Bus struct{ speed int }
type Bike struct { speed int}

// create an interface for Car and Bus

type VehicleSpeed interface {
    Speed() int
    setSpeed() 
}

func (c Car) Speed() int { return c.speed }
func (s *Car) setSpeed() { s.speed = 100 } 

func (b Bus) Speed() int { return b.speed }
func (s *Bus) setSpeed() { s.speed = 50 }

func (b Bike) Speed() int { return b.speed }
func (b *Bike) setSpeed() { b.speed = 200 }

// Lets add a helper function to print the speed of vehicles
func getVehicleSpeed(v VehicleSpeed) {
    fmt.Println(reflect.TypeOf(v).Elem().Name(),"runs at:", v.Speed())
}

func main() {
    // Without interface
/*
    c := Car{}
    c.setSpeed()
    fmt.Println("Car speed:", c.Speed())
    b := Bus{}
    b.setSpeed()
    fmt.Println("Bus speed:", b.Speed())
    bk := Bike{}
    bk.setSpeed()
    fmt.Println("Bike speed", bk.Speed())

    getCarSpeed(c)
    */
// With interface
    c := &Car{}
    b := &Bus{}
    bk := &Bike{}
    fmt.Println("using interfaces:")

    vehicles := []VehicleSpeed{c, b}
    vehicles = append(vehicles, bk)
        for _,v  := range vehicles {
        v.setSpeed()
        fmt.Println(reflect.TypeOf(v).Elem().Name(), "speed:", v.Speed())
        getVehicleSpeed(v)
    }

}
