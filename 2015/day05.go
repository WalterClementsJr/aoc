package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	vowelsRegex = regexp.MustCompile(`[aeiou]`)
	rule3       = regexp.MustCompile(`(ab|cd|pq|xy)`)
)

func main() {
	file, _ := os.Open("./day05.txt")
	scanner := bufio.NewScanner(file)

	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		isNice := checkIsNiceStringPartTwo(line)

		if isNice {
			niceCount += 1
		}
		fmt.Printf("%s is nice:%t\n", line, isNice)
	}
	fmt.Println(niceCount)
	file.Close()
}

func checkIsNiceStringPartOne(input string) bool {
	isRule1Pass := len(vowelsRegex.FindAllString(input, -1)) >= 3
	isRule2Pass := checkRule2(input)
	isRule3Pass := rule3.MatchString(input)

	return isRule1Pass && isRule2Pass && !isRule3Pass
}

func checkIsNiceStringPartTwo(input string) bool {
	ruleOne := false
	ruleTwo := false

	ruleOneMap := make(map[string]int)

	for i := 0; i < len(input)-1; i++ {
		current2Letter := string(rune(input[i])) + string(rune(input[i+1]))

		if (ruleOneMap[current2Letter]) == 0 {
			ruleOneMap[current2Letter] = i + 1
		} else {
			// if occurs right after
			if ruleOneMap[current2Letter]-i == 0 {
				continue
			}
			ruleOne = true
			break
		}
	}

	for i := 0; i < len(input)-2; i++ {
		if rune(input[i]) == rune(input[i+2]) {
			ruleTwo = true
		}
	}

	return ruleOne && ruleTwo
}

func checkRule2(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}
