package day9

import (
	"math"
	"strconv"
	"utils"
)

func Day9() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

type file struct {
	id       int
	index    int
	length   int
	isCopied bool
}

type gap struct {
	length       int
	index        int
	freedUpSpace int
	files        []file
}

func Part1(fileContent string) int {
	numsLen := len(fileContent)
	nums := make([]int, numsLen)
	numOfFiles := int(math.Ceil(float64(numsLen) / 2))
	for i, char := range fileContent {
		num, _ := strconv.Atoi(string(char))
		nums[i] = num
	}
	remFileId := numOfFiles - 1
	remPosition := numsLen - 1
	currentDiskPosition := 0
	gapsPosition := 1
	compressedDisk := make([]int, 0)

	currentFileId := 0
	for {
		if currentFileId != 0 {
			for nums[gapsPosition] > 0 {
				compressedDisk = append(compressedDisk, remFileId)
				nums[remPosition] -= 1
				nums[gapsPosition] -= 1
				if nums[remPosition] == 0 {
					remPosition -= 2
					remFileId -= 1
				}
			}
		}
		for 0 < nums[currentDiskPosition] {
			compressedDisk = append(compressedDisk, currentFileId)
			nums[currentDiskPosition] -= 1
		}
		currentDiskPosition += 2
		if currentDiskPosition > numsLen-1 || (nums[currentDiskPosition] == 0 && gapsPosition+1 >= remPosition) {
			break
		}
		if nums[remPosition] > 0 {
			for i := 0; nums[remPosition] > 0 && i < nums[gapsPosition]; i++ {
				compressedDisk = append(compressedDisk, remFileId)
				nums[remPosition] -= 1
				nums[gapsPosition] -= 1
			}
			if nums[remPosition] == 0 {
				remPosition -= 2
				remFileId -= 1
				if remFileId == currentFileId {
					break
				}
			}
		}
		currentFileId++
		if nums[gapsPosition] == 0 {
			gapsPosition += 2
		}

	}

	sum := 0
	for i, num := range compressedDisk {
		sum += num * i
	}
	return sum
}

func CalculateChecksum(files []file, gaps []gap) int {
	sum := 0
	currentNum := 0
	for i := 0; i < len(files); i++ {
		currentFile := files[i]
		var currentGap gap
		if i < len(gaps) {
			currentGap = gaps[i]
		}
		for j := 0; j < currentFile.length; j++ {
			if !currentFile.isCopied {
				sum += currentFile.id * currentNum
				currentNum += 1
			} else {
				currentNum += 1
			}
		}
		for _, f := range currentGap.files {
			for k := 0; k < f.length; k++ {
				sum += f.id * currentNum
				currentNum += 1
			}
		}
		for k := 0; k < currentGap.length; k++ {
			currentNum += 1
		}
	}
	return sum
}

func Part2(fileContent string) int {
	files := make([]file, 0)
	gaps := make([]gap, 0)
	for i, char := range fileContent {
		num, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			files = append(files, file{index: i, id: int(i / 2), length: num, isCopied: false})
		} else {
			gaps = append(gaps, gap{index: i, length: num, freedUpSpace: 0, files: []file{}})
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(gaps); j++ {
			if gaps[j].index >= files[i].index {
				break
			}
			if !files[i].isCopied && files[i].length <= gaps[j].length {
				gaps[j].files = append(gaps[j].files, files[i])
				gaps[j].length -= files[i].length
				files[i].isCopied = true
				break
			}
		}
	}

	sum := CalculateChecksum(files, gaps)
	return sum
}
