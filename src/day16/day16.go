package day16

import (
	"fmt"
	"log"
	"strings"
	"utils"
)

type Node struct {
	x     int
	y     int
	price int
	// direction [2]int
	visited bool
	prev    *Node
}

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func GetLowestUnvisitedNode(nodes map[[2]int]*Node) (*Node, bool) {
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

func pop(q *[]*Node) int {
	lowestPrice := 2147483647
	var lowestNode *Node
	for _, node := range *q {
		if node.price < lowestPrice && node.price != 0 {
			lowestNode = node
			lowestPrice = node.price
		}
	}
	*q = append((*q)[:0], (*q)[1:]...)
	return lowestNode.price
}

func (node *Node) VisitNeighbors(grid [][]rune, q *[]*Node, visited map[[4]int]bool, nodes map[[2]int]*Node) {
	neighbors := make([]*Node, 0)
	log.Println("Getting neighbors for [", node.y, node.x, "]", string(grid[node.y][node.x]), "Price", node.price)
	prevDirection := [2]int{0, -1}
	if node.prev != nil {
		prevDirection = [2]int{node.prev.y - node.y, node.prev.x - node.x}
	}
	visited[[4]int{node.y, node.x, prevDirection[0], prevDirection[1]}] = true
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
			price = node.price + 1001
		}

		if nextNode.price == 0 || price < nextNode.price {
			nextNode.price = price
			nextNode.prev = node
		}
		log.Println("Neighbor", nextNode.y, nextNode.x, string(grid[nextNode.y][nextNode.x]), "Price", nextNode.price, "PrevDir", prevDirection, "Visited", nextNode.visited)
		// if node.prev != nil {
		// 	log.Println("\tPrev", node.y, node.x)
		// }
		neighbors = append(neighbors, nextNode)

	}
	// return neighbors
}

func Day16() {
	measureTime := true
	// log.SetOutput(io.Discard)
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func Part1(fileContent string) int {
	return 0
	lines := strings.Split(fileContent, "\n")
	var startingNode *Node
	var finalNode *Node
	nodes := make(map[[2]int]*Node, 0)
	// nodesQueue := make([]*Node, 0)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
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
	// current := startingNode
	// i := 0
	for {
		// i += 1
		// if i > 10 {
		// 	break
		// }
		// current.VisitNeighbors(grid, )

		// if nodes[[2]int{finalNode.y, finalNode.x}].visited {
		// 	break
		// }
		// for key, node := range nodes {
		// 	log.Println("\t", key, *node)
		// }
		// val, ok := GetLowestUnvisitedNode(nodes)
		// if !ok {
		// break
		// }
		// current = val
		log.Println("---------")
	}
	sum := finalNode.price

	for finalNode.prev != nil {
		log.Println(finalNode.prev.x, finalNode.prev.y)
		grid[finalNode.prev.y][finalNode.prev.x] = 'O'
		finalNode = finalNode.prev
	}
	fmt.Println(nodes[[2]int{4, 2}].prev, nodes[[2]int{4, 2}].price)
	fmt.Println(nodes[[2]int{2, 4}].prev, nodes[[2]int{2, 4}].price)

	utils.PrintRuneGrid(grid)
	return sum
}

func Part2(fileContent string) int {
	return 0
}
