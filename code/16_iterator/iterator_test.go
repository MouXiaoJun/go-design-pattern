package _16_iterator

import "testing"

func TestSliceIterator(t *testing.T) {
	c := NewCollection("a", "b", "c")

	var got []string
	for it := c.Iterator(); it.Next(); {
		got = append(got, it.Value())
	}

	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestChannelIter(t *testing.T) {
	c := NewCollection("x", "y")

	var got []string
	for s := range c.Iter() {
		got = append(got, s)
	}

	if len(got) != 2 || got[0] != "x" || got[1] != "y" {
		t.Fatalf("got %v, want [x y]", got)
	}
}
