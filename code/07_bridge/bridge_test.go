package _7_bridge

import "testing"

func TestBridge(t *testing.T) {
	circle := NewCircle(VectorRenderer{}, 1.5)
	if got := circle.Draw(); got != "矢量圆 r=1.5" {
		t.Fatalf("unexpected: %s", got)
	}

	square := NewSquare(RasterRenderer{}, 2.0)
	if got := square.Draw(); got != "栅格方 a=2.0" {
		t.Fatalf("unexpected: %s", got)
	}
}
