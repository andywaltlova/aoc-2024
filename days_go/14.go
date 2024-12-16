package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type velocity struct {
	x, y int
}

type robot struct {
	x, y int
	v    velocity
}

func (r *robot) move(n int, maxX int, maxY int) {
	for i := 0; i < n; i++ {
		// if new position is out of bounds, robot teleports to the opposite side
		r.x = (r.x + r.v.x + maxX) % maxX
		r.y = (r.y + r.v.y + maxY) % maxY
	}
	// Print final position
	// for y := 0; y < maxY; y++ {
	// 	for x := 0; x < maxX; x++ {
	// 		if r.x == x && r.y == y {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func part1(robots []robot, n int, maxX int, maxY int) int {
	for i := range robots {
		(&robots[i]).move(n, maxX, maxY)
	}

	// Get robots in each quadrant
	q1, q2, q3, q4 := 0, 0, 0, 0
	horizontalMiddle := maxY / 2
	verticalMiddle := maxX / 2
	for _, r := range robots {
		if r.x < verticalMiddle && r.y < horizontalMiddle {
			q1 += 1
		} else if r.x > verticalMiddle && r.y < horizontalMiddle {
			q2 += 1
		} else if r.x < verticalMiddle && r.y > horizontalMiddle {
			q3 += 1
		} else if r.x > verticalMiddle && r.y > horizontalMiddle {
			q4 += 1
		}
	}
	return q1 * q2 * q3 * q4
}

func parseRobots(filename string) ([]robot, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var robots []robot

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		parts := strings.Fields(lines[i])
		coordinates := strings.Split(strings.TrimPrefix(parts[0], "p="), ",")
		velocityCoord := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		vx, _ := strconv.Atoi(velocityCoord[0])
		vy, _ := strconv.Atoi(velocityCoord[1])

		robots = append(robots, robot{
			x: x,
			y: y,
			v: velocity{x: vx, y: vy},
		})
	}
	return robots, nil
}

func main() {
	// robots, _ := parseRobots("../data/14_test.txt")
	// fmt.Println(part1(robots, 100, 11, 7)) // Test
	robots, _ := parseRobots("../data/14.txt")
	fmt.Println(part1(robots, 100, 101, 103))
}
