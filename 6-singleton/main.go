package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

// 单例模式
func main() {
	instance := GetInstance()
	fmt.Printf("ins: +v", instance)
}
