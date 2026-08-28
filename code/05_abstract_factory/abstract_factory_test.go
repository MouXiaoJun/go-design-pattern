package _5_abstract_factory

import (
	"strings"
	"testing"
)

func TestRenderUI(t *testing.T) {
	dark := RenderUI(DarkThemeFactory{})
	light := RenderUI(LightThemeFactory{})

	if dark == light {
		t.Fatal("暗色与浅色主题应渲染出不同结果")
	}
	if !strings.Contains(dark, "#222") || !strings.Contains(light, "#fff") {
		t.Fatalf("主题渲染结果不符合预期: dark=%s light=%s", dark, light)
	}
}
