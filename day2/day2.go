package day2

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func Day2() {
	utils.LogDay(2)

	fileContent := utils.FileToString("../day2/input.txt")

	start := utils.CurrentTime()
	fmt.Println("Part1: ", Part1(fileContent))
	utils.ExcutionTime(start)

	start = utils.CurrentTime()
	fmt.Println("Part2: ", Part2(fileContent))
	utils.ExcutionTime(start)
}

func Part1(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	notSafeLines := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		rising := val1-val2 < 0
		for i := 0; i < len(parts)-1; i++ {
			if !AreLevelsSafe(parts[i], parts[i+1], rising) {
				notSafeLines += 1
				break
			}
		}
	}
	return len(lines) - notSafeLines
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	safeLines := 0

	for _, line := range lines {
		isLineSafe := make([]bool, 2)

		// traverse forward
		isReportSafe := true
		parts := strings.Split(line, " ")
		val1, _ := strconv.Atoi(parts[0])
		val2, _ := strconv.Atoi(parts[1])
		rising := val1-val2 < 0
		problems := 0
		for i := 0; i < len(parts)-1; i++ {
			isSafe := AreLevelsSafe(parts[i], parts[i+1], rising)
			if !isSafe && problems == 0 {
				if i+2 >= len(parts) {
					break
				}
				isNextSafe := AreLevelsSafe(parts[i], parts[i+2], rising)
				if !isNextSafe {
					isReportSafe = false
					break
				}
				problems += 1
				i += 1
			} else if !isSafe && problems == 1 {
				isReportSafe = false
				break
			}
		}
		if isReportSafe {
			isLineSafe[0] = true
		}

		// traverse backward
		isReportSafe = true
		parts = strings.Split(line, " ")
		val1, _ = strconv.Atoi(parts[len(parts)-1])
		val2, _ = strconv.Atoi(parts[len(parts)-2])
		rising = val1-val2 < 0
		problems = 0
		for i := len(parts) - 1; i >= 1; i-- {
			isSafe := AreLevelsSafe(parts[i], parts[i-1], rising)
			if !isSafe && problems == 0 {
				if i-2 < 0 {
					break
				}
				isNextSafe := AreLevelsSafe(parts[i], parts[i-2], rising)
				if !isNextSafe {
					isReportSafe = false
					break
				}
				problems += 1
				i -= 1
			} else if !isSafe && problems == 1 {
				isReportSafe = false
				break
			}
		}
		if isReportSafe {
			isLineSafe[1] = true
		}

		if isLineSafe[0] || isLineSafe[1] {
			safeLines += 1
		}
	}
	return safeLines
}

func AreLevelsSafe(level1 string, level2 string, rising bool) bool {
	val1, _ := strconv.Atoi(level1)
	val2, _ := strconv.Atoi(level2)
	diff := val1 - val2
	if rising != (diff < 0) || utils.Abs(diff) > 3 || utils.Abs(diff) < 1 {
		return false
	}
	return true
}
