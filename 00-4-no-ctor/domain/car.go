package domain

import "errors"

type Car struct {
	make  string
	model string
	price float64
}

func NewCar(make, model string, price float64) (Car, error) {
	if make == "" {
		return Car{}, errors.New("'make' cannot be empty")
	}

	if model == "" {
		return Car{}, errors.New("'model' cannot be empty")
	}

	if price < 0 {
		return Car{}, errors.New("'price' cannot be lower than '0'")
	}

	return Car{
		make: make,
		model: model,
		price: price,
	}, nil
}

func (car *Car) Price() float64 {
	return car.price
}

func (car *Car) Make() string {
	return car.make
}

func (car *Car) Model() string {
	return car.model
}

