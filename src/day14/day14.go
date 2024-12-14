package day14

import (
	"strconv"
	"strings"
	"utils"
)

type vector struct {
	x      int
	y      int
	speedX int
	speedY int
}

func Day14() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func (v vector) Move(sizeX int, sizeY int, n int) vector {
	newX := (v.x + v.speedX*n) % sizeX
	newY := (v.y + v.speedY*n) % sizeY
	if newX < 0 {
		newX = sizeX + newX
	}
	if newY < 0 {
		newY = sizeY + newY
	}
	newVector := vector{newX, newY, v.speedX, v.speedY}
	return newVector
}

func GridRobots(robots []vector, sizeX int, sizeY int) ([][]int, bool) {
	overlap := false
	grid := make([][]int, sizeY)
	for i := 0; i < sizeY; i++ {
		grid[i] = make([]int, sizeX)
	}
	for _, robot := range robots {
		grid[robot.y][robot.x] += 1
		if grid[robot.y][robot.x] > 1 {
			overlap = true
		}
	}
	return grid, overlap
}

func Part1(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	sizeY := 103
	sizeX := 101
	robots := make([]vector, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		positions := strings.Split(strings.Replace(parts[0], "p=", "", -1), ",")
		speeds := strings.Split(strings.Replace(parts[1], "v=", "", -1), ",")
		posX := utils.Must(strconv.Atoi(positions[0]))
		posY := utils.Must(strconv.Atoi(positions[1]))
		speedX := utils.Must(strconv.Atoi(speeds[0]))
		speedY := utils.Must(strconv.Atoi(speeds[1]))
		vector := vector{posX, posY, speedX, speedY}
		robots[i] = vector
	}
	for i, robot := range robots {
		robots[i] = robot.Move(sizeX, sizeY, 100)
	}

	gridRobots, _ := GridRobots(robots, sizeX, sizeY)
	firstQuadrant := 0
	secondQuadrant := 0
	thirdQuadrant := 0
	fourthQuadrant := 0
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			if gridRobots[i][j] != 0 {
				currentNum := gridRobots[i][j]
				if i < sizeY/2 && j < sizeX/2 {
					firstQuadrant += currentNum
				} else if i < sizeY/2 && j >= sizeX/2+1 {
					secondQuadrant += currentNum
				} else if i >= sizeY/2+1 && j < sizeX/2 {
					thirdQuadrant += currentNum
				} else if i >= sizeY/2+1 && j >= sizeX/2+1 {
					fourthQuadrant += currentNum
				}
			}
		}
	}
	return firstQuadrant * secondQuadrant * thirdQuadrant * fourthQuadrant
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	sum := 0
	for j := 0; j <= 10000; j++ {
		sizeY := 103
		sizeX := 101
		robots := make([]vector, len(lines))
		for i, line := range lines {
			parts := strings.Split(line, " ")
			positions := strings.Split(strings.Replace(parts[0], "p=", "", -1), ",")
			speeds := strings.Split(strings.Replace(parts[1], "v=", "", -1), ",")
			posX := utils.Must(strconv.Atoi(positions[0]))
			posY := utils.Must(strconv.Atoi(positions[1]))
			speedX := utils.Must(strconv.Atoi(speeds[0]))
			speedY := utils.Must(strconv.Atoi(speeds[1]))
			vector := vector{posX, posY, speedX, speedY}
			robots[i] = vector
		}
		for i, robot := range robots {
			robots[i] = robot.Move(sizeX, sizeY, j)
		}

		_, overlap := GridRobots(robots, sizeX, sizeY)
		if !overlap {
			sum = j
		}

	}
	return sum
}
