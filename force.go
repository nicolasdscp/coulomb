package coulomb

// Force define a 2D force applied in X and Y.
type Force struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewForce create and return a new instance of Force struct.
func NewForce(x, y float64) Force {
	return Force{X: x, Y: y}
}
