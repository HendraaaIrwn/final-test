package main

import "strings"

func main() {
	stringWord1 := "MalaM"
	stringWord2 := "levEl"
	stringWord3 := "kaSur iNi ruSak"
	stringWord4 := "saYur"

	println(isPalindrome(stringWord1))
	println(isPalindrome(stringWord2))
	println(isPalindrome(stringWord3))
	println(isPalindrome(stringWord4))

}

func isPalindrome(s string) bool {
	loweredWord := []rune(strings.ToLower(s))

	for i := 0; i < len(s)/2; i++ {
		if loweredWord[i] != loweredWord[len(s)-i-1] {
			return false
		}
	}
	return true
}
