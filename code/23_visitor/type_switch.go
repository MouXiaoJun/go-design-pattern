package _23_visitor

import "math"

// AreaBySwitch 用类型分派直接求面积。
// 类型集小且稳定时，这种方式比访问者更简单直接。
func AreaBySwitch(s Shape) float64 {
	switch v := s.(type) {
	case *Circle:
		return math.Pi * v.Radius * v.Radius
	case *Rectangle:
		return v.Width * v.Height
	}
	return 0
}
