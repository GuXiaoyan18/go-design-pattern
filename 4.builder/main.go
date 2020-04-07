package main

import "github.com/gin-gonic/gin/ginS"

// 建造者模式，将一个复杂对象的表示与构建分离，使得同样的构建过程可以构建不同的表示。
// 建造者模式的四个部分，Builder：建造者，ConcreteBuilder：具体建造者，Director：指挥者，产品：Product

// 产品
type Character struct {
	Name string
	Level uint32
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) SetLevel(level uint32) {
	c.Level = level
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) GetLevel() uint32 {
	return c.Level
}

// 抽象建造者
type Builder interface {
	SetName(string) Builder
	SetLevel(uint32) Builder
	Build() *Character
}

// 具体建造者
type CharacterBuilder struct {
	character *Character
}

func (c *CharacterBuilder) SetName(name string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetName(name)
	return c
}

func (c *CharacterBuilder) SetLevel(level uint32) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetLevel(level)
	return c
}

func (c *CharacterBuilder) Build() *Character {
	return c.character
}

// 导演者，作用是规范生产流程
type Director struct {
	builder Builder
}

func main() {
	df := &Director{&CharacterBuilder{}}
	df.builder.SetName("jack").SetLevel(30).Build()
}

