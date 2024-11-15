package main

import "fmt"

// Car interface defines the methods that any Car object must implement.
type Car interface {
	Cost() int
	Description() string
}

// Concrete Component: SimpleCar represents a basic Car.
type SimpleCar struct{}

func (s SimpleCar) Cost() int {
	return 10
}

func (s SimpleCar) Description() string {
	return "Simple Car"
}

// Decorator: Sunroof adds sunroof to the car.
type Sunroof struct {
	car Car
}

func (c Sunroof) Cost() int {
	return c.car.Cost() + 5
}

func (c Sunroof) Description() string {
	return c.car.Description() + " with Sunroof"
}

// Decorator: AC adds ac to the car.
type AC struct {
	car Car
}

func (c AC) Cost() int {
	return c.car.Cost() + 7
}

func (c AC) Description() string {
	return c.car.Description() + " with AC"
}

func main() {
	// Create a SimpleCar object.
	simple := SimpleCar{}

	// Decorate with AC.
	acCar := AC{car: simple}
	fmt.Println(acCar.Description(), acCar.Cost())

	// Decorate with Sunroof.
	sunroofCar := Sunroof{car: simple}
	fmt.Println(sunroofCar.Description(), sunroofCar.Cost())

	// Decorate with both AC and Sunroof.
	acAndSunroofCar := Sunroof{car: acCar} // Chain decorators
	fmt.Println(acAndSunroofCar.Description(), acAndSunroofCar.Cost())
}
