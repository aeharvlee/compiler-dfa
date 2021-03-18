package hard_wired

import "github.com/aeharvlee/hardwired-dfa/dfa"

type hardWiredDFA struct {
}

func NewHardWiredDFA() dfa.DFA {
	return &hardWiredDFA{}
}

func (d *hardWiredDFA) Recognize(input string) bool {
	currentState := P
	pos := 0

	for pos < len(input)-1 {
		ch := input[pos]

		// Only 0 or 1 is allowed as input character.
		if ch != '0' && ch != '1' {
			return false
		}

		switch currentState {
		case P:
			if ch == '0' {
				currentState = Q
				break
			}
		case Q:
			if ch == '0' {
				currentState = R
				break
			}
			if ch == '1' {
				currentState = P
				break
			}
		}
		pos++
	}

	if currentState == R {
		return true
	}
	return false
}