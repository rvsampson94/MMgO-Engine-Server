package engine

// Vector is a 2D vector
type Vector struct {
	X float64
	Y float64
}

// NewVector creates a new 2D vector
func NewVector(x float64, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}
