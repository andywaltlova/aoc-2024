package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type button struct {
	x, y int
	cost int
}

type clawMachine struct {
	buttonA button
	buttonB button
	targetX int
	targetY int
}

func parseInput(filename string) ([]clawMachine, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var machines []clawMachine

	parseButton := func(line string, label string, cost int) button {
		line = strings.ReplaceAll(line, ",", "")
		parts := strings.Fields(strings.ReplaceAll(line, "Button "+label+":", ""))
		x, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "X+"))
		y, _ := strconv.Atoi(strings.TrimPrefix(parts[1], "Y+"))
		return button{x: x, y: y, cost: cost}
	}

	for i := 0; i < len(lines); i += 4 {
		buttonA := parseButton(lines[i], "A", 3)
		buttonB := parseButton(lines[i+1], "B", 1)

		prizeLine := strings.ReplaceAll(lines[i+2], "Prize:", "")
		prizeLine = strings.ReplaceAll(prizeLine, ",", "")

		targetParts := strings.Fields(prizeLine)
		targetX, _ := strconv.Atoi(strings.TrimPrefix(targetParts[0], "X="))
		targetY, _ := strconv.Atoi(strings.TrimPrefix(targetParts[1], "Y="))

		machines = append(machines, clawMachine{
			buttonA: buttonA,
			buttonB: buttonB,
			targetX: targetX,
			targetY: targetY,
		})
	}
	return machines, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return int(math.Abs(float64(a)))
}

func (c *clawMachine) solveUsingCramer() (int, int, bool) {
	a_x, b_x := c.buttonA.x, c.buttonB.x
	a_y, b_y := c.buttonA.y, c.buttonB.y
	X, Y := c.targetX, c.targetY

	// Determinant of the coefficient matrix
	delta := a_x*b_y - a_y*b_x
	if delta == 0 {
		// No unique solution exists
		return -1, -1, false
	}

	delta_m := X*b_y - Y*b_x
	delta_n := a_x*Y - a_y*X
	if delta_m%delta != 0 || delta_n%delta != 0 {
		return -1, -1, false
	}

	m := delta_m / delta
	n := delta_n / delta

	if m < 0 || n < 0 {
		return -1, -1, false
	}

	return m, n, true
}

func (cm *clawMachine) getMinimumTokens() (bool, int) {
	gcdX := gcd(cm.buttonA.x, cm.buttonB.x)
	gcdY := gcd(cm.buttonA.y, cm.buttonB.y)

	if cm.targetX%gcdX != 0 || cm.targetY%gcdY != 0 {
		return false, -1
	}

	// linear combinations
	minTokens := math.MaxInt64
	for pressesA := 0; pressesA <= cm.targetX/cm.buttonA.x; pressesA++ {

		remainingX := cm.targetX - pressesA*cm.buttonA.x
		remainingY := cm.targetY - pressesA*cm.buttonA.y

		if remainingX%cm.buttonB.x == 0 && remainingY%cm.buttonB.y == 0 {
			pressesB := remainingX / cm.buttonB.x
			if remainingY/cm.buttonB.y == pressesB {
				tokens := pressesA*cm.buttonA.cost + pressesB*cm.buttonB.cost
				if tokens < minTokens {
					minTokens = tokens
				}
			}
		}
	}

	if minTokens == math.MaxInt64 {
		return false, -1
	}

	return true, minTokens
}

func part1(machines []clawMachine) int {
	totalTokens := 0
	for _, cm := range machines {
		// if ok, tokens := cm.getMinimumTokens(); ok {
		// 	totalTokens += tokens
		// }
		Apresses, Bpresses, possible := cm.solveUsingCramer()
		if possible {
			totalTokens += Apresses*cm.buttonA.cost + Bpresses*cm.buttonB.cost
		}
	}
	return totalTokens
}

func part2(machines []clawMachine) int {
	totalTokens := 0
	for _, cm := range machines {
		cm.targetX += 10000000000000
		cm.targetY += 10000000000000
		Apresses, Bpresses, possible := cm.solveUsingCramer()
		if possible {
			totalTokens += Apresses*cm.buttonA.cost + Bpresses*cm.buttonB.cost
		}
	}
	return totalTokens
}

func main() {
	machines, err := parseInput("../data/13.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part1(machines))
	fmt.Println(part2(machines))

}
