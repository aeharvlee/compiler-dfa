package hard_wired

type State int

// Top of State means "Start state"
// Bottom of State means "Final state"
const (
	P State = iota
	Q
	R
)
