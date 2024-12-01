package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogDay(dayNum int) {
	log.Printf("Running day %d...", dayNum)
}

func CurrentTime() time.Time {
	return time.Now()
}

func ExcutionTime(start time.Time) {
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func FileToString(fileName string) string {
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
