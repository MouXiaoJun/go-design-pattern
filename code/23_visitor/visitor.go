package _23_visitor

import "math"

// Visitor 访问者接口：为每种元素类型声明一个访问方法
type Visitor interface {
	VisitCircle(*Circle) float64
	VisitRectangle(*Rectangle) float64
}

// Shape 元素接口：接受访问者，实现双分派
type Shape interface {
	Accept(Visitor) float64
}

// Circle 圆形
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) float64 {
	return v.VisitCircle(c)
}

// Rectangle 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(v Visitor) float64 {
	return v.VisitRectangle(r)
}

// AreaVisitor 面积访问者：给元素结构添加「求面积」操作
type AreaVisitor struct{}

func (AreaVisitor) VisitCircle(c *Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}

func (AreaVisitor) VisitRectangle(r *Rectangle) float64 {
	return r.Width * r.Height
}

// PerimeterVisitor 周长访问者：给元素结构添加「求周长」操作
type PerimeterVisitor struct{}

func (PerimeterVisitor) VisitCircle(c *Circle) float64 {
	return 2 * math.Pi * c.Radius
}

func (PerimeterVisitor) VisitRectangle(r *Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
