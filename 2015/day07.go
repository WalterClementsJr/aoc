package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// store calculated wire
var cache map[string]uint16 = make(map[string]uint16)

func main() {
	wires := make(map[string]string)

	file, _ := os.Open("./day07.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		comm := strings.Split(line, "->")

		expression := strings.Trim(comm[0], " ")
		wire := strings.Trim(comm[1], " ")

		// since a wire is only assigned once, we can do this
		wires[wire] = expression
	}
	aValue := evaluateWire("a", wires)
	fmt.Println(aValue)

	// part 2: assign value of wire a to b, then run calculation of wire a again
	cache = make(map[string]uint16)
	wires["b"] = strconv.Itoa(int(aValue))

	fmt.Println(evaluateWire("a", wires))
}

func evaluateWire(wire string, wires map[string]string) uint16 {
	// search the cache
	if cachedValue, exist := cache[wire]; exist {
		return cachedValue
	}
	// try to parse number
	value, err := strconv.Atoi(wire)
	if err == nil {
		return uint16(value)
	}
	// single variable assignment
	val, exist := wires[wire]
	if exist == true && val != "" {
		return evaluateWire(wires[wire], wires)
	}

	var result uint16
	// wire is a gate instruction
	ins := strings.Split(wire, " ")
	if strings.Contains(wire, "NOT") {
		result = ^evaluateWire(ins[1], wires)
	} else if strings.Contains(wire, "AND") {
		result = evaluateWire(ins[0], wires) & evaluateWire(ins[2], wires)
	} else if strings.Contains(wire, "OR") {
		result = evaluateWire(ins[0], wires) | evaluateWire(ins[2], wires)
	} else if strings.Contains(wire, "LSHIFT") {
		result = evaluateWire(ins[0], wires) << evaluateWire(ins[2], wires)
	} else if strings.Contains(wire, "RSHIFT") {
		result = evaluateWire(ins[0], wires) >> evaluateWire(ins[2], wires)
	} else {
		panic("Not supported")
	}
	cache[wire] = result

	return result
}
