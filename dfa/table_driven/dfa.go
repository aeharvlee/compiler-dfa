package table_driven

import (
	"fmt"

	"github.com/aeharvlee/compiler-dfa/dfa"
)

type tableDrivenDFA struct {
	StateTransitionDiagram map[State]map[byte]State
}

func NewTableDrivenDFA() dfa.DFA {
	return &tableDrivenDFA{
		StateTransitionDiagram: map[State]map[byte]State{
			P: {
				'0': Q,
				'1': P,
			},
			Q: {
				'0': R,
				'1': P,
			},
			R: {
				'0': R,
				'1': R,
			},
		},
	}
}

func (t *tableDrivenDFA) Recognize(input string) bool {
	currentState := P
	pos := 0

	for pos < len(input)-1 {
		ch := input[pos]

		// Only 0 or 1 is allowed as input character.
		if ch != '0' && ch != '1' {
			return false
		}

		// State Transition
		newState := t.StateTransitionDiagram[currentState][ch]
		fmt.Printf("ch: %c | State Transition from %c to %c\n", ch, currentState.Char(), newState.Char())
		currentState = newState

		pos++
	}

	if currentState == R {
		return true
	}
	return false
}
