package dfa

type DFA interface {
	Recognize(input string) bool
}