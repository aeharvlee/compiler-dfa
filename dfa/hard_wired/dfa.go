package hard_wired

import (
	"fmt"

	"github.com/aeharvlee/compiler-dfa/dfa"
)

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

		var newState State
		switch currentState {
		case P:
			if ch == '0' {
				newState = Q
				break
			}
		case Q:
			if ch == '0' {
				newState = R
				break
			}
			if ch == '1' {
				newState = P
				break
			}
		case R:
			if ch == '0' {
				newState = R
				break
			}
			if ch == '1' {
				newState = R
				break
			}
		}

		fmt.Printf("ch: %c | State Transition from %c to %c\n", ch, currentState.Char(), newState.Char())
		currentState = newState

		pos++
	}

	if currentState == R {
		return true
	}
	return false
}
