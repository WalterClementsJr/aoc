package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var dimensionRegex = regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

const MAX_INT = int(^uint(0) >> 1) // max int

func main() {
	file, _ := os.Open("./day02-input.txt")
	scanner := bufio.NewScanner(file)

	totalWrapperFeet := 0
	totalRibbonLengthFeet := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		l, w, h, _ := parseDimensions(line)

		totalWrapperFeet = totalWrapperFeet + getSurfaceArea(l, w, h)
		totalRibbonLengthFeet = totalRibbonLengthFeet + calculateRibbonLength(l, w, h)
	}

	fmt.Println("total wrapping paper in square feet is", totalWrapperFeet)
	fmt.Println("total ribbon needed in feet is", totalRibbonLengthFeet)

	file.Close()
}

func parseDimensions(dimensionString string) (int, int, int, error) {
	matches := dimensionRegex.FindStringSubmatch(dimensionString)

	if len(matches) < 4 {
		return -1, -1, -1, errors.New("Wrong format")
	}
	length, _ := strconv.Atoi(matches[1])
	width, _ := strconv.Atoi(matches[2])
	height, _ := strconv.Atoi(matches[3])
	return length, width, height, nil
}

func getSurfaceArea(l int, w int, h int) int {
	surfaceArea := 2*l*w + 2*w*h + 2*l*h

	smallestSurface := min_func(l*w, w*h, l*h)

	return surfaceArea + smallestSurface
}

func calculateRibbonLength(l int, w int, h int) int {
	bowLength := l * w * h

	smallestPerimeter := min_func(2*(l+w), 2*(w+h), 2*(l+h))

	return bowLength + smallestPerimeter
}

func min_func(args ...int) int {
	smallestVal := MAX_INT
	for _, s := range args {
		if smallestVal > s {
			smallestVal = s
		}
	}
	return smallestVal
}
