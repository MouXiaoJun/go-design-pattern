package _6_adapter

import (
	"errors"
	"fmt"
)

// Order 新系统下单结构
type Order struct {
	Amount float64
	CardNo string
}

// PaymentService 新系统期望的统一支付接口
type PaymentService interface {
	Pay(order Order) (string, error)
}

// LegacyGateway 老系统支付网关，接口与新系统不兼容
type LegacyGateway struct{}

// Charge 老接口：按金额和卡号扣款
func (LegacyGateway) Charge(amount float64, cardNo string) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be positive")
	}
	return fmt.Sprintf("charge %.2f via card %s", amount, cardNo), nil
}

// GatewayAdapter 对象适配器：包装老网关，实现新接口
type GatewayAdapter struct {
	gateway LegacyGateway
}

func NewGatewayAdapter(gateway LegacyGateway) *GatewayAdapter {
	return &GatewayAdapter{gateway: gateway}
}

func (a *GatewayAdapter) Pay(order Order) (string, error) {
	return a.gateway.Charge(order.Amount, order.CardNo)
}

// PayFunc 函数适配器：让普通函数实现 PaymentService 接口（类比 http.HandlerFunc）
type PayFunc func(order Order) (string, error)

func (f PayFunc) Pay(order Order) (string, error) {
	return f(order)
}
