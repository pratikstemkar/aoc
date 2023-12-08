package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func readInput(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)

	return str
}

func part1() int {
	str := readInput("input.txt")

	strList := strings.Split(str, "\n")
	strList = strList[:len(strList)-1]

	total := 0
	for _, str := range strList {
		var firstDigit rune
		for _, char := range str {
			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		var lastDigit rune
		for i := len(str) - 1; i >= 0; i-- {
			char := rune(str[i])
			if unicode.IsDigit(char) {
				lastDigit = char
				break
			}
		}

		firstInt := int(firstDigit - '0')
		lastInt := int(lastDigit - '0')

		number := firstInt*10 + lastInt

		total += number
	}

	return total
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
