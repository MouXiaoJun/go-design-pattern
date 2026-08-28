package _19_observer

import "testing"

func TestStockNotify(t *testing.T) {
	stock := NewStock("AAPL")
	ch := stock.Subscribe("u1")

	stock.Update(100.0)

	select {
	case price := <-ch:
		if price != 100.0 {
			t.Fatalf("期望价格 100, 实际 %v", price)
		}
	default:
		t.Fatal("订阅者应收到价格更新")
	}
}

func TestStockUnsubscribe(t *testing.T) {
	stock := NewStock("AAPL")
	ch := stock.Subscribe("u1")
	stock.Unsubscribe("u1")

	stock.Update(200.0)

	if _, ok := <-ch; ok {
		t.Fatal("取消订阅后通道应已关闭")
	}
}
