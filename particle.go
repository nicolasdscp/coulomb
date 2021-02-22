package coulomb

// Particle define a single charge in space. The charge have a charge value in coulomb,
// mass, position, velocity and acceleration.
type Particle struct {
	Q   float64 `json:"q"`		// Charge of the particle in Coulomb
	M   float64 `json:"m"`		// Mass of the particle in Kg
	Pos Point   `json:"pos"`	// Position of the particle
	Vel Force   `json:"vel"`	// Velocity of the particle in m.s-1
	Acc Force   `json:"acc"`	// Acceleration of the particle in m.s-2
}

// NewParticle create and return a new instance of Charge struct.
func NewParticle(q, m float64) Particle {
	return Particle{
		Q:   q,
		M:   m,
		Pos: NewPoint(0, 0),
		Vel: NewForce(0, 0),
		Acc: NewForce(0, 0),
	}
}
