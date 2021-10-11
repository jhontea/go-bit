package main

import (
	"fmt"
	"strings"
)

func main() {
	result := findFirstStringInBracketRefactor("lis(ten)ing")
	fmt.Println(result)
}

func findFirstStringInBracketRefactor(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")

		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")

		if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 0 {
			runes := []rune(wordsAfterFirstBracket)
			return string(runes[1:indexClosingBracketFound])
		}
	}
	return ""
}
