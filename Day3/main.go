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


func getPriority() map[string] int{
	var priorityList = make(map[string]int)
	i := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		priorityList[string(ch)] = i
		i++
    }
	for ch := 'A'; ch <= 'Z'; ch++ {
		priorityList[string(ch)] = i
		i++
	}
	return priorityList
}

func getRepeat(data string) string{
	for i:=0; i < len(data)/2; i++{
		for j:=len(data)/2; j < len(data); j++{
			if data[i] == data[j]{
				return string(data[i])
			}
		}
	}
	return ":"
}

 func summ(arr []int) int {
	total := 0 
	for _, a := range arr {
	  total += a
	}
	return total
}

func solutionOne(data []string){
	priority := getPriority()
	var items []int
	for _, line := range(data){
		repeated := priority[getRepeat(line)]
		items = append(items, repeated)
	}
	fmt.Println(summ(items))
}

func getRepeatGroups(one string, two string, tre string) string{
	for i:=0; i < len(one); i++{
		for j:=0; j < len(two); j++{
			for k:=0; k < len(tre); k++{
				if one[i] == two[j] && two[j] == tre[k]{
					return string(one[i])
				}
			}
		}
	}
	return ":"
}

func solutionTwo(data []string){
	priority := getPriority()
	fmt.Println(getRepeatGroups(data[0], data[1], data[2]))
	var items []int
	for i:=0; i<len(data); i+=3{
		repeated := priority[getRepeatGroups(data[i], data[i+1], data[i+2])]
		items = append(items, repeated)
	}
	fmt.Println(summ(items))
	

}

func main() {
	lines := ReadFile("test.txt")
	solutionOne(lines)
	solutionTwo(lines)
}