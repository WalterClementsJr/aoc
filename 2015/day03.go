package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("./day03.txt")
	directions := string(file)
	// directions = "^>v<"

	mapGrid := make(map[int]map[int]int)

	var x, y int
	var roboX, roboY int

	isSantaTurn := true

	// delivery present at starting location
	recordDelivery(mapGrid, x, y)
	recordDelivery(mapGrid, roboX, roboY)

	for _, direction := range directions {
		if isSantaTurn {
			applyDirection(direction, &x, &y)
			recordDelivery(mapGrid, x, y)
		} else {
			applyDirection(direction, &roboX, &roboY)
			recordDelivery(mapGrid, roboX, roboY)
		}
		isSantaTurn = !isSantaTurn
	}

	count := 0
	for xIndex := range mapGrid {
		for yIndex := range mapGrid[xIndex] {
			if mapGrid[xIndex][yIndex] >= 1 {
				count += 1
			}
		}
	}
	fmt.Println("houses with at least 1 present", count)
}

// increment the present delivered at current coordinate (x, y)
func recordDelivery(mapGrid map[int]map[int]int, x int, y int) {
	// init inner map if not present
	if mapGrid[x] == nil {
		mapGrid[x] = make(map[int]int)
	}
	mapGrid[x][y] = mapGrid[x][y] + 1
}

func applyDirection(dir rune, x, y *int) {
	switch dir {
	case '>':
		*x += 1
	case '<':
		*x -= 1
	case '^':
		*y += 1
	case 'v':
		*y -= 1
	}
}
