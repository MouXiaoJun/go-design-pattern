package _23_visitor

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaVisitor(t *testing.T) {
	area := AreaVisitor{}
	circle := &Circle{Radius: 1}
	rect := &Rectangle{Width: 2, Height: 3}

	assert.InDelta(t, math.Pi, circle.Accept(area), 1e-9)
	assert.Equal(t, 6.0, rect.Accept(area))
}

func TestPerimeterVisitor(t *testing.T) {
	perimeter := PerimeterVisitor{}
	circle := &Circle{Radius: 1}
	rect := &Rectangle{Width: 2, Height: 3}

	assert.InDelta(t, 2*math.Pi, circle.Accept(perimeter), 1e-9)
	assert.Equal(t, 10.0, rect.Accept(perimeter))
}

func TestAreaBySwitch(t *testing.T) {
	assert.InDelta(t, math.Pi, AreaBySwitch(&Circle{Radius: 1}), 1e-9)
	assert.Equal(t, 6.0, AreaBySwitch(&Rectangle{Width: 2, Height: 3}))
}
