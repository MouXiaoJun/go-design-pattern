package _11_flyweight

import "testing"

func TestGlyphReuse(t *testing.T) {
	f := NewGlyphFactory()
	Render(f, "hello hello", 0, 0)

	// 文本中不同字符只有 h/e/l/o/空格 共 5 个享元。
	if f.Len() != 5 {
		t.Fatalf("glyph count = %d, want 5", f.Len())
	}

	g1 := f.GetGlyph('l', "consolas", 12, "black")
	g2 := f.GetGlyph('l', "consolas", 12, "black")
	if g1 != g2 {
		t.Fatal("相同内在状态应共享同一个 Glyph 实例")
	}
}
