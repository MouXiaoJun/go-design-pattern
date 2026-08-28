package _10_facade

import "testing"

func TestPlaceOrderSuccess(t *testing.T) {
	o := NewOrderService()
	if err := o.PlaceOrder("u1", "phone", 100); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.stock.inventory["phone"] != 9 {
		t.Fatalf("inventory = %d, want 9", o.stock.inventory["phone"])
	}
	if len(o.notifier.sent) != 1 {
		t.Fatalf("sent count = %d, want 1", len(o.notifier.sent))
	}
}

func TestPlaceOrderOutOfStock(t *testing.T) {
	o := NewOrderService()
	if err := o.PlaceOrder("u1", "unknown", 100); err == nil {
		t.Fatal("want error, got nil")
	}
}
