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

func contains(arr []int, number int) bool {
   for _, a := range arr {
      if a == number {
         return true
      }
   }
   return false
}


func summ(arr []int) int {
	total := 0 
	for _, a := range arr {
      total += a
   }
   return total

}
func solution(data []string) {
	var top []int
	for i := 0; i < 3; i++ {
		sum := 0
		max := 0
		for _, line := range data {
			number, err := strconv.Atoi(line)
			if err != nil{
				if sum > max && !contains(top, sum) {
					max = sum
				}
				fmt.Println(sum)
				sum = 0
			} else {
				sum += number
			}
		}
		top = append(top, max)
	}
	fmt.Println(top)
	fmt.Println(summ(top))

}

func main() {
	lines := ReadFile("./test.txt")
	solution(lines)
}