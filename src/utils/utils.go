package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

func CurrentTime() time.Time {
	return time.Now()
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func ExcutionTime(start time.Time) {
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func FileToString(dayNum int) string {
	f, _ := os.OpenFile("../results.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	fmt.Printf("-------------------------------------------------------------\n")
	fileName := "../day" + strconv.Itoa(dayNum) + "/input.txt"
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetStringGrid(fileContent string) [][]string {
	lines := strings.Split(fileContent, "\n")

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}
	return grid
}

func GetRuneGrid(fileContent string) [][]rune {
	lines := strings.Split(fileContent, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	return grid
}

func GetRuneGridFromLines(lines []string) [][]rune {
	// lines := strings.Split(fileContent, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}
	return grid
}

func GetIntGrid(fileContent string) [][]int {
	lines := strings.Split(fileContent, "\n")

	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			if char == '.' {
				grid[i][j] = 2
				continue
			}
			num, _ := strconv.Atoi(string(char))
			grid[i][j] = num
		}
	}
	return grid
}

func PrintStringGrid(grid [][]string) {
	for _, line := range grid {
		for _, num := range line {
			fmt.Printf("%s ", num)
		}
		fmt.Println()
	}
}

func PrintRuneGrid(grid [][]rune) {
	for _, line := range grid {
		for _, num := range line {
			fmt.Printf("%c", num)
		}
		fmt.Println()
	}
}

func PrintGrid(grid [][]int) {
	for _, line := range grid {
		for _, num := range line {
			if num == 0 {
				fmt.Printf("%s", ".")
			} else {
				fmt.Printf("%d", num)
			}
		}
		fmt.Println()
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetCurrentDay() int {
	pc, _, _, _ := runtime.Caller(1)
	callerFunctionName := runtime.FuncForPC(pc).Name()
	currentDay, _ := strconv.Atoi(strings.Replace(strings.Split(callerFunctionName, ".")[0], "day", "", -1))
	return currentDay
}

func ProcessArgs(functions map[int]func(), args []string) []int {
	f, _ := os.Create("../../README.md")
	defer f.Close()

	f.WriteString(fmt.Sprintf("| %-17s | %-17s | %-17s |\n", "Function", "Result", "Execution time"))
	f.WriteString(fmt.Sprintf("| ----------------- | ----------------- | ----------------- |\n"))

	keys := make([]int, 0, len(functions))
	for k := range functions {
		keys = append(keys, k)
	}
	var daysToRun []int = keys
	if len(os.Args) > 1 {
		param := os.Args[1]
		daysToRun = make([]int, 0, len(functions))
		if strings.Contains(param, ",") {
			keysStrings := strings.Split(param, ",")
			for _, keyString := range keysStrings {
				key, _ := strconv.Atoi(keyString)
				daysToRun = append(daysToRun, key)
			}
		} else {
			key, _ := strconv.Atoi(param)
			daysToRun = []int{key}
		}
	}
	slices.Sort(daysToRun)
	return daysToRun
}

func RunFunc(funcToRun func(string) int, fileContent string, measureTime bool) int {
	f, err := os.OpenFile("../../README.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n := 1000

	if !measureTime {
		n = 1
	}

	um := color.New(color.FgGreen).SprintFunc()
	msFast := color.New(color.FgBlue).SprintFunc()
	msSlow := color.New(color.FgYellow).SprintFunc()
	s := color.New(color.FgRed).SprintFunc()

	start := time.Now()
	result := funcToRun(fileContent)
	end := time.Since(start).Microseconds()

	if result == 0 {
		n = 1
	}
	if !measureTime {
		fmt.Println("Result:", result)
		fmt.Println("Execution time:", time.Since(start))
		return result
	}

	functionName := GetFunctionName(funcToRun)

	if end > 1000000 {
		n = 1
	} else if end > 1000 {
		if n >= 500 {
			n = n / 100
		}
	}

	var sum float64 = 0
	for i := 0; i < n; i++ {
		start := time.Now()
		funcToRun(fileContent)
		sum += float64(time.Since(start).Microseconds())
	}
	avg := sum / float64(n)
	timeString := ""
	timeStringNormal := ""
	timeStringHtml := ""
	if avg > 1000 && avg < 10000 {
		timeStringNormal = fmt.Sprintf("%.2f ms", avg/float64(1000))
		timeStringHtml = fmt.Sprintf("![#678feb](https://placehold.co/10x10/678feb/678feb.png)   %s", timeStringNormal)
		timeString = msFast(timeStringNormal)
	} else if avg > 10000 && avg < 1000000 {
		timeStringNormal = fmt.Sprintf("%.2f ms", avg/float64(1000))
		timeStringHtml = fmt.Sprintf("![#f5b051](https://placehold.co/10x10/f5b051/f5b051.png)   %s", timeStringNormal)
		timeString = msSlow(timeStringNormal)
	} else if avg > 1000000 {
		timeStringNormal = fmt.Sprintf("%.2f s", avg/float64(1000000))
		timeStringHtml = fmt.Sprintf("![#f24663](https://placehold.co/10x10/f24663/f24663.png)   %s", timeStringNormal)
		timeString = s(timeStringNormal)
	} else {
		timeStringNormal = fmt.Sprintf("%.2f μs", float64(avg))
		timeStringHtml = fmt.Sprintf("%5s %s", "![#94ff6e](https://placehold.co/10x10/94ff6e/94ff6e.png)", timeStringNormal)
		timeString = um(timeStringNormal)
	}
	timePadding := 26
	timePaddingNormal := 17
	if result == 0 {
		timePadding = 17
		timeString = "N/A"
		timeStringNormal = "N/A"
		timeStringHtml = "N/A"
	}
	fmt.Printf("| %-17s | %-17d | %-*s |\n", functionName, result, timePadding, timeString)
	f.WriteString(fmt.Sprintf("| %-17s | %-17d | %-*s |\n", functionName, result, timePaddingNormal, timeStringHtml))

	return result
}

func RunWithGoroutines(wg *sync.WaitGroup, function func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		function()
	}()
}
