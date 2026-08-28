package _15_interpreter

// Expression 表达式接口：每个语法节点都能求值。
type Expression interface {
	Interpret() int
}

// Number 数字节点，终结符。
type Number struct {
	value int
}

func (n *Number) Interpret() int { return n.value }

// Add 加法节点，非终结符。
type Add struct {
	left, right Expression
}

func (a *Add) Interpret() int { return a.left.Interpret() + a.right.Interpret() }

// Sub 减法节点。
type Sub struct {
	left, right Expression
}

func (s *Sub) Interpret() int { return s.left.Interpret() - s.right.Interpret() }

// Mul 乘法节点。
type Mul struct {
	left, right Expression
}

func (m *Mul) Interpret() int { return m.left.Interpret() * m.right.Interpret() }

// Div 除法节点。
type Div struct {
	left, right Expression
}

func (d *Div) Interpret() int {
	right := d.right.Interpret()
	if right == 0 {
		// 与 Go 原生行为一致：整数除零会 panic，这里给出明确的错误信息。
		// 真实项目里应返回 error（把 Expression 接口改成 Interpret() (int, error)）。
		panic("division by zero")
	}
	return d.left.Interpret() / right
}
