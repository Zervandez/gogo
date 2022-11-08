package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")

	Ariel := Person{true, "Ariel", 21, "Student", "Albaner", "Jewish"}

	var ptr *bool = &Ariel.Gender

	fmt.Println(ptr)
	fmt.Println(Ariel.Name)

}

type Person struct {
	Gender      bool
	Name        string
	age         int
	occupation  string
	nationality string
	race        string
}

type Vehicle struct {
	Make     string
	Model    string
	Year     int
	FuelType string
	price    float32
	maxSpeed int
}
