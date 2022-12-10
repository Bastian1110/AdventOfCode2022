package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func summ(arr []int) int {
	total := 0
	for _, a := range arr {
		total += a
	}
	return total
}

func makeMatrix(data []string) [][]int {
	res := [][]int{}
	for i := 0; i < len(data); i++ {
		temp := []int{}
		for j := 0; j < len(data[i]); j++ {
			number, _ := strconv.Atoi(string(data[i][j]))
			temp = append(temp, number)
		}
		res = append(res, temp)
	}
	return res
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("")
	}
}

func isVisible(number int, line []int) bool {
	for _, n := range line {
		if n >= number {
			return false
		}
	}
	return true
}

func checkY(matrix [][]int, x int, y int) bool {
	yAxis := []int{}
	for i := 0; i < len(matrix); i++ {
		yAxis = append(yAxis, matrix[i][x])
	}
	if isVisible(matrix[y][x], yAxis[y+1:]) {
		return true
	}
	if isVisible(matrix[y][x], yAxis[:y]) {
		return true
	}
	return false
}

func checkX(matrix [][]int, x int, y int) bool {
	if isVisible(matrix[y][x], matrix[y][x+1:]) {
		return true
	}
	if isVisible(matrix[y][x], matrix[y][:x]) {
		return true
	}
	return false
}

func solutionOne(data []string) {
	matrix := makeMatrix(data)
	total := (len(matrix)*2 - 4) + (len(matrix[0]) * 2)
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			if checkX(matrix, j, i) || checkY(matrix, j, i) {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func untilBlocked(line []int, inverted bool, n int) int {
	if inverted {
		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}
	}
	fmt.Println("Line", line)
	res := 1
	for i := 0; i < len(line); i++ {
		if line[i] >= n {
			break
		}
		if i+1 > len(line)-1 {
			break
		}
		if line[i] <= line[i+1] {
			res += 1
		}
	}
	return res
}

func getScenic(matrix [][]int, x int, y int) int {
	fmt.Println(matrix[y][x])
	scores := []int{}
	if x != len(matrix[0])-1 {
		east := untilBlocked(matrix[y][x+1:], false, matrix[y][x])
		fmt.Println("->", east)
		scores = append(scores, east)
	}
	if x != 0 {
		west := untilBlocked(matrix[y][:x], true, matrix[y][x])
		scores = append(scores, west)
		fmt.Println("<-", west)
	}
	yAxis := []int{}
	for i := 0; i < len(matrix); i++ {
		yAxis = append(yAxis, matrix[i][x])
	}
	if y != len(matrix)-1 {
		north := untilBlocked(yAxis[:y], true, matrix[y][x])
		scores = append(scores, north)
		fmt.Println("^", north)
	}
	if y != 0 {
		south := untilBlocked(yAxis[y+1:], false, matrix[y][x])
		scores = append(scores, south)
		fmt.Println("v", south)
	}
	score := 1
	for _, n := range scores {
		score = score * n
	}
	return score
}
func max(list []int) int {
	max := -1
	for _, n := range list {
		if n > max {
			max = n
		}
	}
	return max
}
func solutionTwo(data []string) {
	matrix := makeMatrix(data)
	scores := []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			scores = append(scores, getScenic(matrix, j, i))
		}
	}
	fmt.Println(max(scores))
}

func main() {
	lines := ReadFile("test.txt")
	solutionTwo(lines)

}
