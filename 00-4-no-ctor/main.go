package main

import (
	"fmt"
	"./domain"
)

func main() {
	// Legit use
	car1, err := domain.NewCar("Renault", "Clio", 12.750)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", car1)

	// Frowned upon use
	car2 := domain.Car{}
	fmt.Printf("%+v\n", car2)
}
