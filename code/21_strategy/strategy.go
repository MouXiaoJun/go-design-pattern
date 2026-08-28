package _21_strategy

import "fmt"

// Payment 支付策略接口，定义一族可替换的支付算法
type Payment interface {
	Pay(amount int) string
}

// Alipay 支付宝支付
type Alipay struct{}

func (Alipay) Pay(amount int) string {
	return fmt.Sprintf("使用支付宝支付 %d 元", amount)
}

// WechatPay 微信支付
type WechatPay struct{}

func (WechatPay) Pay(amount int) string {
	return fmt.Sprintf("使用微信支付 %d 元", amount)
}

// BankCard 银行卡支付
type BankCard struct{}

func (BankCard) Pay(amount int) string {
	return fmt.Sprintf("使用银行卡支付 %d 元", amount)
}

// Order 订单上下文，持有当前支付策略，客户端负责选择策略
type Order struct {
	payment Payment
}

func NewOrder(payment Payment) *Order {
	return &Order{payment: payment}
}

func (o *Order) SetPayment(payment Payment) {
	o.payment = payment
}

func (o *Order) Checkout(amount int) string {
	return o.payment.Pay(amount)
}
