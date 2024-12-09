package day3

import (
	"regexp"
	"strconv"
	"utils"
)

func Day3() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func Part1(fileContent string) int {
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	r2, _ := regexp.Compile(`\d+`)
	sum := 0
	matches := r.FindAllString(fileContent, -1)
	for _, match := range matches {
		nums := r2.FindAllString(match, -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}
	return sum
}

func Part2(fileContent string) int {
	r, _ := regexp.Compile(`((mul\(\d+,\d+\))|do\(\)|don't\(\))`)
	r2, _ := regexp.Compile(`\d+`)
	matches := r.FindAllString(fileContent, -1)
	enabled := true
	sum := 0
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		} else if match == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			nums := r2.FindAllString(match, -1)
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			sum += num1 * num2
		}
	}
	return sum
}
