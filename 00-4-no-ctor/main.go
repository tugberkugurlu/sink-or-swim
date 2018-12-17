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

	// Frowned upon use, cannot do this anymore as the type is unexported
	// car2 := domain.car{}
	// fmt.Printf("%+v\n", car2)
}
