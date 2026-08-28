package _18_memento

// Memento 快照：保存发起者在某个时刻的内部状态，对外只读。
type Memento struct {
	content string
}

// Content 返回快照保存的内容。
func (m *Memento) Content() string { return m.content }

// Editor 发起者：持有可被保存和恢复的状态。
type Editor struct {
	content string
}

// Type 追加输入，模拟编辑器编辑动作。
func (e *Editor) Type(text string) {
	e.content += text
}

func (e *Editor) Content() string { return e.content }

// Save 把当前状态打包成一个快照。
func (e *Editor) Save() *Memento {
	return &Memento{content: e.content}
}

// Restore 用快照恢复状态。
func (e *Editor) Restore(m *Memento) {
	if m == nil {
		return
	}
	e.content = m.content
}

// History 负责人：只负责保存与取出快照，不关心快照内部结构。
type History struct {
	mementos []*Memento
}

func (h *History) Push(m *Memento) {
	h.mementos = append(h.mementos, m)
}

// Pop 弹出最近一次快照，空栈时返回 nil。
func (h *History) Pop() *Memento {
	if len(h.mementos) == 0 {
		return nil
	}
	last := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return last
}
