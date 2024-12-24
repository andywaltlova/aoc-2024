package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gate struct {
	input1    string
	input2    string
	operation string
	output    string
}

type wire struct {
	name  string
	value int
}

type dependencyGraph struct {
	gates      map[string]gate
	wireValues map[string]int
	dependents map[string][]string // Tracks which wires depend on a given wire
	evaluated  map[string]bool
}

func newDependencyGraph() *dependencyGraph {
	return &dependencyGraph{
		gates:      make(map[string]gate),
		wireValues: make(map[string]int),
		dependents: make(map[string][]string),
		evaluated:  make(map[string]bool),
	}
}

func (dg *dependencyGraph) addGate(input1, input2, operation, output string) {
	dg.gates[output] = gate{input1, input2, operation, output}
	dg.dependents[input1] = append(dg.dependents[input1], output)
	dg.dependents[input2] = append(dg.dependents[input2], output)
}

func (dg *dependencyGraph) setWireValue(wire string, value int) {
	dg.wireValues[wire] = value
	dg.evaluated[wire] = true
}

func (dg *dependencyGraph) evaluate(wire string) int {
	if dg.evaluated[wire] {
		return dg.wireValues[wire]
	}

	gate, _ := dg.gates[wire] // Asuming input is correct and gate exists

	val1 := dg.evaluate(gate.input1)
	val2 := dg.evaluate(gate.input2)
	result := compute(gate.operation, val1, val2)

	// Cache
	dg.wireValues[wire] = result
	dg.evaluated[wire] = true
	return result
}

func compute(op string, val1, val2 int) int {
	switch op {
	case "AND":
		return val1 & val2
	case "OR":
		return val1 | val2
	case "XOR":
		return val1 ^ val2
	default:
		panic("Unknown operation: " + op)
	}
}

func (dg *dependencyGraph) collectZOutputs() string {
	var binaryString strings.Builder
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if _, exists := dg.gates[wire]; !exists {
			break
		}
		value := dg.evaluate(wire)
		binaryString.WriteString(strconv.Itoa(value))
	}

	// Reverse the string (least significant bit z00 first)
	binaryStringString := binaryString.String()
	binaryString.Reset()
	for i := len(binaryStringString) - 1; i >= 0; i-- {
		binaryString.WriteByte(binaryStringString[i])
	}
	return binaryString.String()
}

func parseInput(filename string) ([]wire, []gate) {
	data, _ := os.ReadFile(filename)
	parts := strings.Split(string(data), "\n\n")

	wires := strings.Split(parts[0], "\n")
	var wiresResult []wire
	for _, w := range wires {
		values := strings.Split(w, ": ")
		name := values[0]
		value, _ := strconv.Atoi(values[1])
		wiresResult = append(wiresResult, wire{name, value})
	}
	gates := strings.Split(strings.TrimSpace(parts[1]), "\n")
	var gatesResult []gate
	for _, g := range gates {
		values := strings.Split(g, " ")
		gatesResult = append(
			gatesResult,
			gate{values[0], values[2], values[1], values[4]},
		)
	}
	return wiresResult, gatesResult
}

func main() {
	graph := newDependencyGraph()

	wires, gates := parseInput("../data/24.txt")

	for _, wire := range wires {
		graph.setWireValue(wire.name, wire.value)
	}
	for _, gate := range gates {
		graph.addGate(gate.input1, gate.input2, gate.operation, gate.output)
	}

	binaryString := graph.collectZOutputs()
	fmt.Println("Binary string:", binaryString)
	result, _ := strconv.ParseInt(binaryString, 2, 64)
	fmt.Printf("Output decimal number: %d\n", result)

}
