package main

import "fmt"

type PaymentStratege interface {
	Pay(amount float64) string
}

type WechatPay struct{}

func (w WechatPay) Pay(amount float64) string {
	return fmt.Sprintf("微信支付了%.2f元", amount)
}

type AliPay struct{}

func (a AliPay) Pay(amount float64) string {
	return fmt.Sprintf("支付宝支付了%.2f元", amount)
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("信用卡支付了%.2f元", amount)
}

type Cash struct{}

func (c Cash) Pay(amount float64) string {
	return fmt.Sprintf("使用现金支付了%.2f元", amount)
}

func main() {
	PaymentMethods := map[string]PaymentStratege{
		"wechat": WechatPay{},
		"alipay": AliPay{},
		"card":   CreditCard{},
		"cash":   Cash{},
	}

	userChoice := "cash"
	amount := 69.90

	method, exists := PaymentMethods[userChoice]

	if exists {
		result := method.Pay(amount)
		fmt.Println("支付成功", result)
	} else {
		fmt.Println("不支持的支付方式")
	}
}
