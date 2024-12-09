package day4

import (
	"strings"
	"utils"
)

func Day4() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func CheckWord(grid [][]string, i int, j int, word string, wordLen int) bool {
	return false
}

func CheckPoint(grid [][]string, i int, j int, word string, wordLen int, combinations [][]int) int {
	words := make([]string, 8)

	for c, combination := range combinations {
		incX := 0
		incY := 0
		currentWord := ""
		for k := 0; k < wordLen; k++ {
			currentWord += grid[i+incX][j+incY]
			incX += combination[0]
			incY += combination[1]
		}
		words[c] = currentWord
	}

	howMany := 0
	for _, foundWord := range words {
		if word == foundWord {
			howMany++
		}
	}

	return howMany
}

func Part1(fileContent string) int {
	searchWord := "XMAS"
	searchWordLen := len(searchWord)
	resizeLen := 2 * (searchWordLen - 1)
	resizeHalf := searchWordLen - 1
	lines := strings.Split(fileContent, "\n")
	lineLen := len(lines[0])

	combinations := [][]int{
		{-1, -1}, // top left
		{0, -1},  // top
		{1, -1},  // top right
		{1, 0},   // right
		{1, 1},   // bottom right
		{0, 1},   // bottom
		{-1, 1},  // bottom left
		{-1, 0},  // left
	}

	grid := make([][]string, len(lines)+resizeLen)
	for i := 0; i < len(lines)+resizeLen; i++ {
		grid[i] = make([]string, lineLen+resizeLen)
	}

	for i := resizeHalf; i < len(lines)+resizeHalf; i++ {
		line := lines[i-resizeHalf]
		for j := resizeHalf; j < lineLen+resizeHalf; j++ {
			grid[i][j] = string(line[j-resizeHalf])
		}
	}

	howMany := 0
	for i := resizeHalf; i < len(lines)+resizeHalf; i++ {
		for j := resizeHalf; j < lineLen+resizeHalf; j++ {
			if grid[i][j] == "X" {
				howMany += CheckPoint(grid, i, j, searchWord, searchWordLen, combinations)
			}
		}
	}
	return howMany
}

func Part2(fileContent string) int {
	searchWord := "MAS"
	searchWordLen := len(searchWord)
	resizeLen := 2 * (searchWordLen - 1)
	resizeHalf := searchWordLen - 1
	lines := strings.Split(fileContent, "\n")
	lineLen := len(lines[0])
	howMany := 0

	grid := make([][]string, len(lines)+resizeLen)
	for i := 0; i < len(lines)+resizeLen; i++ {
		grid[i] = make([]string, lineLen+resizeLen)
		for j := 0; j < lineLen+resizeLen; j++ {
			grid[i][j] = "0"
		}
	}

	for i := resizeHalf; i < len(lines)+resizeHalf; i++ {
		line := lines[i-resizeHalf]
		for j := resizeHalf; j < lineLen+resizeHalf; j++ {
			grid[i][j] = string(line[j-resizeHalf])
		}
	}

	topLeft := [][]int{{-1, -1}}
	topRight := [][]int{{1, -1}}
	bottomRight := [][]int{{1, 1}}
	bottomLeft := [][]int{{-1, 1}}

	for i := resizeHalf; i < len(lines)+resizeHalf; i++ {
		for j := resizeHalf; j < lineLen+resizeHalf; j++ {
			if grid[i][j] == "M" {
				if ((CheckPoint(grid, i, j, searchWord, searchWordLen, bottomRight) > 0) && (CheckPoint(grid, i+2, j, searchWord, searchWordLen, bottomLeft) > 0)) ||
					((CheckPoint(grid, i, j, searchWord, searchWordLen, bottomRight) > 0) && (CheckPoint(grid, i, j+2, searchWord, searchWordLen, topRight) > 0)) {
					howMany += 1
				}
				if ((CheckPoint(grid, i, j, searchWord, searchWordLen, topLeft) > 0) && (CheckPoint(grid, i-2, j, searchWord, searchWordLen, topRight) > 0)) ||
					((CheckPoint(grid, i, j, searchWord, searchWordLen, topLeft) > 0) && (CheckPoint(grid, i, j-2, searchWord, searchWordLen, bottomLeft) > 0)) {
					howMany += 1
				}

			}
		}
	}
	return howMany
}
