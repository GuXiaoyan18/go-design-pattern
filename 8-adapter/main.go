package main

import "fmt"

// 适配器接口
type Adapter interface {
	Request() string
}

// 被适配的接口
type Target interface {
	SpecificRequest() string
}

// 工厂函数创建被适配接口的实例
func NewTarget() Target {
	return &TargetIpl{}
}

func NewAdapter(target Target) Adapter {
	return &AdapterIpl{Target: target}
}

type TargetIpl struct {
}

func (t *TargetIpl) SpecificRequest() string {
	msg := "a targetIpl is running"
	fmt.Println(msg)
	return msg
}

type AdapterIpl struct {
	Target
}

func (a *AdapterIpl) Request() string {
	return a.SpecificRequest()
}

// 适配器模式
func main() {
	target := NewTarget()
	adapter := NewAdapter(target)
	adapter.Request()
}
