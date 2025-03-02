package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	vowelsRegex = regexp.MustCompile(`[aeiou]`)
	// rule2       = regexp.MustCompile(`^(\w)(?=\1)$`)
	rule3 = regexp.MustCompile(`(ab|cd|pq|xy)`)
)

func main() {
	// nice, naughty
	// testSuit := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "uurcxstgmygtbstg", "ieodomkazucvgmuy"}

	// for _, input := range testSuit {
	// 	fmt.Println(input, checkIsNiceStringPartTwo(input))
	// }
	// return

	file, _ := os.Open("./day05.txt")
	scanner := bufio.NewScanner(file)

	niceCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		// isNice := checkIsNiceStringPartOne(line)
		isNice := checkIsNiceStringPartTwo(line)

		if isNice {
			niceCount += 1
		}
		fmt.Printf("%s is nice:%t\n", line, isNice)
	}
	fmt.Println(niceCount)
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

		// fmt.Println("\tcurrent letter", current2Letter, ruleOneMap[current2Letter])
		if (ruleOneMap[current2Letter]) == 0 {
			ruleOneMap[current2Letter] = i + 1
			// fmt.Println("\tnew key", ruleOneMap)
		} else {
			// if occurs right after
			if ruleOneMap[current2Letter]-i == 0 {
				continue
			}
			ruleOne = true
			// fmt.Println("\tfound rule 1 at index", i, ruleOneMap)
			break
		}
	}
	// fmt.Println(ruleOneMap)

	for i := 0; i < len(input)-2; i++ {
		if rune(input[i]) == rune(input[i+2]) {
			// fmt.Println("\tfound rule 2 at index", i, rune(input[i]))
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
