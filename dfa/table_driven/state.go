package table_driven

type State int

// Top of State means "Start state"
// Bottom of State means "Final state"
const (
	P State = iota
	Q
	R
)

func (s State) Char() byte {
	switch s {
	case P:
		return 'P'
	case Q:
		return 'Q'
	case R:
		return 'R'
	}
	return 'X'
}
