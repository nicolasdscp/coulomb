package coulomb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewForce(t *testing.T) {
	f := NewForce(0, 0)

	assert.Equal(t, float64(0), f.X)
	assert.Equal(t, float64(0), f.Y)

	f = NewForce(10, 10)

	assert.Equal(t, float64(10), f.X)
	assert.Equal(t, float64(10), f.Y)
}
