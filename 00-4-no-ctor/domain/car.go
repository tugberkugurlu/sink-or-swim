package domain

import "errors"

type car struct {
	make  string
	model string
	price float64
}

func NewCar(make, model string, price float64) (car, error) {
	if make == "" {
		return car{}, errors.New("'make' cannot be empty")
	}

	if model == "" {
		return car{}, errors.New("'model' cannot be empty")
	}

	if price < 0 {
		return car{}, errors.New("'price' cannot be lower than '0'")
	}

	return car{
		make: make,
		model: model,
		price: price,
	}, nil
}

func (car *car) Price() float64 {
	return car.price
}

func (car *car) Make() string {
	return car.make
}

func (car *car) Model() string {
	return car.model
}

