package main

import (
	"fmt"
	"unicode"
)

func isAlphaNumerical(r rune) bool {
	return (unicode.IsDigit(r) || unicode.IsLetter(r))
}

func isPalindrome(s string) bool {
	runeString := []rune(s)
	if len(runeString) <= 1 {
		return true
	}
	for st, end := 0, len(runeString)-1; st < end;{
		r1, r2 := runeString[st], runeString[end]
		if !isAlphaNumerical(r1){
			st ++
			continue;
		}
		if !isAlphaNumerical(r2){
			end --
			continue;
		}
		if unicode.ToLower(r1) != unicode.ToLower(r2){
			return false
		}
		st ++
		end --

	}
	return true
}

func main() {

	s := "ab"
	fmt.Print(isPalindrome(s))
}