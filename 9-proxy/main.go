package main

type Subject interface {
	Do() string
}

type RealSubject struct {

}

func (r RealSubject) Do() string {
	return "Realsubject Do"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	res := "pre"
	res += p.real.Do()
	res += "after"
	return res
}

// 代理模式
func main() {
	var sub Subject
	sub = Proxy{}
	sub.Do()
}
