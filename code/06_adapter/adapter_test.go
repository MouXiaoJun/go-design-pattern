package _6_adapter

import "testing"

func TestGatewayAdapter(t *testing.T) {
	adapter := NewGatewayAdapter(LegacyGateway{})
	got, err := adapter.Pay(Order{Amount: 9.9, CardNo: "6222"})
	if err != nil {
		t.Fatal(err)
	}
	if got == "" {
		t.Fatal("Pay 返回空结果")
	}
}

func TestPayFunc(t *testing.T) {
	var svc PaymentService = PayFunc(func(order Order) (string, error) {
		return "ok", nil
	})
	if _, err := svc.Pay(Order{Amount: 1}); err != nil {
		t.Fatal(err)
	}
}
