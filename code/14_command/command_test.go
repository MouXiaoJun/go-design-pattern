package _14_command

import "testing"

func TestEditorUndoRedo(t *testing.T) {
	e := &Editor{}
	h := NewHistory()

	h.Do(&InsertCommand{editor: e, pos: 0, text: "hello"})
	h.Do(&InsertCommand{editor: e, pos: 5, text: " world"})
	if e.Text() != "hello world" {
		t.Fatalf("want %q, got %q", "hello world", e.Text())
	}

	h.Do(&DeleteCommand{editor: e, pos: 5, n: 6})
	if e.Text() != "hello" {
		t.Fatalf("want %q, got %q", "hello", e.Text())
	}

	h.Undo()
	if e.Text() != "hello world" {
		t.Fatalf("after undo want %q, got %q", "hello world", e.Text())
	}

	h.Redo()
	if e.Text() != "hello" {
		t.Fatalf("after redo want %q, got %q", "hello", e.Text())
	}
}
