// 
// The code you've provided demonstrates a great example of the Builder Pattern in Go. This pattern is typically used to construct complex objects step by step, allowing for different representations of the same type of object. In this case, you're building different types of cars (SportsCar, ElectricCar, and FamilyCar), each with a unique configuration of components.

// Key Components of the Builder Pattern:
// Product (Car): The complex object you're building. In this case, Car has various attributes like Engine, Color, Wheels, Seats, etc.



// Concrete Builders: These classes (SportsCarBuilder, ElectricCarBuilder, FamilyCarBuilder) implement the CarBuilder interface and provide specific implementations for constructing different types of cars.

// Director: The Director is responsible for managing the building process. It ensures that a car is built step by step in the correct order, using a specific builder.
package main

import "fmt"

// Car represents the product being constructed.
// Product (Car): The complex object you're building. In this case, Car has various attributes like Engine, Color, Wheels, Seats, etc.

type Car struct {
	Engine       string
	Color        string
	Wheels       int
	Seats        int
	Transmission string
	Sunroof      bool
	GPS          bool
	Bluetooth    bool
}

// String method to display the Car details.
func (c *Car) String() string {
	return fmt.Sprintf("Car [Engine=%s, Color=%s, Wheels=%d, Seats=%d, Transmission=%s, Sunroof=%v, GPS=%v, Bluetooth=%v]",
		c.Engine, c.Color, c.Wheels, c.Seats, c.Transmission, c.Sunroof, c.GPS, c.Bluetooth)
}

// CarBuilder interface defines methods for constructing different parts of a Car.
type CarBuilder interface {
	BuildEngine() CarBuilder
	BuildColor() CarBuilder
	BuildWheels() CarBuilder
	BuildSeats() CarBuilder
	BuildTransmission() CarBuilder
	AddSunroof() CarBuilder
	AddGPS() CarBuilder
	AddBluetooth() CarBuilder
	GetResult() *Car
}

// SportsCarBuilder for building sports cars
// Concrete Builders: These classes (SportsCarBuilder, ElectricCarBuilder, FamilyCarBuilder) implement the CarBuilder interface 
// and provide specific implementations for constructing different types of cars.
type SportsCarBuilder struct {
	car *Car
}

func NewSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{
		car: &Car{},
	}
}

func (b *SportsCarBuilder) BuildEngine() CarBuilder {
	b.car.Engine = "V12 Engine"
	return b
}

func (b *SportsCarBuilder) BuildColor() CarBuilder {
	b.car.Color = "Red"
	return b
}

func (b *SportsCarBuilder) BuildWheels() CarBuilder {
	b.car.Wheels = 4
	return b
}

func (b *SportsCarBuilder) BuildSeats() CarBuilder {
	b.car.Seats = 2
	return b
}

func (b *SportsCarBuilder) BuildTransmission() CarBuilder {
	b.car.Transmission = "Manual"
	return b
}

func (b *SportsCarBuilder) AddSunroof() CarBuilder {
	b.car.Sunroof = true
	return b
}

func (b *SportsCarBuilder) AddGPS() CarBuilder {
	// Sports car doesn't need GPS
	return b
}

func (b *SportsCarBuilder) AddBluetooth() CarBuilder {
	// Sports car doesn't need Bluetooth
	return b
}

func (b *SportsCarBuilder) GetResult() *Car {
	return b.car
}

// ElectricCarBuilder for building electric cars
type ElectricCarBuilder struct {
	car *Car
}

func NewElectricCarBuilder() *ElectricCarBuilder {
	return &ElectricCarBuilder{
		car: &Car{},
	}
}

func (b *ElectricCarBuilder) BuildEngine() CarBuilder {
	b.car.Engine = "Electric Motor"
	return b
}

func (b *ElectricCarBuilder) BuildColor() CarBuilder {
	b.car.Color = "Blue"
	return b
}

func (b *ElectricCarBuilder) BuildWheels() CarBuilder {
	b.car.Wheels = 4
	return b
}

func (b *ElectricCarBuilder) BuildSeats() CarBuilder {
	b.car.Seats = 5
	return b
}

func (b *ElectricCarBuilder) BuildTransmission() CarBuilder {
	b.car.Transmission = "Automatic"
	return b
}

func (b *ElectricCarBuilder) AddSunroof() CarBuilder {
	// Electric car doesn't need Sunroof
	return b
}

func (b *ElectricCarBuilder) AddGPS() CarBuilder {
	b.car.GPS = true
	return b
}

func (b *ElectricCarBuilder) AddBluetooth() CarBuilder {
	// Electric car doesn't need Bluetooth
	return b
}

func (b *ElectricCarBuilder) GetResult() *Car {
	return b.car
}

// FamilyCarBuilder for building family cars
type FamilyCarBuilder struct {
	car *Car
}

func NewFamilyCarBuilder() *FamilyCarBuilder {
	return &FamilyCarBuilder{
		car: &Car{},
	}
}

func (b *FamilyCarBuilder) BuildEngine() CarBuilder {
	b.car.Engine = "V6 Engine"
	return b
}

func (b *FamilyCarBuilder) BuildColor() CarBuilder {
	b.car.Color = "White"
	return b
}

func (b *FamilyCarBuilder) BuildWheels() CarBuilder {
	b.car.Wheels = 4
	return b
}

func (b *FamilyCarBuilder) BuildSeats() CarBuilder {
	b.car.Seats = 7
	return b
}

func (b *FamilyCarBuilder) BuildTransmission() CarBuilder {
	b.car.Transmission = "Automatic"
	return b
}

func (b *FamilyCarBuilder) AddSunroof() CarBuilder {
	// Family car doesn't need Sunroof
	return b
}

func (b *FamilyCarBuilder) AddGPS() CarBuilder {
	// Family car doesn't need GPS
	return b
}

func (b *FamilyCarBuilder) AddBluetooth() CarBuilder {
	b.car.Bluetooth = true
	return b
}

func (b *FamilyCarBuilder) GetResult() *Car {
	return b.car
}

// Director will construct the Car using the CarBuilder.
// Director: The Director is responsible for managing the building process. 
// It ensures that a car is built step by step in the correct order, using a specific builder.
type Director struct {
	builder CarBuilder
}

// NewDirector creates a new Director with the given CarBuilder.
func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

// Construct builds the car step by step.
func (d *Director) Construct() *Car {
	return d.builder.BuildEngine().
		BuildColor().
		BuildWheels().
		BuildSeats().
		BuildTransmission().
		GetResult()
}

func main() {
	// Example 1: Building a Sports Car
	sportsCarBuilder := NewSportsCarBuilder()
	director := NewDirector(sportsCarBuilder)
	sportsCar := director.Construct()
	sportsCar = sportsCarBuilder.AddSunroof().GetResult() // Add sunroof during building

	fmt.Println("Sports Car:", sportsCar)

	// Example 2: Building an Electric Car
	electricCarBuilder := NewElectricCarBuilder()
	director = NewDirector(electricCarBuilder)
	electricCar := director.Construct()
	electricCar = electricCarBuilder.AddGPS().GetResult() // Add GPS during building

	fmt.Println("Electric Car:", electricCar)

	// Example 3: Building a Family Car
	familyCarBuilder := NewFamilyCarBuilder()
	director = NewDirector(familyCarBuilder)
	familyCar := director.Construct()
	familyCar = familyCarBuilder.AddBluetooth().GetResult() // Add Bluetooth during building

	fmt.Println("Family Car:", familyCar)
}
