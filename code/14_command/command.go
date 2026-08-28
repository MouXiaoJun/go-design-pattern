package _14_command

// Command 命令接口：可执行、可撤销。
type Command interface {
	Execute()
	Undo()
}

// InsertCommand 在指定位置插入一段文本。
type InsertCommand struct {
	editor *Editor
	pos    int
	text   string
}

func (c *InsertCommand) Execute() { c.editor.insert(c.pos, c.text) }

func (c *InsertCommand) Undo() { c.editor.delete(c.pos, len(c.text)) }

// DeleteCommand 删除指定位置的一段文本，执行时保存被删内容以便撤销。
type DeleteCommand struct {
	editor  *Editor
	pos     int
	n       int
	removed string
}

func (c *DeleteCommand) Execute() { c.removed = c.editor.delete(c.pos, c.n) }

func (c *DeleteCommand) Undo() { c.editor.insert(c.pos, c.removed) }
