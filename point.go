package coulomb

// Point define a point in 2D space
type Point struct {
	X float64 `json:"x"` // X coordinate in meter
	Y float64 `json:"y"` // Y coordinate in meter
}

// NewPoint create and return a new instance of Point struct.
func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}
