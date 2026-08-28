package _11_flyweight

import (
	"fmt"
	"sync"
)

// Glyph 是享元对象：保存字符的内在状态（字符本身、字体、字号、颜色），不可变。
type Glyph struct {
	char  rune
	font  string
	size  int
	color string
}

// draw 接收外在状态 x,y 作为参数，因此同一个 Glyph 可被多处复用。
func (g *Glyph) draw(x, y int) string {
	return fmt.Sprintf("draw %q at (%d,%d) font=%s size=%d color=%s", g.char, x, y, g.font, g.size, g.color)
}

// GlyphFactory 享元工厂：按内在状态缓存 Glyph，相同字符共享同一实例。
type GlyphFactory struct {
	mu     sync.Mutex
	glyphs map[string]*Glyph
}

func NewGlyphFactory() *GlyphFactory {
	return &GlyphFactory{glyphs: make(map[string]*Glyph)}
}

func (f *GlyphFactory) GetGlyph(char rune, font string, size int, color string) *Glyph {
	key := fmt.Sprintf("%c|%s|%d|%s", char, font, size, color)
	f.mu.Lock()
	defer f.mu.Unlock()
	if g, ok := f.glyphs[key]; ok {
		return g
	}
	g := &Glyph{char: char, font: font, size: size, color: color}
	f.glyphs[key] = g
	return g
}

func (f *GlyphFactory) Len() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.glyphs)
}

// Render 渲染文本：每个字符从工厂取享元，位置作为外在状态逐个传入。
func Render(f *GlyphFactory, text string, x, y int) []string {
	var out []string
	for _, c := range text {
		g := f.GetGlyph(c, "consolas", 12, "black")
		out = append(out, g.draw(x, y))
		x += g.size
	}
	return out
}
