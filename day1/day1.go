package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func Day1() {
	utils.LogDay(1)

	fileContent := utils.FileToString("../day1/input.txt")

	start := utils.CurrentTime()
	fmt.Println("Part1: ", Part1(fileContent))
	utils.ExcutionTime(start)

	start = utils.CurrentTime()
	fmt.Println("Part2: ", Part2(fileContent))
	utils.ExcutionTime(start)

}

func Part1(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	numOfLines := len(lines)
	list1 := make([]int, numOfLines)
	list2 := make([]int, numOfLines)
	for i, line := range lines {
		parts := strings.Split(line, "   ")
		firstNum, _ := strconv.Atoi(parts[0])
		secondNum, _ := strconv.Atoi(parts[1])
		list1[i] = int(firstNum)
		list2[i] = int(secondNum)
	}
	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := 0; i < numOfLines; i++ {
		sum += utils.Abs(list1[i] - list2[i])
	}
	return sum
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")

	numOfLines := len(lines)
	list1 := make([]int, numOfLines)
	nums := make(map[int]int)

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		firstNum, _ := strconv.Atoi(parts[0])
		secondNum, _ := strconv.Atoi(parts[1])
		firstNum = int(firstNum)
		secondNum = int(secondNum)
		val, ok := nums[secondNum]
		if ok {
			nums[secondNum] = val + 1
		} else {
			nums[secondNum] = 1
		}
		list1[i] = firstNum
	}

	sum := 0
	for i := 0; i < numOfLines; i++ {
		sum += list1[i] * nums[list1[i]]
	}
	return sum
}
