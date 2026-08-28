package _14_command

// Editor 文本编辑器，保存当前文本，是命令的接收者。
type Editor struct {
	text string
}

func (e *Editor) Text() string { return e.text }

// insert 在 pos 处插入 s，越界则忽略。
func (e *Editor) insert(pos int, s string) {
	if pos < 0 || pos > len(e.text) {
		return
	}
	e.text = e.text[:pos] + s + e.text[pos:]
}

// delete 删除从 pos 开始 n 个字符，返回被删除的内容。
func (e *Editor) delete(pos, n int) string {
	if pos < 0 || pos >= len(e.text) {
		return ""
	}
	end := pos + n
	if end > len(e.text) {
		end = len(e.text)
	}
	removed := e.text[pos:end]
	e.text = e.text[:pos] + e.text[end:]
	return removed
}
