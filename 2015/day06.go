package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Instruction struct {
	op    string
	start Point
	end   Point
}

var instructionRegex = regexp.MustCompile(`^([\w\s]*) ([\d,]*) through ([\d,]*)$`)

func main() {
	file, _ := os.Open("./day06.txt")
	scanner := bufio.NewScanner(file)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		res := instructionRegex.FindAllStringSubmatch(line, -1)

		aCoord := strings.Split(res[0][2], ",")

		x, _ := strconv.Atoi(aCoord[0])
		y, _ := strconv.Atoi(aCoord[1])

		pointA := Point{
			x: x,
			y: y,
		}

		bCoord := strings.Split(res[0][3], ",")

		x, _ = strconv.Atoi(bCoord[0])
		y, _ = strconv.Atoi(bCoord[1])

		pointB := Point{
			x: x,
			y: y,
		}

		op := res[0][1]

		for i := pointA.x; i <= pointB.x; i++ {
			for j := pointA.y; j <= pointB.y; j++ {
				switch op {
				case "turn on":
					grid[i][j] += 1
				case "turn off":
					if grid[i][j] <= 0 {
						continue
					}
					grid[i][j] -= 1
				case "toggle":
					grid[i][j] += 2
				}
			}
		}
	}

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += grid[i][j]
		}
	}
	fmt.Println(count)
	file.Close()
}
