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

func isIn(list []int, number int) bool {
	for _, n := range list {
		if n == number {
			return true
		}
	}
	return false
}

func solutionOne(data []string) {
	x := 1
	cycles := 1
	cyclesInfo := map[int]int{}
	for i := 0; i < len(data); i++ {
		splited := strings.Split(data[i], " ")
		if splited[0] == "noop" {
			cycles += 1
			cyclesInfo[cycles] = x
		} else {
			cycles += 1
			cyclesInfo[cycles] = x
			number, _ := strconv.Atoi(splited[1])
			x += number
			cycles += 1
			cyclesInfo[cycles] = x
		}
	}
	sum := 0
	for i := 20; i <= 240; i += 40 {
		sum += cyclesInfo[i] * i
	}
	fmt.Println(sum)
}

func solutionTwo(data []string) {
}

func main() {
	lines := ReadFile("test.txt")
	solutionOne(lines)

}
