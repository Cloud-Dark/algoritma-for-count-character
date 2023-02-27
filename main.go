package main

import (
	"fmt"
	"strings"
)

func countCharacters(inputStr string) map[string]int {
	counts := make(map[string]int)

	for i := 0; i < len(inputStr); i++ {
		char := string(inputStr[i])
		counts[char] = counts[char] + 1
	}

	return counts
}

func printCharacterCounts(inputStr string) string {
	counts := countCharacters(inputStr)
	var outputStr strings.Builder

	for char, count := range counts {
		if count == 1 {
			outputStr.WriteString(char)
		} else {
			outputStr.WriteString(fmt.Sprintf("%d%s", count, char))
		}
	}

	return outputStr.String()
}

func removeSpaces(inputStr string) string {
	return strings.ReplaceAll(inputStr, " ", "")
}

func changetolowercase(input string) string {
	return strings.ToLower(input)
}

func main() {
	input := "dani Maulana"
	fmt.Printf("text yang anda masukkan adalah: %s\n", input)
	output := printCharacterCounts(removeSpaces(changetolowercase(input)))
	fmt.Printf("Hasil setelah di kelola: %s\n", output)

	input2 := "Mas Syahdan Filsafan"
	fmt.Printf("text yang anda masukkan adalah: %s\n", input2)
	output2 := printCharacterCounts(removeSpaces(changetolowercase(input2)))
	fmt.Printf("Hasil setelah di kelola: %s\n", output2)
}
