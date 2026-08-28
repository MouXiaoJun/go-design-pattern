package _20_state

import "testing"

func TestOrderHappyPath(t *testing.T) {
	order := NewOrder()

	if order.State().Name() != "待支付" {
		t.Fatalf("初始状态应为待支付, 实际 %s", order.State().Name())
	}

	if err := order.Ship(); err == nil {
		t.Fatal("未支付订单发货应报错")
	}

	if err := order.Pay(); err != nil {
		t.Fatalf("支付应成功: %v", err)
	}
	if order.State().Name() != "已支付" {
		t.Fatalf("支付后状态应为已支付, 实际 %s", order.State().Name())
	}

	if err := order.Ship(); err != nil {
		t.Fatalf("发货应成功: %v", err)
	}
	if err := order.Complete(); err != nil {
		t.Fatalf("完成应成功: %v", err)
	}
	if order.State().Name() != "已完成" {
		t.Fatalf("完成后状态应为已完成, 实际 %s", order.State().Name())
	}

	if err := order.Pay(); err == nil {
		t.Fatal("已完成订单支付应报错")
	}
}
