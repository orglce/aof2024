package day6

import (
	"utils"
)

func Day6() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func Part1(fileContent string) int {
	return 0
}

func Part2(fileContent string) int {
	return 0
}
