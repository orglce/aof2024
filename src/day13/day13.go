package day13

import (
	"math"
	"strconv"
	"strings"
	"utils"

	"gonum.org/v1/gonum/mat"
)

func Day13() {
	measureTime := true
	fileContent := utils.FileToString(utils.GetCurrentDay())

	utils.RunFunc(Part1, fileContent, measureTime)
	utils.RunFunc(Part2, fileContent, measureTime)
}

func FloatToIntIfClose(value float64, epsilon float64) int {
	rounded := math.Round(value)
	if math.Abs(value-rounded) <= epsilon {
		return int(rounded)
	}
	return -1
}

func Part1(fileContent string) int {
	sum := 0
	lines := strings.Split(fileContent, "\n")
	for i := 0; i < (len(lines)+1)/4; i++ {
		line1 := lines[i*4]
		line2 := lines[i*4+1]
		line3 := lines[i*4+2]

		nums := strings.Split(strings.Split(line1, ": ")[1], ", ")
		numA1 := strings.Replace(nums[0], "X", "", -1)
		numA2 := strings.Replace(nums[1], "Y", "", -1)
		num1Int, _ := strconv.Atoi(numA1)
		num2Int, _ := strconv.Atoi(numA2)
		nums = strings.Split(strings.Split(line2, ": ")[1], ", ")
		numB1 := strings.Replace(nums[0], "X", "", -1)
		numB2 := strings.Replace(nums[1], "Y", "", -1)
		num3Int, _ := strconv.Atoi(numB1)
		num4Int, _ := strconv.Atoi(numB2)
		nums = strings.Split(strings.Split(line3, ": ")[1], ", ")
		numC1 := strings.Replace(nums[0], "X=", "", -1)
		numC2 := strings.Replace(nums[1], "Y=", "", -1)
		num5Int, _ := strconv.Atoi(numC1)
		num6Int, _ := strconv.Atoi(numC2)

		matrix := mat.NewDense(2, 2, []float64{
			float64(num1Int), float64(num3Int),
			float64(num2Int), float64(num4Int),
		})

		vector := mat.NewVecDense(2, []float64{float64(num5Int), float64(num6Int)})

		inverse := mat.NewDense(2, 2, nil)
		inverse.Inverse(matrix)

		result := mat.NewVecDense(2, nil)
		result.MulVec(inverse, vector)

		first := FloatToIntIfClose(result.At(0, 0), 0.001)
		second := FloatToIntIfClose(result.At(1, 0), 0.001)

		if first == -1 || second == -1 {
			continue
		}
		sum += first*3 + second
	}

	return sum
}

func Part2(fileContent string) int {

	sum := 0
	lines := strings.Split(fileContent, "\n")
	for i := 0; i < (len(lines)+1)/4; i++ {
		line1 := lines[i*4]
		line2 := lines[i*4+1]
		line3 := lines[i*4+2]

		nums := strings.Split(strings.Split(line1, ": ")[1], ", ")
		numA1 := strings.Replace(nums[0], "X", "", -1)
		numA2 := strings.Replace(nums[1], "Y", "", -1)
		num1Int, _ := strconv.Atoi(numA1)
		num2Int, _ := strconv.Atoi(numA2)
		nums = strings.Split(strings.Split(line2, ": ")[1], ", ")
		numB1 := strings.Replace(nums[0], "X", "", -1)
		numB2 := strings.Replace(nums[1], "Y", "", -1)
		num3Int, _ := strconv.Atoi(numB1)
		num4Int, _ := strconv.Atoi(numB2)
		nums = strings.Split(strings.Split(line3, ": ")[1], ", ")
		numC1 := strings.Replace(nums[0], "X=", "", -1)
		numC2 := strings.Replace(nums[1], "Y=", "", -1)
		num5Int, _ := strconv.Atoi(numC1)
		num6Int, _ := strconv.Atoi(numC2)

		num5Int = num5Int + 10000000000000
		num6Int = num6Int + 10000000000000

		matrix := mat.NewDense(2, 2, []float64{
			float64(num1Int), float64(num3Int),
			float64(num2Int), float64(num4Int),
		})

		vector := mat.NewVecDense(2, []float64{float64(num5Int), float64(num6Int)})

		inverse := mat.NewDense(2, 2, nil)
		inverse.Inverse(matrix)

		result := mat.NewVecDense(2, nil)
		result.MulVec(inverse, vector)

		first := FloatToIntIfClose(result.At(0, 0), 0.001)
		second := FloatToIntIfClose(result.At(1, 0), 0.001)

		if first == -1 || second == -1 {
			continue
		}
		sum += first*3 + second
	}

	return sum
}
