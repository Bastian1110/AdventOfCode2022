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

func summ(arr []int) int {
	total := 0
	for _, a := range arr {
		total += a
	}
	return total
}

func getPairs(data string) [][]string {
	var res [][]string
	pairs := strings.Split(data, ",")
	for _, pair := range pairs {
		res = append(res, strings.Split(pair, "-"))
	}
	return res
}

func convertToInt(data []string) []int {
	var res []int
	for _, i := range data {
		number, _ := strconv.Atoi(i)
		res = append(res, number)
	}
	return res
}

func checkContain(data [][]string) bool {
	pairOne := convertToInt(data[0])
	pairTwo := convertToInt(data[1])
	if (pairOne[0]) <= int(pairTwo[0]) && int(pairOne[1]) >= int(pairTwo[1]) {
		return true
	}
	if (pairTwo[0]) <= int(pairOne[0]) && int(pairTwo[1]) >= int(pairOne[1]) {
		return true
	}
	return false
}

func solutionOne(data []string) {
	x := 0
	for _, line := range data {
		pairs := getPairs(line)
		if checkContain(pairs) {
			x += 1
		}
	}
	fmt.Println(x)
}

func getNumbers(line [][]string) []int {
	var numbers []int
	pairOne := convertToInt(line[0])
	pairTwo := convertToInt(line[1])
	for i := pairOne[0]; i <= pairOne[1]; i++ {
		numbers = append(numbers, i)
	}
	for i := pairTwo[0]; i <= pairTwo[1]; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func overlap(data []int) bool {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] && i != j {
				return true
			}
		}
	}
	return false
}

func solutionTwo(data []string) {
	x := 0
	for _, line := range data {
		pairs := getPairs(line)
		if overlap(getNumbers(pairs)) {
			x += 1
		}
	}
	fmt.Println(x)
}

func main() {
	lines := ReadFile("test.txt")
	//solutionOne((lines))
	solutionTwo(lines)
}
