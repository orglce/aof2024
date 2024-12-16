package day15

import (
	"strings"
	"utils"
)

func GetDirection(instruction string) (int, int) {
	switch instruction {
	case "<":
		return 0, -1
	case ">":
		return 0, 1
	case "^":
		return -1, 0
	case "v":
		return 1, 0
	}
	return 0, 0
}

func Day15() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func CalculateSum(grid [][]rune, gridLen int) int {
	sum := 0
	for i := 0; i < gridLen; i++ {
		for j := 0; j < gridLen; j++ {
			if grid[i][j] == 'O' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func Part1(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	gridLen := len(lines[0])
	instructionLines := lines[gridLen+1:]
	grid := make([][]rune, len(lines[:gridLen]))

	currentX := -1
	currentY := -1
	for i, line := range lines[:gridLen] {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
			if char == '@' {
				currentX = j
				currentY = i
			}
		}
	}

	for _, instructionLine := range instructionLines {
		for _, instruction := range strings.Split(instructionLine, "") {
			dirY, dirX := GetDirection(instruction)
			nextField := grid[currentY+dirY][currentX+dirX]
			if nextField == '#' {
				continue
			} else if nextField == '.' {
				grid[currentY][currentX] = '.'
				currentX += dirX
				currentY += dirY
				grid[currentY][currentX] = '@'
			} else if nextField == 'O' {
				startingX := currentX
				startingY := currentY
				currentX += dirX
				currentY += dirY
				tileToMove := []int{currentY, currentX}
				for {
					nextField = grid[currentY][currentX]
					if nextField == '#' {
						currentX = startingX
						currentY = startingY
						break
					} else if nextField == '.' {
						tile := tileToMove
						grid[tile[0]+dirY][tile[1]+dirX] = grid[tile[0]][tile[1]]
						grid[startingY][startingX] = '.'
						grid[startingY+dirY][startingX+dirX] = '@'
						currentX = startingX + dirX
						currentY = startingY + dirY
						break
					} else {
						tileToMove = []int{currentY, currentX}
					}
					currentX += dirX
					currentY += dirY
				}
			}
		}
	}
	return CalculateSum(grid, gridLen)
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	gridLen := len(lines[0])
	grid := make([][]rune, len(lines[:gridLen]))

	currentX := -1
	currentY := -1
	for i, line := range lines[:gridLen] {
		grid[i] = make([]rune, len(line)*2)
		for j, char := range line {
			if char == '#' {
				grid[i][j*2] = '#'
				grid[i][j*2+1] = '#'
			} else if char == 'O' {
				grid[i][j*2] = '['
				grid[i][j*2+1] = ']'
			} else if char == '.' {
				grid[i][j*2] = '.'
				grid[i][j*2+1] = '.'
			} else {
				grid[i][j*2] = char
				grid[i][j*2+1] = '.'
				currentX = j * 2
				currentY = i
			}
		}
	}

	for _, instructionLine := range lines[gridLen+1:] {
		for _, instruction := range strings.Split(instructionLine, "") {
			dirY, dirX := GetDirection(instruction)
			nextField := grid[currentY+dirY][currentX+dirX]
			if nextField == '#' {
				continue
			} else if nextField == '.' {
				grid[currentY][currentX] = '.'
				currentX += dirX
				currentY += dirY
				grid[currentY][currentX] = '@'
			} else if (instruction == "<" || instruction == ">") && (nextField == '[' || nextField == ']') {
				startingX := currentX
				startingY := currentY
				currentX += dirX
				currentY += dirY
				tilesToMove := make([][]int, 0)
				tilesToMove = append(tilesToMove, []int{currentY, currentX})
				for {
					nextField = grid[currentY][currentX]
					if nextField == '#' {
						currentX = startingX
						currentY = startingY
						break
					} else if nextField == '.' {
						for t := len(tilesToMove) - 1; t >= 0; t-- {
							tile := tilesToMove[t]
							grid[tile[0]+dirY][tile[1]+dirX] = grid[tile[0]][tile[1]]
						}
						grid[startingY][startingX] = '.'
						grid[startingY+dirY][startingX+dirX] = '@'
						currentX = startingX + dirX
						currentY = startingY + dirY
						break
					} else {
						tilesToMove = append(tilesToMove, []int{currentY, currentX})
					}
					currentX += dirX
					currentY += dirY
				}
			} else if (instruction == "^" || instruction == "v") && (nextField == '[' || nextField == ']') {
				startingX := currentX
				startingY := currentY
				tilesToMove := make(map[[2]int][2]rune, 0)
				tilesToMove[[2]int{currentY, currentX}] = [2]rune{'@', '.'}
				frontier := make(map[[2]int]bool, 0)
				frontier[[2]int{currentY, currentX}] = true
				for {
					newFrontier := make(map[[2]int]bool, 0)
					isBlocked := false
					for tile := range frontier {
						currentX, currentY = tile[1]+dirX, tile[0]+dirY
						nextField = grid[currentY][currentX]
						if nextField == '#' {
							currentX = startingX
							currentY = startingY
							isBlocked = true
							break
						} else if nextField == '[' {
							tilesToMove[[2]int{currentY, currentX}] = [2]rune{grid[currentY][currentX], grid[currentY-dirY][currentX-dirX]}
							newChar := grid[currentY-dirY][currentX-dirX+1]
							if _, ok := tilesToMove[[2]int{currentY, currentX + 1}]; !ok {
								newChar = '.'
							}
							tilesToMove[[2]int{currentY, currentX + 1}] = [2]rune{grid[currentY][currentX+1], newChar}
							newFrontier[[2]int{currentY, currentX}] = true
							newFrontier[[2]int{currentY, currentX + 1}] = true
						} else if nextField == ']' {
							tilesToMove[[2]int{currentY, currentX}] = [2]rune{grid[currentY][currentX], grid[currentY-dirY][currentX-dirX]}
							newChar := grid[currentY-dirY][currentX-dirX-1]
							if _, ok := tilesToMove[[2]int{currentY, currentX - 1}]; !ok {
								newChar = '.'
							}
							tilesToMove[[2]int{currentY, currentX - 1}] = [2]rune{grid[currentY][currentX-1], newChar}
							newFrontier[[2]int{currentY, currentX}] = true
							newFrontier[[2]int{currentY, currentX - 1}] = true
						}
					}
					if isBlocked {
						break
					} else if len(newFrontier) == 0 {
						for k, v := range tilesToMove {
							grid[k[0]+dirY][k[1]+dirX] = v[0]
							grid[k[0]][k[1]] = v[1]
						}
						currentX = startingX + dirX
						currentY = startingY + dirY
						break
					}
					frontier = newFrontier
				}
			}
		}
	}
	sum := 0
	for i := 0; i < gridLen; i++ {
		for j := 0; j < gridLen*2; j++ {
			if grid[i][j] == '[' {
				sum += 100*i + j
			}
		}
	}
	return sum
}
