package day10

import (
	"utils"
)

func Day10() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func ScanForStep(i int, j int, grid [][]int, foundTrails map[[2]int]int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	currentStep := grid[i][j]

	// down up right left
	nextSteps := [][]int{{1, 0, -1}, {-1, 0, -1}, {0, 1, -1}, {0, -1, -1}}

	if i < len(grid)-1 {
		nextSteps[0][2] = grid[i+1][j]
	}
	if i > 0 {
		nextSteps[1][2] = grid[i-1][j]
	}
	if j < len(grid[0])-1 {
		nextSteps[2][2] = grid[i][j+1]
	}
	if j > 0 {
		nextSteps[3][2] = grid[i][j-1]
	}

	for _, nextStep := range nextSteps {
		if currentStep == 8 && nextStep[2] == 9 {
			finalPoint := [2]int{i + nextStep[0], j + nextStep[1]}
			foundTrails[finalPoint] += 1
		}
		if nextStep[2]-1 == currentStep {
			ScanForStep(i+nextStep[0], j+nextStep[1], grid, foundTrails)
		} else {
			continue
		}
	}
}

func Part1(fileContent string) int {
	grid := utils.GetIntGrid(fileContent)

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				foundTrails := make(map[[2]int]int)
				ScanForStep(i, j, grid, foundTrails)
				sum += len(foundTrails)
			}
		}
	}

	return sum
}

func Part2(fileContent string) int {
	grid := utils.GetIntGrid(fileContent)

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				foundTrails := make(map[[2]int]int)
				ScanForStep(i, j, grid, foundTrails)
				for _, v := range foundTrails {
					if v > 0 {
						sum += v
					}
				}
			}
		}
	}

	return sum
}
