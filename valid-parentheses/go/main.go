package main

import (
	"fmt"
)

// bracketPairs maps an opening bracket to its matching closing bracket.
var bracketPairs = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

// isValid checks if a string that contains brackets ('()', '{}', '[]') is valid,
// according to the following rules:
//
// 1. Open brackets must be closed by the same type of brackets.
//
// 2. Open brackets must be closed in the correct order.
func isValid(s string) bool {
	stack := runeStack{}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack.push(r)
		case ')', '}', ']':
			v, err := stack.pop()
			if err != nil {
				return false
			}

			if v != bracketPairs[r] {
				return false
			}
		}
	}

	return stack.count == 0
}

func main() {
	valid := isValid("((()))")
	fmt.Printf("valid: %v\n", valid)
}
