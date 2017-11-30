package computer

// Spec is the computer specification
type Spec struct { // exported struct (Capitalized)
	Maker string // exported field (Capitalized)
	model string // not-exported field
	Price int    // exported field (Capitalized)
}
