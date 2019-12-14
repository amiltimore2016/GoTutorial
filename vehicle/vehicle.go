package main

import (
	"fmt"
)

type Car string

func (c Car) Accelerate() {
	fmt.Println("Car is accelerating")
}

func (c Car) Brake() {
	fmt.Println("Car is braking")
}

func (c Car) Steer(direction string) {
	fmt.Println("Car is steering ", direction)
}

type Truck string

func (t Truck) Accelerate(truckb string) {
	fmt.Println(truckb, " is accelerating")
}

func (t Truck) Brake(truckb string) {
	fmt.Println(truckb, " is braking")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Truck is steering ", direction)
}

type Vehicle interface {
	Accelerate(string)
	Steer(string)
	Brake(string)
}

func main() {
	var TheTruck Vehicle
	TheTruck = Truck("ford")
	TheTruck.Accelerate("ford")
}
