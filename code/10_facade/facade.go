package _10_facade

import (
	"errors"
	"fmt"
)

// OrderService 是门面：客户端只依赖它下单，无需了解库存、支付、通知子系统。
type OrderService struct {
	stock    *StockService
	pay      *PaymentService
	notifier *NotifyService
}

func NewOrderService() *OrderService {
	return &OrderService{
		stock:    newStockService(),
		pay:      &PaymentService{},
		notifier: &NotifyService{},
	}
}

// PlaceOrder 一键下单：校验库存 → 扣款 → 扣减库存 → 通知用户。
func (o *OrderService) PlaceOrder(userID, goodsID string, amount int) error {
	if err := o.stock.check(goodsID); err != nil {
		return err
	}
	if err := o.pay.charge(userID, amount); err != nil {
		return err
	}
	o.stock.deduct(goodsID)
	o.notifier.send(userID, fmt.Sprintf("订单创建成功：%s", goodsID))
	return nil
}

// StockService 库存子系统。
type StockService struct {
	inventory map[string]int
}

func newStockService() *StockService {
	return &StockService{inventory: map[string]int{"phone": 10}}
}

func (s *StockService) check(goodsID string) error {
	if s.inventory[goodsID] <= 0 {
		return errors.New("库存不足")
	}
	return nil
}

func (s *StockService) deduct(goodsID string) {
	s.inventory[goodsID]--
}

// PaymentService 支付子系统。
type PaymentService struct{}

func (p *PaymentService) charge(userID string, amount int) error {
	if amount <= 0 {
		return errors.New("金额非法")
	}
	return nil
}

// NotifyService 通知子系统，这里仅收集消息用于演示。
type NotifyService struct {
	sent []string
}

func (n *NotifyService) send(userID, msg string) {
	n.sent = append(n.sent, userID+": "+msg)
}
