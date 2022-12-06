package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func getStackId(idLine string) [][]int {
	var res [][]int
	for x, ch := range idLine {
		if unicode.IsDigit(ch) {
			number, _ := strconv.Atoi(string(ch))
			temp := [2]int{number, x}
			res = append(res, temp[:])
		}
	}
	return res
}

func getStackData(index int, cero int, data []string) []string {
	var numbers []string
	for x, line := range data {
		if x < cero-1 {
			if unicode.IsLetter(rune(line[index])) {
				number := string(line[index])
				numbers = append(numbers, number)
			}
		}
	}
	return numbers
}

type Stack struct {
	stackArray []interface{}
}

func (stack *Stack) init() *Stack {
	stack.stackArray = make([]interface{}, 0)
	return stack
}

func NewStack() *Stack {
	return new(Stack).init()
}

func (stack *Stack) Push(data interface{}) {
	stack.stackArray = append(stack.stackArray, data)
}

func (a *Stack) Add(b *Stack, n int) {
	for i := b.Length() - 1; i > b.Length()-n-1; i-- {
		a.Push(b.stackArray[i])
	}
	for i := 0; i < n; i++ {
		b.Pop()
	}
}

func (a *Stack) AddTwo(b *Stack, n int) {
	for i := b.Length() - n; i < b.Length(); i++ {
		a.Push(b.stackArray[i])
	}
	for i := 0; i < n; i++ {
		b.Pop()
	}
}

func (stack *Stack) Pop() {
	stack.stackArray = stack.stackArray[:len(stack.stackArray)-1]
}

func (stack *Stack) Reset() {
	stack.stackArray = stack.stackArray[0:0]
}

func (stack *Stack) Top() interface{} {
	return stack.stackArray[len(stack.stackArray)-1]
}

func (stack *Stack) Length() int {
	return len(stack.stackArray)
}

func (stack *Stack) Print() {
	for i := len(stack.stackArray) - 1; i >= 0; i-- {
		fmt.Println(stack.stackArray[i])
	}
	fmt.Println("=====")
}

func makeStacks(data []string) map[int]*Stack {
	var stacks = make(map[int]*Stack)
	cero := 0
	for x, line := range data {
		if len(line) == 0 {
			cero = x
		}
	}
	stackId := getStackId(data[cero-1])
	for _, pair := range stackId {
		data := getStackData(pair[1], cero, data)
		stack := NewStack()
		for i := len(data) - 1; i >= 0; i-- {
			stack.Push(data[i])
		}
		stacks[pair[0]] = stack
	}
	return stacks
}

func getInstructions(data []string) [][]int {
	var instructions [][]int
	cero := 0
	for x, line := range data {
		if len(line) == 0 {
			cero = x
		}
	}
	for x, line := range data {
		if x > cero {
			splited := strings.Split(line, " ")
			quantity, _ := strconv.Atoi(splited[1])
			from, _ := strconv.Atoi(splited[3])
			to, _ := strconv.Atoi(splited[5])
			instructions = append(instructions, []int{quantity, from, to})
		}
	}
	return instructions
}

func solutionOne(data []string) {
	stacks := makeStacks(data)
	instructions := getInstructions(data)
	for _, instruction := range instructions {
		stacks[instruction[2]].Add(stacks[instruction[1]], instruction[0])
		for _, stack := range stacks {
			stack.Print()
		}
	}
}

func solutionTwo(data []string) {
	stacks := makeStacks(data)
	instructions := getInstructions(data)
	for _, instruction := range instructions {
		stacks[instruction[2]].AddTwo(stacks[instruction[1]], instruction[0])
		for _, stack := range stacks {
			stack.Print()
		}
		fmt.Println("Instructions : ", instruction)
	}
	for x, stack := range stacks {
		fmt.Println(x)
		fmt.Println(stack.Top())
	}
}

func main() {
	lines := ReadFile("test.txt")
	solutionTwo(lines)

}
