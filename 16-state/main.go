package main

import (
	"fmt"
	"time"
)

// 状态模式，可以解决项目中大量的if-else
// 状态模式适用于某一个对象的行为取决于对象的状态，并且对象的状态会在运行时转换，又或者有很多的if-else判断，
// 而这些判断只是因为状态不同而不断切换行为

type Game interface {
	Start(*Person)
}

type Person struct {
	Game
}

func (p *Person) SetGame(game Game) {
	p.Game = game
}

func (p *Person) Start() {
	p.Game.Start(p)
}

type Run struct {
}

func (r *Run) Start(p *Person) {
	fmt.Println("Running")
	time.Sleep(2*time.Second)
	p.SetGame(&Walk{})
}

type Swim struct {
}

func (s *Swim) Start(p *Person) {
	fmt.Println("Swimming")
	time.Sleep(2*time.Second)
	p.SetGame(&Walk{})
}

type Walk struct {
}

func (w *Walk) Start(p *Person) {
	fmt.Println("Walking")
}

func main() {
	p := &Person{
		Game: &Run{},
	}
	p.Start()
	p.SetGame(&Swim{})
	p.Start()
}
