package main

import (
	"fmt"
)

// 抽象工厂模式，具有一个抽象工厂和多个抽象产品，有多少个抽象产品，抽象工厂中就需要有多少个对应的创建方法，对应有多少个具体工厂

// 抽象产品，汽车
type Car interface {
	Run()
}

// 抽象产品，电视
type TV interface {
	Show()
}

// 具体产品，odi汽车
type OdiCar struct {
}

func (o OdiCar) Run() {
	fmt.Println("an odi car is running")
}

// 具体产品，odi电视
type OdiTV struct {
}

// 具体产品，bmw汽车
type BMWCar struct {
}

func (b BMWCar) Run() {
	fmt.Println("an bmw car is running")
}

// 具体产品，bmw电视
type BMWTV struct {
}

func (b BMWTV) Show() {
	fmt.Println("an bmw tv is showing")
}

func (o OdiTV) Show() {
	fmt.Println("an odi TV is showing")
}

// 抽象工厂
type Factory interface {
	CreateCar() Car
	CreateTV() TV
}

// 具体工厂，odi
type OdiFactory struct {
}

func (o OdiFactory) CreateCar() Car {
	return OdiCar{}
}

func (o OdiFactory) CreateTV() TV {
	return OdiTV{}
}

// 具体工厂，bmw
type BMWFactory struct {
}

func (b BMWFactory) CreateCar() Car {
	return BMWCar{}
}

func (b BMWFactory) CreateTV() TV {
	return BMWTV{}
}

func main()  {
	var factory Factory
	factory = OdiFactory{}
	factory.CreateCar().Run()
	factory.CreateTV().Show()

	factory = BMWFactory{}
	factory.CreateCar().Run()
	factory.CreateTV().Show()
}
