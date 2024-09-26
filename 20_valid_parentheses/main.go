package main

import "fmt"

func main() {
	s := "([]{}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := []rune{}
	for _, symb := range s {
		if symb == '(' || symb == '[' || symb == '{' { //1 case: if i is "([{"
			stack = append(stack, symb)
		} else { //2 case if i not in "([{"
			if len(stack) == 0 { // when s like ")[]"
				return false
			}
			last := stack[len(stack)-1] // last symbol from stack
			stack = stack[:len(stack)-1]
			comb := string(last) + string(symb)
			if Valid(comb) {
				continue
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
func Valid(s string) bool {
	if s == "()" || s == "[]" || s == "{}" {
		return true
	}
	return false
}
