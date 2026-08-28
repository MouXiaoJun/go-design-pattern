package _18_memento

import "testing"

func TestEditorUndo(t *testing.T) {
	editor := &Editor{}
	history := &History{}

	editor.Type("hello")
	history.Push(editor.Save())

	editor.Type(" world")
	if editor.Content() != "hello world" {
		t.Fatalf("期望 hello world, 实际 %q", editor.Content())
	}

	editor.Restore(history.Pop())
	if editor.Content() != "hello" {
		t.Fatalf("撤销后期望 hello, 实际 %q", editor.Content())
	}
}

func TestHistoryPopEmpty(t *testing.T) {
	history := &History{}
	if m := history.Pop(); m != nil {
		t.Fatalf("空栈应返回 nil, 实际 %v", m)
	}
}
