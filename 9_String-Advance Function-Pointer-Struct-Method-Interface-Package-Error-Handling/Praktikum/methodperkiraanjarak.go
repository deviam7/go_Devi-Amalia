package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) EstimateDistance() float64 {
	return car.FuelIn * 1.5
}

func main() {
	myCar := Car{Type: "Sedan", FuelIn: 45.0}

	estimatedDistance := myCar.EstimateDistance()

	fmt.Printf("My %s can travel approximately %.2f miles with the current fuel.\n", myCar.Type, estimatedDistance)
}
