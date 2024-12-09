package day8

import (
	"fmt"
	"strings"
	"utils"
)

func Day8() {
	utils.LogDay(8)

	fileContent := utils.FileToString("../day8/input.txt")

	start := utils.CurrentTime()
	fmt.Println("Part1: ", Part1(fileContent))
	utils.ExcutionTime(start)

	start = utils.CurrentTime()
	fmt.Println("Part2: ", Part2(fileContent))
	utils.ExcutionTime(start)
}

type vector struct {
	x int
	y int
}

func Points(v vector, u vector) []vector {
	diff := vector{u.x - v.x, u.y - v.y}
	points := make([]vector, 2)
	points[0] = vector{u.x + diff.x, u.y + diff.y}
	points[1] = vector{v.x - diff.x, v.y - diff.y}
	return points
}

func Points2(v vector, u vector, sizeX int, sizeY int) []vector {
	diff := vector{u.x - v.x, u.y - v.y}
	points := []vector{vector{v.x, v.y}}
	originalPoint := vector{v.x, v.y}
	currentPoint := vector{v.x, v.y}
	for currentPoint.x < sizeX && currentPoint.y < sizeY && currentPoint.x >= 0 && currentPoint.y >= 0 {
		points = append(points, currentPoint)
		currentPoint = vector{currentPoint.x + diff.x, currentPoint.y + diff.y}
	}
	currentPoint = originalPoint
	for currentPoint.x > 0 && currentPoint.y > 0 && currentPoint.x < sizeX && currentPoint.y < sizeY {
		points = append(points, currentPoint)
		currentPoint = vector{currentPoint.x - diff.x, currentPoint.y - diff.y}
	}
	return points
}

func combinations(arr []vector) [][]vector {
	combs := make([][]vector, 0)
	for _, vec1 := range arr {
		for _, vec2 := range arr {
			if vec1 == vec2 {
				continue
			}
			comb := []vector{vec1, vec2}
			combs = append(combs, comb)
		}
	}
	return combs
}

func CreateGridAndAntennas(fileContent string) ([][]string, map[string][]vector) {
	lines := strings.Split(fileContent, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
	}

	antennas := make(map[string][]vector)
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "." {
				continue
			}
			antennas[char] = append(antennas[char], vector{j, i})
		}
	}

	return grid, antennas
}

func Part1(fileContent string) int {
	grid, antennas := CreateGridAndAntennas(fileContent)

	num := 0

	for _, v := range antennas {
		combs := combinations(v)
		for _, comb := range combs {
			points := Points(comb[0], comb[1])
			for _, point := range points {
				if point.x < 0 || point.y < 0 || point.x >= len(grid) || point.y >= len(grid[0]) {
					continue
				}
				if grid[point.x][point.y] == "X" {
					continue
				} else {
					num += 1
					grid[point.x][point.y] = "X"
				}
			}
		}
	}
	return num
}

func Part2(fileContent string) int {
	grid, antennas := CreateGridAndAntennas(fileContent)

	num := 0
	for _, v := range antennas {
		combs := combinations(v)
		for _, comb := range combs {
			points := Points2(comb[0], comb[1], len(grid), len(grid[0]))
			for _, point := range points {
				if grid[point.x][point.y] == "X" {
					continue
				} else {
					num += 1
					grid[point.x][point.y] = "X"
				}
			}
		}
	}

	return num
}
