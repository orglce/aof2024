package day1

import (
	"slices"
	"strconv"
	"strings"
	"utils"
)

func Day1() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
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
		nums[secondNum] += 1
		list1[i] = firstNum
	}

	sum := 0
	for i := 0; i < numOfLines; i++ {
		sum += list1[i] * nums[list1[i]]
	}
	return sum
}
