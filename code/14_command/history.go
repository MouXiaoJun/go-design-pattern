package _14_command

// History 记录已执行的命令，提供撤销与重做。
type History struct {
	done   []Command
	undone []Command
}

func NewHistory() *History { return &History{} }

// Do 执行命令并压入历史栈，同时清空重做栈。
func (h *History) Do(c Command) {
	c.Execute()
	h.done = append(h.done, c)
	h.undone = nil
}

// Undo 撤销最近一次命令，成功返回 true。
func (h *History) Undo() bool {
	if len(h.done) == 0 {
		return false
	}
	c := h.done[len(h.done)-1]
	h.done = h.done[:len(h.done)-1]
	c.Undo()
	h.undone = append(h.undone, c)
	return true
}

// Redo 重做最近一次被撤销的命令，成功返回 true。
func (h *History) Redo() bool {
	if len(h.undone) == 0 {
		return false
	}
	c := h.undone[len(h.undone)-1]
	h.undone = h.undone[:len(h.undone)-1]
	c.Execute()
	h.done = append(h.done, c)
	return true
}
