package main

import (
	"bufio"
	"fmt"
	"os"
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

func winner(first string, second string) int{
	gameRulesOne := map[string] string {"A" : "Y", "B" : "Z", "C" : "X"}
	gameRulesTwo := map[string] string {"X" : "B", "Y" : "C", "Z" : "A"}
	if gameRulesOne[first] == second{
		return 6
	}
	if gameRulesTwo[second] == first {
		return 0
	}
	return 3
}

func summ(arr []int) int {
    total := 0 
    for _, a := range arr {
      total += a
   	}
   	return total
}

func firstSolution(data []string){
	itemValues := map[string] int {"X" : 1, "Y" : 2, "Z" : 3}
	var points[] int
	for _, line := range(data) {
		info := strings.Split(line, " ")
		oponent := info[0]
		me := info[1]
		roundPoitns := winner(oponent, me) + itemValues[me]
		points = append(points, roundPoitns)
	}
	fmt.Println(summ(points))
}

func needShape(input string, first string) string{
	gameRulesL := map[string] string {"A" : "B", "B" : "C", "C" : "A"}
	gameRulesW := map[string] string {"A" : "C", "B" : "A", "C" : "B"}
	if input == "Z"{
		return gameRulesL[first]
	}
	if input == "X"{
		return gameRulesW[first]
	}
	return first
}

func secondSolution(data []string){
	itemValues := map[string] int {"A" : 1, "B" : 2, "C" : 3}
	gamePoints := map[string] int {"X" : 0, "Y" : 3, "Z" : 6}
	var points[] int
	for _, line := range(data) {
		info := strings.Split(line, " ")
		oponent := info[0]
		need := info[1]
		me := needShape(need, oponent)
		fmt.Println("Oponent : ", oponent, "Me :", me, " Play :",need)
		roundPoitns := gamePoints[need] + itemValues[me]
		fmt.Println("RP : ", roundPoitns)
		points = append(points, roundPoitns)
	}
	fmt.Println(summ(points))
}

func main() {
	lines := ReadFile("test.txt")
	firstSolution(lines)
	secondSolution(lines)
}