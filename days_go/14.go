package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

type velocity struct {
	x, y int
}

type robot struct {
	c coordinate
	v velocity
}

func (r *robot) move(n int, maxX int, maxY int) {
	for i := 0; i < n; i++ {
		// if new position is out of bounds, robot teleports to the opposite side
		r.c.x = (r.c.x + r.v.x + maxX) % maxX
		r.c.y = (r.c.y + r.v.y + maxY) % maxY
	}
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
		if r.c.x < verticalMiddle && r.c.y < horizontalMiddle {
			q1 += 1
		} else if r.c.x > verticalMiddle && r.c.y < horizontalMiddle {
			q2 += 1
		} else if r.c.x < verticalMiddle && r.c.y > horizontalMiddle {
			q3 += 1
		} else if r.c.x > verticalMiddle && r.c.y > horizontalMiddle {
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
			c: coordinate{x: x, y: y},
			v: velocity{x: vx, y: vy},
		})
	}
	return robots, nil
}

func part2(robots []robot, maxX int, maxY int) int {
	n := 0

	// Get while cycle to break by user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		n++
		for i := range robots {
			(&robots[i]).move(1, maxX, maxY)
		}

		// Print final positions
		robotMap := make(map[coordinate]int)
		for _, r := range robots {
			robotMap[coordinate{x: r.c.x, y: r.c.y}] += 1
		}

		for y := 0; y < maxY; y++ {
			line := 0
			for x := 0; x < maxX; x++ {
				// Eh what do I know, maybe there is a line of robots?
				if robotMap[coordinate{x: x, y: y}] > 0 {
					line++
				} else {
					if line > 12 {
						// Check picture
						for y := 0; y < maxY; y++ {
							for x := 0; x < maxX; x++ {
								if robotMap[coordinate{x: x, y: y}] > 0 {
									fmt.Print("#")
								} else {
									fmt.Print(".")
								}
							}
							fmt.Println()
						}

						fmt.Printf("(%d) Enter command (type 'finish' to exit): ", n)
						scanner.Scan()
						input := strings.TrimSpace(scanner.Text())
						if input == "finish" {
							fmt.Println("Merry Christmas!")
							break
						}
					}
					line = 0
				}
			}

		}

	}
}

func main() {
	// robots, _ := parseRobots("../data/14_test.txt")
	// fmt.Println(part1(robots, 100, 11, 7)) // Test
	robots, _ := parseRobots("../data/14.txt")
	// fmt.Println(part1(robots, 100, 101, 103))
	part2(robots, 101, 103)

}
