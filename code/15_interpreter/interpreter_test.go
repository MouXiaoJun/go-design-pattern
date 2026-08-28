package _15_interpreter

import "testing"

func TestEval(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"1+2*3", 7},
		{"(1+2)*3", 9},
		{"10-2-3", 5},
		{"8/4*3", 6},
	}
	for _, c := range cases {
		got, err := Eval(c.input)
		if err != nil {
			t.Fatalf("Eval(%q) error: %v", c.input, err)
		}
		if got != c.want {
			t.Fatalf("Eval(%q) = %d, want %d", c.input, got, c.want)
		}
	}
}

func TestEvalInvalid(t *testing.T) {
	if _, err := Eval("1+"); err == nil {
		t.Fatal("want error, got nil")
	}
	if _, err := Eval("(1+2"); err == nil {
		t.Fatal("want error, got nil")
	}
}
