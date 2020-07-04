package main

import "fmt"

// 策略模式，定义一系列的算法，把每一个算法都封装起来，使得它们可以互相替换

type Payment struct {
	context *PaymentContext
	strategy PaymentStrategy
}

func (p *Payment) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentContext struct {
	Name  string
	Money int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type CashPay struct {
}

func (c *CashPay) Pay(context *PaymentContext) {
	fmt.Printf("cash pay, name: %s, money: %d\n", context.Name, context.Money)
}

type AliPay struct {
}

func (a *AliPay) Pay(context *PaymentContext) {
	fmt.Printf("ali pay, name: %s, monney: %d\n", context.Name, context.Money)
}

func main() {
	context := &PaymentContext{Name: "jack", Money: 100}
	payment := &Payment{context: context}

	payment.SetStrategy(&CashPay{})
	payment.Pay()

	payment.SetStrategy(&AliPay{})
	payment.Pay()
}
