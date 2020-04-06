package main

import "fmt"

type Car interface {
	Run()
}

type OdiCar struct {
}

func (o *OdiCar) Run() {
	fmt.Println("an Odi is running")
}

type BMWCar struct {
}

func (b *BMWCar) Run() {
	fmt.Println("an BMWC is running")
}

func NewCar(brand string) Car {
	if brand == "odi" {
		return &OdiCar{}
	} else if brand == "bmw" {
		return &BMWCar{}
	}
	return nil
}

func main() {
	car := NewCar("odi")
	car.Run()
	car = NewCar("bmw")
	car.Run()
}
