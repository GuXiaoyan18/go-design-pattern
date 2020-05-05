package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Student struct {
	Name string
	Age uint32
}

// 这里是浅拷贝，值类型是完全拷贝一份，引用类型是拷贝其地址
func (s *Student) Clone() *Student{
	student := *s
	return &student
}

// 深拷贝
func (s *Student) DeepClone() *Student {
	student := new(Student)
	if err := deepClone(student, s); err != nil {
		log.Fatal(err.Error())
	}
	return student
}

func deepClone(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// 原型模式
func main() {
	a := func (a uint32) uint32 {
		a++
		return a
	}
	b := uint32(2)
	fmt.Println(a(b))
	fmt.Println(a(b))


	s1 := &Student{
		Name: "Jack",
		Age: 18,
	}

	s2 := s1.Clone()
	s3 := s1.DeepClone()

	if s1 == s2 || s1 == s3 {
		log.Println("copy error")
	}
}
