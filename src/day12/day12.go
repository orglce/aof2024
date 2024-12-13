package day12

import (
	"utils"
)

func Day12() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func DFS(grid [][]rune, row, col int, visited *[][]bool, connectedComponent *[][2]int, gridNeighbors *[][][2]int, gridDirections *[]int) {
	(*visited)[row][col] = true
	*connectedComponent = append(*connectedComponent, [2]int{row, col})
	neighbors := getCachedNeighbors(grid, row, col, gridNeighbors)
	for _, neighbor := range neighbors {
		newRow, newCol := neighbor[0], neighbor[1]
		if !(*visited)[newRow][newCol] {
			DFS(grid, newRow, newCol, visited, connectedComponent, gridNeighbors, gridDirections)
		}
	}

}

func CheckElement(grid [][]rune, row, col, rows int, cols int, currentElement rune) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == currentElement
}

func GetNeighbors(grid [][]rune, row, col int, rowLen, colLen int) ([][2]int, int) {
	neighbors := [][2]int{}
	neighborDirections := 0
	currentElement := grid[row][col]

	for i, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newRow, newCol := row+d[0], col+d[1]
		if CheckElement(grid, newRow, newCol, rowLen, colLen, currentElement) {
			neighbors = append(neighbors, [2]int{newRow, newCol})
			neighborDirections += (i + 1) * (i + 1)
		}
	}
	return neighbors, neighborDirections
}

func GetCachedNeighborsAndDirections(grid [][]rune, row, col int, gridNeighbors *[][][2]int, gridDirections *[]int) ([][2]int, int) {
	neighbors := (*gridNeighbors)[Index(len(grid), row, col)]
	directions := (*gridDirections)[Index(len(grid), row, col)]
	return neighbors, directions
}

func getCachedNeighbors(grid [][]rune, row, col int, gridNeighbors *[][][2]int) [][2]int {
	neighbors := (*gridNeighbors)[Index(len(grid), row, col)]
	return neighbors
}

func Index(gridSize int, i, j int) int {
	return i*gridSize + j
}

func Part1(fileContent string) int {
	grid := utils.GetRuneGrid(fileContent)
	gridSize := len(grid)
	gridNeighbors := make([][][2]int, len(grid)*len(grid))
	gridDirections := make([]int, len(grid)*len(grid))
	rowLen, colLen := len(grid), len(grid[0])

	visited := make([][]bool, len(grid))
	for i, line := range grid {
		visited[i] = make([]bool, len(line))
		for j, _ := range line {
			neighbors, directions := GetNeighbors(grid, i, j, rowLen, colLen)
			gridNeighbors[Index(gridSize, i, j)] = neighbors
			gridDirections[Index(gridSize, i, j)] = directions
		}
	}

	connectedComponents := make([][][2]int, 0)
	for i, line := range grid {
		for j, _ := range line {
			if !visited[i][j] {
				connectedComponent := make([][2]int, 0)
				DFS(grid, i, j, &visited, &connectedComponent, &gridNeighbors, &gridDirections)
				connectedComponents = append(connectedComponents, connectedComponent)
			}
		}
	}

	sum := 0
	for _, connectedComponent := range connectedComponents {
		sumComponent := 0
		for _, element := range connectedComponent {
			neighbors := getCachedNeighbors(grid, element[0], element[1], &gridNeighbors)
			sumComponent += 4 - len(neighbors)
		}
		sum += sumComponent * len(connectedComponent)
	}
	return sum
}

func Part2(fileContent string) int {
	grid := utils.GetRuneGrid(fileContent)
	gridSize := len(grid)
	gridNeighbors := make([][][2]int, len(grid)*len(grid))
	gridDirections := make([]int, len(grid)*len(grid))
	rowLen, colLen := len(grid), len(grid[0])

	visited := make([][]bool, len(grid))
	for i, line := range grid {
		visited[i] = make([]bool, len(line))
		for j, _ := range line {
			neighbors, directions := GetNeighbors(grid, i, j, rowLen, colLen)
			gridNeighbors[Index(gridSize, i, j)] = neighbors
			gridDirections[Index(gridSize, i, j)] = directions
		}
	}

	connectedComponents := make([][][2]int, 0)
	for i, line := range grid {
		for j, _ := range line {
			if !visited[i][j] {
				connectedComponent := make([][2]int, 0)
				DFS(grid, i, j, &visited, &connectedComponent, &gridNeighbors, &gridDirections)
				connectedComponents = append(connectedComponents, connectedComponent)
			}
		}
	}

	sum := 0
	for _, connectedComponent := range connectedComponents {
		sumComponent := 0
		currentConnectedComponentString := grid[connectedComponent[0][0]][connectedComponent[0][1]]
		for _, element := range connectedComponent {
			neighbors, directions := GetCachedNeighborsAndDirections(grid, element[0], element[1], &gridNeighbors, &gridDirections)
			numOfNeighbors := len(neighbors)
			if numOfNeighbors == 0 {
				sumComponent += 4
			} else if numOfNeighbors == 1 {
				sumComponent += 2
			} else if numOfNeighbors == 2 {
				if directions == 5 || directions == 25 {
					sumComponent += 0
				} else {
					sumComponent += 1
				}
			}
			if CheckElement(grid, element[0], element[1]+1, len(grid), len(grid[0]), currentConnectedComponentString) {
				if CheckElement(grid, element[0]-1, element[1]+1, len(grid), len(grid[0]), currentConnectedComponentString) &&
					!CheckElement(grid, element[0]-1, element[1], len(grid), len(grid[0]), currentConnectedComponentString) {
					sumComponent += 1
				}
				if CheckElement(grid, element[0]+1, element[1]+1, len(grid), len(grid[0]), currentConnectedComponentString) &&
					!CheckElement(grid, element[0]+1, element[1], len(grid), len(grid[0]), currentConnectedComponentString) {
					sumComponent += 1
				}
				if CheckElement(grid, element[0]+1, element[1], len(grid), len(grid[0]), currentConnectedComponentString) &&
					!CheckElement(grid, element[0]+1, element[1]+1, len(grid), len(grid[0]), currentConnectedComponentString) {
					sumComponent += 1
				}
			} else {
				if CheckElement(grid, element[0]+1, element[1]+1, len(grid), len(grid[0]), currentConnectedComponentString) &&
					CheckElement(grid, element[0]+1, element[1], len(grid), len(grid[0]), currentConnectedComponentString) {
					sumComponent += 1
				}

			}
		}
		sum += sumComponent * len(connectedComponent)
	}
	return sum
}
