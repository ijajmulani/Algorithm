// Factory design pattern
// The Factory Pattern is a creational design pattern used to create objects without specifying the exact class of object that will be created.
// Instead of using a direct constructor call, the pattern provides a method in a factory class that returns an instance of an object.

package main

import "fmt"

const (
	SedanType     = 0
	HatchBackType = 1
)

func main() {
	c := GetCarFactory(SedanType)
	c.GetCar()

	s := GetCarFactory(HatchBackType)
	s.GetCar()
}

type Car interface {
	GetCar()
}

type SedanCar struct {
	Name string
}

func (s SedanCar) GetCar() {
	fmt.Println(s.Name)
}

type HatchBackCar struct {
	Name string
}

func (s HatchBackCar) GetCar() {
	fmt.Println(s.Name)
}

func GetCarFactory(t int) Car {
	switch t {
	case HatchBackType:
		return HatchBackCar{Name: "hatchback"}
	case SedanType:
		return SedanCar{Name: "sedan"}
	default:
		return nil
	}
}
