package main

import (
	"day1"
	"day10"
	"day11"
	"day12"
	"day13"
	"day2"
	"day3"
	"day4"
	"day5"
	"day6"
	"day7"
	"day8"
	"day9"
	"os"
	"utils"
)

func main() {
	functions := make(map[int]func())
	functions[1] = day1.Day1
	functions[2] = day2.Day2
	functions[3] = day3.Day3
	functions[4] = day4.Day4
	functions[5] = day5.Day5
	functions[6] = day6.Day6
	functions[7] = day7.Day7
	functions[8] = day8.Day8
	functions[9] = day9.Day9
	functions[10] = day10.Day10
	functions[11] = day11.Day11
	functions[12] = day12.Day12
	functions[13] = day13.Day13

	daysToRun := utils.ProcessArgs(functions, os.Args)

	for _, day := range daysToRun {
		functions[day]()
	}
}
