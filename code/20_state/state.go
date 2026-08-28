package _20_state

import "fmt"

// State 订单状态接口：每个状态对象封装自己处理动作的行为与转移。
type State interface {
	Name() string
	Pay(order *Order) error
	Ship(order *Order) error
	Complete(order *Order) error
}

// Order 上下文：持有当前状态，把动作委托给状态对象处理。
type Order struct {
	state State
}

func NewOrder() *Order {
	return &Order{state: &PendingPayment{}}
}

func (o *Order) State() State { return o.state }

func (o *Order) Pay() error      { return o.state.Pay(o) }
func (o *Order) Ship() error     { return o.state.Ship(o) }
func (o *Order) Complete() error { return o.state.Complete(o) }

// PendingPayment 待支付：只能支付。
type PendingPayment struct{}

func (s *PendingPayment) Name() string { return "待支付" }

func (s *PendingPayment) Pay(o *Order) error {
	o.state = &Paid{}
	return nil
}

func (s *PendingPayment) Ship(o *Order) error {
	return fmt.Errorf("订单未支付，不能发货")
}

func (s *PendingPayment) Complete(o *Order) error {
	return fmt.Errorf("订单未支付，不能完成")
}

// Paid 已支付：可以发货。
type Paid struct{}

func (s *Paid) Name() string { return "已支付" }

func (s *Paid) Pay(o *Order) error {
	return fmt.Errorf("订单已支付，请勿重复支付")
}

func (s *Paid) Ship(o *Order) error {
	o.state = &Shipped{}
	return nil
}

func (s *Paid) Complete(o *Order) error {
	return fmt.Errorf("订单未发货，不能完成")
}

// Shipped 已发货：可以完成。
type Shipped struct{}

func (s *Shipped) Name() string { return "已发货" }

func (s *Shipped) Pay(o *Order) error {
	return fmt.Errorf("订单已发货，不能支付")
}

func (s *Shipped) Ship(o *Order) error {
	return fmt.Errorf("订单已发货，请勿重复发货")
}

func (s *Shipped) Complete(o *Order) error {
	o.state = &Completed{}
	return nil
}

// Completed 已完成：终态，不能再有任何动作。
type Completed struct{}

func (s *Completed) Name() string { return "已完成" }

func (s *Completed) Pay(o *Order) error {
	return fmt.Errorf("订单已完成，不能支付")
}

func (s *Completed) Ship(o *Order) error {
	return fmt.Errorf("订单已完成，不能发货")
}

func (s *Completed) Complete(o *Order) error {
	return fmt.Errorf("订单已完成，请勿重复操作")
}
