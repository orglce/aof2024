package day18

import (
	"io"
	"log"
	"strconv"
	"strings"
	"utils"
)

type Node struct {
	x       int
	y       int
	price   int
	visited bool
	prev    *Node
}

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func GetLowestUnvisitedNode18(nodes map[[2]int]*Node) (*Node, bool) {
	lowestPrice := 2147483647
	var lowestNode *Node
	for _, node := range nodes {
		// log.Println(node.y, node.x, node.price, node.visited)
		if !node.visited && node.price < lowestPrice && node.price != 0 {
			lowestNode = node
			lowestPrice = node.price
		}
	}
	if lowestNode == nil {
		return lowestNode, false
	}
	return lowestNode, true
}

func (node *Node) VisitNeighbors18(grid [][]rune, nodes map[[2]int]*Node) {
	neighbors := make([]*Node, 0)
	log.Println("Getting neighbors for [", node.y, node.x, "]", string(grid[node.y][node.x]), "Price", node.price)
	prevDirection := [2]int{0, -1}
	node.visited = true
	if node.prev != nil {
		prevDirection = [2]int{node.prev.y - node.y, node.prev.x - node.x}
	}
	for _, direction := range directions {
		nextNode, ok := nodes[[2]int{node.y + direction[0], node.x + direction[1]}]
		if !ok {
			continue
		}
		if node.prev != nil && (nextNode.x == node.prev.x && nextNode.y == node.prev.y) {
			continue
		}
		price := 0
		if (prevDirection[0] == 0 && direction[0] == 0) || (prevDirection[0] != 0) && (direction[0] != 0) {
			price = node.price + 1
		} else if (prevDirection[0] == 0 && direction[1] == 0) || (prevDirection[1] == 0 && direction[0] == 0) {
			price = node.price + 1
		}

		if nextNode.price == 0 || price < nextNode.price {
			nextNode.price = price
			nextNode.prev = node
		}
		log.Println("Neighbor", nextNode.y, nextNode.x, string(grid[nextNode.y][nextNode.x]), "Price", nextNode.price, "PrevDir", prevDirection, "Visited", nextNode.visited)
		neighbors = append(neighbors, nextNode)

	}
}

func Day18() {
	measureTime := true
	log.SetOutput(io.Discard)
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func Part1(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	gridSize := 70 + 1
	grid := make([][]rune, gridSize+2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, gridSize+2)
		if i == 0 || i == len(grid)-1 {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = '#'
			}
		} else {
			for j := 0; j < len(grid[i]); j++ {
				if j == 0 || j == len(grid)-1 {
					grid[i][j] = '#'
				} else {
					grid[i][j] = '.'
				}
			}
		}
	}
	for i := 0; i < 1024; i++ {
		nums := strings.Split(lines[i], ",")
		firstNum, _ := strconv.Atoi(nums[0])
		secondNum, _ := strconv.Atoi(nums[1])
		grid[secondNum+1][firstNum+1] = '#'
	}
	grid[1][1] = 'S'
	grid[len(grid)-2][len(grid)-2] = 'E'
	// utils.PrintRuneGrid(grid)
	// return 0
	var startingNode *Node
	var finalNode *Node
	nodes := make(map[[2]int]*Node, 0)
	// nodesQueue := make([]*Node, 0)
	// grid = make([][]rune, len(lines))
	for i, line := range grid {
		// grid[i] = make([]rune, len(line))
		for j, char := range line {
			// grid[i][j] = char
			if char == '#' {
				continue
			}
			currentNode := &Node{x: j, y: i, prev: startingNode}
			if char == 'S' {
				startingNode = currentNode
			} else if char == 'E' {
				finalNode = currentNode
			}
			nodes[[2]int{i, j}] = currentNode

		}
	}

	// grid := utils.GetRuneGrid(fileContent)
	// prev := Node{x: 1, y: 13, prev: nil}
	current := startingNode
	// i := 0
	for {
		// i += 1
		// if i > 10 {
		// 	break
		// }
		current.VisitNeighbors18(grid, nodes)

		// if nodes[[2]int{finalNode.y, finalNode.x}].visited {
		// 	break
		// }
		// for key, node := range nodes {
		// 	log.Println("\t", key, *node)
		// }
		val, ok := GetLowestUnvisitedNode18(nodes)
		if !ok {
			break
		}
		current = val
		log.Println("---------")
	}
	// sum := finalNode.price

	sum := 0
	for finalNode.prev != nil {
		log.Println(finalNode.prev.x, finalNode.prev.y)
		grid[finalNode.prev.y][finalNode.prev.x] = 'O'
		finalNode = finalNode.prev
		sum += 1
	}

	// utils.PrintRuneGrid(grid)
	return sum
}

func Part2(fileContent string) int {
	lines := strings.Split(fileContent, "\n")
	gridSize := 70 + 1
	grid := make([][]rune, gridSize+2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, gridSize+2)
		if i == 0 || i == len(grid)-1 {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = '#'
			}
		} else {
			for j := 0; j < len(grid[i]); j++ {
				if j == 0 || j == len(grid)-1 {
					grid[i][j] = '#'
				} else {
					grid[i][j] = '.'
				}
			}
		}
	}

	grid[1][1] = 'S'
	grid[len(grid)-2][len(grid)-2] = 'E'
	// utils.PrintRuneGrid(grid)
	foundAt := 0

	lowerBound := 1
	upperBound := len(lines)
	t := 0
	// for k := 0; k < len(lines); k++ {
	for {
		t += 1
		if t > 10 {
			break
		}
		// fmt.Println("lowerBound", lowerBound, "upperBound", upperBound)
		currentCounter := (upperBound + lowerBound) / 2
		notFound := 0
		for o := 0; o < 2; o++ {
			for j := 0; j < currentCounter; j++ {
				nums := strings.Split(lines[j], ",")
				firstNum, _ := strconv.Atoi(nums[0])
				secondNum, _ := strconv.Atoi(nums[1])
				grid[secondNum+1][firstNum+1] = '#'
			}
			var startingNode *Node
			var finalNode *Node
			nodes := make(map[[2]int]*Node, 0)
			for i, line := range grid {
				for j, char := range line {
					if char == '#' {
						continue
					}
					currentNode := &Node{x: j, y: i, prev: startingNode}
					if char == 'S' {
						startingNode = currentNode
					} else if char == 'E' {
						finalNode = currentNode
					}
					nodes[[2]int{i, j}] = currentNode
				}
			}
			current := startingNode
			for {

				current.VisitNeighbors18(grid, nodes)

				val, ok := GetLowestUnvisitedNode18(nodes)
				if !ok {
					break
				}
				current = val
				log.Println("---------")
			}

			sum := 0
			for finalNode.prev != nil {
				// log.Println(finalNode.prev.x, finalNode.prev.y)
				grid[finalNode.prev.y][finalNode.prev.x] = 'O'
				finalNode = finalNode.prev
				sum += 1
			}
			// fmt.Println(currentCounter, sum)

			if sum == 1 {
				notFound += 1
			}
			// utils.PrintRuneGrid(grid)

			for i := 1; i < len(grid)-1; i++ {
				for j := 1; j < len(grid[i])-1; j++ {
					grid[i][j] = '.'
				}
			}
			grid[1][1] = 'S'
			grid[len(grid)-2][len(grid)-2] = 'E'
		}
		if notFound == 0 {
			lowerBound = currentCounter + 1
		} else if notFound == 2 {
			upperBound = currentCounter
		}
		// fmt.Println("Current counter", currentCounter, notFound)

		if notFound == 1 {
			foundAt = upperBound
			break
		}
		foundAt = upperBound
	}

	// fmt.Println("At iteration", foundAt)
	// utils.PrintRuneGrid(grid)
	return foundAt
}
