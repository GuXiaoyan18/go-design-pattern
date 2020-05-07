package main

import "fmt"

// 命令接口
type Command interface {
	Execute()
}

// 具体命令
type OnCommand struct {
	light *Light
}

func (o *OnCommand) Execute() {
	o.light.On()
}

// 具体命令
type OffCommand struct {
	light *Light
}

func (o *OffCommand) Execute() {
	o.light.Off()
}

// reciever，接收器，命令的真正实现者
type Light struct {
}

func (l *Light) On() {
	fmt.Println("light on")
}

func (l *Light) Off() {
	fmt.Println("light off")
}

// 请求者，持有命令对象
type Invoker struct {
	command1 Command
	command2 Command
}

func (i *Invoker) Call() {
	i.command1.Execute()
	i.command2.Execute()
}

func NewInvoker(command1, command2 Command) *Invoker {
	return &Invoker{
		command1: command1,
		command2: command2,
	}
}

// 命令模式，调用顺序：调用命令者->接受者->命令，实际上是将请求者和实现者解耦
func main() {
	light := &Light{}

	onCommand := &OnCommand{light: light}
	offCommand := &OffCommand{light: light}

	invoker := NewInvoker(onCommand, offCommand)
	invoker.Call()
}
