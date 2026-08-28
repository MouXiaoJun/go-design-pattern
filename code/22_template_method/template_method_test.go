package _22_template_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperPipeline(t *testing.T) {
	p := NewUpperPipeline()
	assert.Equal(t, "ABC\n", p.Run("abc\n"))
}

func TestTrimPipeline(t *testing.T) {
	p := NewTrimPipeline()
	assert.Equal(t, "hi\n", p.Run("  hi  \n"))
}
