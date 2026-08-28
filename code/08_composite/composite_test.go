package _8_composite

import (
	"strings"
	"testing"
)

func TestComposite(t *testing.T) {
	root := NewDirectory("root")
	home := NewDirectory("home")
	root.Add(NewFile("README.md", 100))
	root.Add(home)
	home.Add(NewFile("a.go", 50))
	home.Add(NewFile("b.go", 60))

	if got := root.Size(); got != 210 {
		t.Fatalf("size=%d want 210", got)
	}

	out := root.Print("")
	if !strings.Contains(out, "root/") || !strings.Contains(out, "a.go") {
		t.Fatalf("unexpected tree:\n%s", out)
	}
}
