package main

import "fmt"

// 聚集抽象接口
type Aggregate interface {
	CreateIterator() Iterator
}

// 抽象迭代器
type Iterator interface {
	First()
	Next() interface{}
	IsDone() bool
}

// 聚集对象
type Numbers struct {
	data []int
}

func (n *Numbers) CreateIterator() Iterator {
	return &NumberIterator{
		numbers: n,
		index:   0,
	}
}

// 具体迭代器
type NumberIterator struct {
	numbers *Numbers
	index   int
}

func (ni *NumberIterator) First() {
	ni.index = 0
}

func (ni *NumberIterator) Next() interface{} {
	if !ni.IsDone() {
		cur := ni.numbers.data[ni.index]
		ni.index++
		return cur
	}
	return nil
}

func (ni *NumberIterator) IsDone() bool {
	if len(ni.numbers.data) == 0 {
		return false
	}
	return ni.index > len(ni.numbers.data)-1
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Printf("%v\n", c)
	}
}

// 迭代器模型，核心是通过聚集对象自身的方法创建一个迭代器，该迭代器包含了对聚集对象的引用
func main() {
	numbers := &Numbers{
		data: []int{1, 2, 3, 4, 5},
	}
	iterator := numbers.CreateIterator()
	IteratorPrint(iterator)
}
