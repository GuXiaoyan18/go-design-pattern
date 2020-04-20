package main

import "fmt"

// 外观类
type Model struct {
	dao Dao // 子系统
	cao Cao // 子系统
}

func (m *Model) Test() {
	m.dao.Query()
	m.cao.Read()
}

func NewModel() *Model {
	return &Model{
		dao: &Mysql{},
		cao: &Redis{},
	}
}

type Dao interface {
	Query()
}

type Cao interface {
	Read()
}

type Mysql struct {
}

func (m *Mysql) Query() {
	fmt.Printf("Mysql Query")
}

type Redis struct {
}

func (r *Redis) Read() {
	fmt.Println("Redis Read")
}

// 外观模式提供一个高层次的接口，对客户端屏蔽子系统的调用过程，使客户端无需与子系统交互，代码变得简单
func main() {
	model := NewModel()
	model.Test()
}
