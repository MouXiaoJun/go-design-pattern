package _21_strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentStrategy(t *testing.T) {
	order := NewOrder(Alipay{})
	assert.Equal(t, "使用支付宝支付 100 元", order.Checkout(100))

	order.SetPayment(WechatPay{})
	assert.Equal(t, "使用微信支付 50 元", order.Checkout(50))

	order.SetPayment(BankCard{})
	assert.Equal(t, "使用银行卡支付 1 元", order.Checkout(1))
}

func TestSortStrategyFunc(t *testing.T) {
	asc := NewSorter(Asc)
	data := []int{3, 1, 2}
	asc.Sort(data)
	assert.Equal(t, []int{1, 2, 3}, data)

	desc := NewSorter(Desc)
	data = []int{3, 1, 2}
	desc.Sort(data)
	assert.Equal(t, []int{3, 2, 1}, data)
}
