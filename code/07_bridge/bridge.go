package _7_bridge

import "fmt"

// Renderer 渲染实现维度：与形状抽象分离，二者可独立扩展
type Renderer interface {
	RenderCircle(radius float64) string
	RenderSquare(side float64) string
}

// VectorRenderer 矢量渲染实现
type VectorRenderer struct{}

func (VectorRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("矢量圆 r=%.1f", radius)
}

func (VectorRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("矢量方 a=%.1f", side)
}

// RasterRenderer 栅格渲染实现
type RasterRenderer struct{}

func (RasterRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("栅格圆 r=%.1f", radius)
}

func (RasterRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("栅格方 a=%.1f", side)
}

// Shape 形状抽象维度：通过组合持有渲染器，二者桥接
type Shape struct {
	renderer Renderer
}

// Circle 圆
type Circle struct {
	Shape
	radius float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{Shape: Shape{renderer: renderer}, radius: radius}
}

func (c *Circle) Draw() string {
	return c.renderer.RenderCircle(c.radius)
}

// Square 正方形
type Square struct {
	Shape
	side float64
}

func NewSquare(renderer Renderer, side float64) *Square {
	return &Square{Shape: Shape{renderer: renderer}, side: side}
}

func (s *Square) Draw() string {
	return s.renderer.RenderSquare(s.side)
}
