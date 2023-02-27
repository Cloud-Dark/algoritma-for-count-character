package main

import (
	"fmt"
	"strings"
)

func countCharacters(inputStr string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range inputStr {
		counts[char] = counts[char] + 1
	}

	return counts
}

func printCharacterCounts(inputStr string) string {
	counts := countCharacters(inputStr)
	var outputStr strings.Builder

	for char, count := range counts {
		if count == 1 {
			outputStr.WriteRune(char)
		} else {
			outputStr.WriteString(fmt.Sprintf("%d%c", count, char))
		}
	}

	return outputStr.String()
}

func removeSpaces(inputStr string) string {
	return strings.ReplaceAll(inputStr, " ", "")
}

func changeToLowercase(inputStr string) string {
	return strings.ToLower(inputStr)
}

func main() {
	input := "dani Maulana"
	output := printCharacterCounts(removeSpaces(changeToLowercase(input)))
	fmt.Println(output)

	input2 := "SYahdan"
	output2 := printCharacterCounts(removeSpaces(changeToLowercase(input2)))
	fmt.Println(output2)
}
