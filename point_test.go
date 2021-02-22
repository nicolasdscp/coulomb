package coulomb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(0, 0)

	assert.Equal(t, float64(0), p.X)
	assert.Equal(t, float64(0), p.Y)

	p = NewPoint(10, 10)

	assert.Equal(t, float64(10), p.X)
	assert.Equal(t, float64(10), p.Y)
}
