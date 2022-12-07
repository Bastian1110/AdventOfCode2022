package main

import (
	"bufio"
	"fmt"
	"os"
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

func repeated(data string) bool {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] && i != j {
				return true
			}
		}
	}
	return false
}

func solutionOne(data string) {
	for i := 0; i < len(data); i++ {
		if i+4 > len(data) {
			break
		}
		actual := data[i : i+4]
		fmt.Println(actual)
		if !repeated(actual) {
			fmt.Println(i + 4)
			break
		}
	}
}

func solutionTwo(data string) {
	for i := 0; i < len(data); i++ {
		if i+14 > len(data) {
			break
		}
		actual := data[i : i+14]
		fmt.Println(actual)
		if !repeated(actual) {
			fmt.Println(i + 14)
			break
		}
	}

}

func main() {
	lines := ReadFile("test.txt")
	//solutionOne(lines[0])
	solutionTwo(lines[0])
}
