package day7

import (
	"strconv"
	"strings"
	"utils"
)

func Day7() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func concatenate(x, y int) int {
	pow := int(10)
	for y >= pow {
		pow *= 10
	}
	return x*pow + y
}

func Part1(fileContent string) int {
	return 0
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testVal, _ := strconv.Atoi(parts[0])
		nums := strings.Split(parts[1], " ")
		numsInt := make([]int, len(nums))
		for i, num := range nums {
			numsInt[i], _ = strconv.Atoi(num)
		}
		isValid := Calc(1, numsInt, 0, testVal, 0)
		if isValid {
			sum += testVal
		}
	}
	return sum
}

func Calc(sum int, nums []int, i int, res int, op int) bool {
	// 0 = *
	// 1 = +
	// 2 = concat
	newRes := 0
	if i == len(nums) {
		return sum == res
	}
	if op == 0 {
		newRes = sum * nums[i]
	} else if op == 1 {
		newRes = sum + nums[i]
	} else if op == 2 {
		newRes = concatenate(sum, nums[i])
	}

	if newRes > res {
		return false
	}
	return Calc(newRes, nums, i+1, res, 0) || Calc(newRes, nums, i+1, res, 1) || Calc(newRes, nums, i+1, res, 2)
}
