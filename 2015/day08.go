package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	totalCodelen := 0
	totalChar := 0
	totalEncodedCharLen := 0

	file, _ := os.Open("day08.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		originalCodeLen := len(input)

		decodedString := decodeToInMemoryString(input)
		encodedString := encodeString(input)
		fmt.Println(encodedString)

		totalCodelen += originalCodeLen
		// get actual char count, not byte count
		totalChar += utf8.RuneCount([]byte(decodedString))
		totalEncodedCharLen += len(encodedString)
	}

	fmt.Println("results 1", totalCodelen-totalChar)
	fmt.Println("results 2", totalEncodedCharLen, totalCodelen, totalEncodedCharLen-totalCodelen)
}

var hexaRegex = regexp.MustCompile(`\\x[a-f0-9]{2}`)

func decodeToInMemoryString(input string) string {
	// trim quotes at beginning
	actualStr := strings.TrimSuffix(input, `"`)
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		actualStr = input[1 : len(input)-1]
	}

	actualStr = strings.ReplaceAll(actualStr, `\"`, `"`)
	actualStr = strings.ReplaceAll(actualStr, `\\`, `\`)

	hexaMatches := hexaRegex.FindAllStringSubmatchIndex(actualStr, -1)

	if len(hexaMatches) > 0 {
		for i := len(hexaMatches) - 1; i >= 0; i-- {
			match := hexaMatches[i]

			// will be something like \xFF
			hexaStr := actualStr[match[0]:match[1]]
			hexaCode := hexaStr[2:]
			decimalValue, err := strconv.ParseInt(hexaCode, 16, 0)
			if err != nil {
				panic("Error decoding hex")
			}
			char := string(rune(decimalValue))
			actualStr = actualStr[:match[0]] + char + actualStr[match[1]:]
		}
	}

	return actualStr
}

func encodeString(input string) string {
	input = strings.ReplaceAll(input, `\`, `\\`)
	input = strings.ReplaceAll(input, `"`, `\"`)

	return `"` + input + `"`
}
