package main

import "fmt"

type Payer interface {
	Pay(amount int)
}

type Alipay struct {
	amount int
}

func (a *Alipay) Pay(amount int) {
	fmt.Printf("使用支付宝支付了%d元", amount)
}

type Applepay struct {
	amount int
}

func (a *Applepay) Pay(amount int) {
	fmt.Printf("使用ApplePay支付了%d元", amount)
}

func main() {
	customers := []Payer{
		&Alipay{},
		&Applepay{},
	}
	for i := range customers {
		customers[i].Pay(100)
	}
}
