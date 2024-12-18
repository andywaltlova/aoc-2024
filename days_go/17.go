package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type registers struct {
	a, b, c int
}

func getOperandValue(operand int, regs *registers) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return regs.a
	case 5:
		return regs.b
	case 6:
		return regs.c
	default:
		panic("Invalid combo operand")
	}
}

func executeProgram(program []int, regs *registers) string {
	instructionPointer := 0
	output := ""

	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]

		switch opcode {
		case 0: // adv
			denominator := int(math.Pow(2, float64(getOperandValue(operand, regs))))
			if denominator != 0 {
				regs.a /= denominator
			}
		case 1: // bxl
			regs.b ^= operand
		case 2: // bst
			regs.b = getOperandValue(operand, regs) % 8
		case 3: // jnz
			if regs.a != 0 {
				instructionPointer = operand
				continue
			}
		case 4: // bxc
			regs.b ^= regs.c
		case 5: // out
			value := getOperandValue(operand, regs) % 8
			if output != "" {
				output += ","
			}
			output += fmt.Sprint(value)
		case 6: // bdv
			denominator := int(math.Pow(2, float64(getOperandValue(operand, regs))))
			if denominator != 0 {
				regs.b = regs.a / denominator
			}
		case 7: // cdv
			denominator := int(math.Pow(2, float64(getOperandValue(operand, regs))))
			if denominator != 0 {
				regs.c = regs.a / denominator
			}
		default:
			panic("Unknown opcode")
		}
		instructionPointer += 2
	}

	return output
}

func parseInput(filePath string) (*registers, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var regs *registers
	var a, b, c int
	var program []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register") {
			values := strings.Fields(line)
			label := strings.TrimSuffix(values[1], ":")
			value, _ := strconv.Atoi(values[2])
			switch label {
			case "A":
				a = value
			case "B":
				b = value
			case "C":
				c = value
			}

		} else if strings.HasPrefix(line, "Program:") {
			values := strings.Split(line[9:], ",")
			for _, v := range values {
				num, _ := strconv.Atoi(strings.TrimSpace(v))
				program = append(program, num)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	regs = &registers{a: a, b: b, c: c}
	return regs, program, nil
}

func main() {
	regs, program, _ := parseInput("../data/17.txt")
	output := executeProgram(program, regs)
	fmt.Println(output)

}
