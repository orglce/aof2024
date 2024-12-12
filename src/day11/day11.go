package day11

import (
	"math"
	"strconv"
	"strings"
	"utils"
)

func Day11() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func ProcessStone(num int, n int, maxN int, results map[int]int) int {
	if n >= maxN {
		return 1
	}
	result := 0
	key := CantorPair(num, n)
	if val, exists := results[key]; exists {
		return val
	}
	switch {
	case num == 0:
		result = ProcessStone(2024, n+2, maxN, results)
	case (int(math.Log10(float64(num))+1.0))%2 == 0:
		num1, num2 := SplitInt(num)
		first := ProcessStone(num1, n+1, maxN, results)
		second := ProcessStone(num2, n+1, maxN, results)
		result = first + second
	default:
		result = ProcessStone(num*2024, n+1, maxN, results)
	}
	results[key] = result
	return result
}

func SplitInt(num int) (int, int) {
	Base := 10
	divisor := Base
	for num/divisor > divisor {
		divisor *= Base
	}
	return num / divisor, num % divisor
}

func CantorPair(k1, k2 int) int {
	return (k1+k2)*(k1+k2+1)/2 + k2
}

func CalculateStones(fileContent string, maxN int) int {
	stones := strings.Split(fileContent, " ")
	stonesNums := make([]int, len(stones))
	for i, stone := range stones {
		num, _ := strconv.Atoi(stone)
		stonesNums[i] = num
	}
	numOfStones := 0
	results := make(map[int]int)
	for i := 0; i < len(stones); i++ {
		numOfStones += ProcessStone(stonesNums[i], 0, maxN, results)
	}
	return numOfStones
}

func Part1(fileContent string) int {
	return CalculateStones(fileContent, 25)
}

func Part2(fileContent string) int {
	return CalculateStones(fileContent, 75)
}
