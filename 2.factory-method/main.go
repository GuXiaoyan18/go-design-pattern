package main

import "fmt"

// 工厂方法模式，只有一个抽象工厂，一个抽象产品

type Car interface {
	Run()
}

type CarFactory interface {
	Create() Car
}

type OdiCarFactory struct {
}

func (o OdiCarFactory) Create() Car {
	return &OdiCar{}
}

type BMWCarFactory struct {
}

func (o BMWCarFactory) Create() Car {
	return &BMWCar{}
}

type OdiCar struct {
}

func (o *OdiCar) Run() {
	fmt.Println("an odi is running")
}

type BMWCar struct {
}

func (b *BMWCar) Run() {
	fmt.Println("an bmw is running")
}

func main() {
	var carFactory CarFactory
	carFactory = OdiCarFactory{}
	carFactory.Create().Run()

	carFactory = BMWCarFactory{}
	carFactory.Create().Run()
}
