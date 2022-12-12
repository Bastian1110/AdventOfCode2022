package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	var res []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cant open the file 😬")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

type End struct {
	x, y int
	id   string
}

func (h *End) init(x int, y int, id string) *End {
	h.x = x
	h.y = y
	h.id = id
	return h
}

func NewEnd(x int, y int, id string) *End {
	return new(End).init(x, y, id)
}

func parseDirection(c string) []int {
	switch c {
	case "U":
		return []int{0, 1}
	case "D":
		return []int{0, -1}
	case "R":
		return []int{1, 0}
	case "L":
		return []int{-1, 0}
	default:
		return []int{0, 0}
	}
}

func Print(width int, height int, h *End, t *End) {
	matrix := [][]string{}
	for i := 0; i < height; i++ {
		temp := []string{}
		for j := 0; j < width; j++ {
			temp = append(temp, ".")
		}
		matrix = append(matrix, temp)
	}

	matrix[t.y][t.x] = t.id
	matrix[h.y][h.x] = h.id

	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("=")
}

func (e *End) Step(instruction []string, t *End) [][]int {
	tailPositions := [][]int{}
	direction := parseDirection(instruction[0])
	magnitud, _ := strconv.Atoi(instruction[1])
	for i := 0; i < magnitud; i++ {
		e.x += direction[0]
		e.y += direction[1]
		if i >= 1 {
			followX := 0
			followY := 0
			directionX := e.x - t.x
			directionY := e.y - t.y
			if directionX < 0 {
				followX = -1
			} else if directionX > 0 {
				followX = 1
			} else {
				followX = 0
			}
			if directionY < 0 {
				followY = -1
			} else if directionY > 0 {
				followY = 1
			} else {
				followY = 0
			}
			if t.x+followX == e.x && t.y+followY == e.y {
				followX = 0
				followY = 0
			}
			t.x += followX
			t.y += followY
			tailPositions = append(tailPositions, []int{t.x, t.y})
		}
		//Print(6, 5, e, t)
	}
	return tailPositions
}

func isIn(list [][]int, check []int) bool {
	for _, cord := range list {
		if cord[0] == check[0] && cord[1] == check[1] {
			return true
		}
	}
	return false
}

func printResult(positions [][]int, width int, height int) {
	matrix := [][]string{}
	for i := 0; i < height; i++ {
		temp := []string{}
		for j := 0; j < width; j++ {
			temp = append(temp, ".")
		}
		matrix = append(matrix, temp)
	}

	for _, h := range positions {
		matrix[h[1]][h[0]] = "#"
	}

	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("=")

}

func solutionOne(data []string) {
	head := NewEnd(0, 0, "H")
	tail := NewEnd(0, 0, "T")
	totalTos := [][]int{}
	for _, line := range data {
		splited := strings.Split(line, " ")
		fmt.Println(splited)
		tailPositions := head.Step(splited, tail)
		for _, pos := range tailPositions {
			if !isIn(totalTos, pos) {
				totalTos = append(totalTos, pos)
			}
		}
	}
	//printResult(totalTos, 6, 5)
	fmt.Println(len(totalTos) + 1)
}

func solutionTwo(data []string) {

}

func main() {
	lines := ReadFile("test.txt")
	solutionOne(lines)

}
