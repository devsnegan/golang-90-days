package main

import "fmt"

type Vehicle interface {
    Start()
}

type Car struct{}

func (c Car) Start() {
    fmt.Println("Car started")
}

func main() {
    var value any = Car{}

    vehicle, ok := value.(Vehicle)

    if ok {
        vehicle.Start()
    }
}