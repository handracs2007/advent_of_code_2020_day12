// https://adventofcode.com/2020/day/12/
package main

import (
    "fmt"
    "io/ioutil"
    "math"
    "strconv"
    "strings"
)

func main() {
    fc, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    var x, y, dx, dy int

    lines := strings.Split(string(fc), "\n")

    dx = 1 // Default to east
    part1(lines, x, dx, y, dy)

    dx = 10
    dy = -1
    part2(lines, x, dx, y, dy)
}

func part1(lines []string, x int, dx int, y int, dy int) {
    for _, line := range lines {
        cmd := line[0]
        dist, _ := strconv.Atoi(line[1:])

        switch cmd {
        case 'F':
            x += dist * dx
            y += dist * dy

        case 'N':
            y -= dist
        case 'S':
            y += dist
        case 'E':
            x += dist
        case 'W':
            x -= dist

        case 'L':
            switch dist {
            case 90:
                if dx == 1 && dy == 0 {
                    // Facing east
                    dx, dy = 0, -1
                } else if dx == -1 && dy == 0 {
                    // Facing west
                    dx, dy = 0, 1
                } else if dx == 0 && dy == -1 {
                    // Facing north
                    dx, dy = -1, 0
                } else if dx == 0 && dy == 1 {
                    // Facing south
                    dx, dy = 1, 0
                }

            case 180:
                dx *= -1
                dy *= -1

            case 270:
                if dx == 1 && dy == 0 {
                    // Facing east
                    dx, dy = 0, 1
                } else if dx == -1 && dy == 0 {
                    // Facing west
                    dx, dy = 0, -1
                } else if dx == 0 && dy == -1 {
                    // Facing north
                    dx, dy = 1, 0
                } else if dx == 0 && dy == 1 {
                    // Facing south
                    dx, dy = -1, 0
                }
            }

        case 'R':
            switch dist {
            case 90:
                if dx == 1 && dy == 0 {
                    // Facing east
                    dx, dy = 0, 1
                } else if dx == -1 && dy == 0 {
                    // Facing west
                    dx, dy = 0, -1
                } else if dx == 0 && dy == -1 {
                    // Facing north
                    dx, dy = 1, 0
                } else if dx == 0 && dy == 1 {
                    // Facing south
                    dx, dy = -1, 0
                }

            case 180:
                dx *= -1
                dy *= -1

            case 270:
                if dx == 1 && dy == 0 {
                    // Facing east
                    dx, dy = 0, -1
                } else if dx == -1 && dy == 0 {
                    // Facing west
                    dx, dy = 0, 1
                } else if dx == 0 && dy == -1 {
                    // Facing north
                    dx, dy = -1, 0
                } else if dx == 0 && dy == 1 {
                    // Facing south
                    dx, dy = 1, 0
                }
            }
        }
    }

    ax := int(math.Abs(float64(x)))
    ay := int(math.Abs(float64(y)))

    fmt.Println(ax + ay)
}

func part2(lines []string, x int, dx int, y int, dy int) {
    for _, line := range lines {
        cmd := line[0]
        dist, _ := strconv.Atoi(line[1:])

        switch cmd {
        case 'F':
            x += dist * dx
            y += dist * dy

        case 'N':
            dy -= dist
        case 'S':
            dy += dist
        case 'E':
            dx += dist
        case 'W':
            dx -= dist

        case 'L':
            switch dist {
            case 90:
                dx, dy = dy, -dx

            case 180:
                dx *= -1
                dy *= -1

            case 270:
                dx, dy = -dy, dx
            }

        case 'R':
            switch dist {
            case 90:
                dx, dy = -dy, dx

            case 180:
                dx *= -1
                dy *= -1

            case 270:
                dx, dy = dy, -dx
            }
        }
    }

    ax := int(math.Abs(float64(x)))
    ay := int(math.Abs(float64(y)))

    fmt.Println(ax + ay)
}
