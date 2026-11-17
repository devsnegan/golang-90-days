package main

import "fmt"

type BasicVehicle interface {
	Drive()
}

type ElectricVehicle interface {
	Charge()
}

type ElectricCar interface {
	BasicVehicle
	ElectricVehicle
}

type Tesla struct{}

func (t Tesla) Drive() {
	fmt.Println("Electric car is driving")
}

func (t Tesla) Charge() {
	fmt.Println("Electric car is charging")
}

func main() {
	var car ElectricCar = Tesla{}

	car.Drive()
	car.Charge()
}