package solutions

import (
	"fmt"
)

// ParenthesisMatch returns true if all opened parenthesis have a corresponding closed parenthesis in the right order.
// Use a stack and go over the string linearly.
func ParenthesisMatch(input string) bool {
	stack := RunesStack{}

	for i, char := range input {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			r, err := stack.Pop()
			if err != nil {
				fmt.Printf("Error: found a problem at char [%d] in input string: %s\n", i, err)
				return false
			}

			if char != getOppositeParenthesis(r) {
				fmt.Printf("Error: found a problem at char [%d] in input string: opening brackets '%c', closing brackets '%c' mismatch\n", i, r, char)
				return false
			}
		}
	}

	if !stack.Empty() {
		fmt.Printf("Error: found a problem with input string: stack still contains elements [count: %d]\n", stack.Length())
		return false
	}

	return true
}

// getOppositeParenthesis returns the opposite parenthesis.
func getOppositeParenthesis(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return 0
	}
}

// Rune stack is a stack of runes.
// The zero value is useful, it represents and empty stack.
type RunesStack struct {
	stack []rune
}

func (rs *RunesStack) Length() int {
	return len(rs.stack)
}

func (rs *RunesStack) Empty() bool {
	return rs.Length() == 0
}

func (rs *RunesStack) Push(r rune) {
	rs.stack = append(rs.stack, r)
}

func (rs *RunesStack) Pop() (rune, error) {
	// if empty return error
	if rs.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	elem := rs.stack[len(rs.stack)-1]

	rs.stack = rs.stack[:len(rs.stack)-1]

	return elem, nil
}
