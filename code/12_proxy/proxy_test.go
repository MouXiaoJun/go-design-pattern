package _12_proxy

import (
	"testing"
	"time"
)

func TestCachedUserRepository(t *testing.T) {
	proxy := NewCachedUserRepository(&DBUserRepository{})

	start := time.Now()
	if _, err := proxy.FindByID(1); err != nil {
		t.Fatalf("first: %v", err)
	}
	first := time.Since(start)

	start = time.Now()
	if _, err := proxy.FindByID(1); err != nil {
		t.Fatalf("second: %v", err)
	}
	second := time.Since(start)

	if second >= first {
		t.Fatalf("缓存命中应明显快于回源：first=%v second=%v", first, second)
	}
}

func TestLazyImage(t *testing.T) {
	img := NewLazyImage("photo.png")
	if img.real != nil {
		t.Fatal("构造时不应立即加载")
	}
	if got := img.draw(); got == "" {
		t.Fatal("draw 结果为空")
	}
	if img.real == nil {
		t.Fatal("首次 draw 后应完成加载")
	}
}
