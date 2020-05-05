package main

import "fmt"

// 主题
type Subject struct {
	observers []Observer
	status    string
}

// 主题初始化
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// 添加观察者对象
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// 通知观察者
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// 更改状态
func (s *Subject) UpdateStatus(status string) {
	s.status = status
	s.notify()
}

// 观察者接口，需要实现Update方法
type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader{
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("Reader %s recieve notify from %s", r.name, s.status)
}

// 观察者模式，当一个对象修改时，可以通知到它的依赖对象
func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.UpdateStatus("begin")
}
