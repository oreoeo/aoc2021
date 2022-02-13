package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Window struct {
	val1 int
	val2 int
	val3 int
	sum  int
}

var prevDepth int = 0
var currentDepth int = 0

func main() {
	file, err := os.Open("./sonar-report.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// partOne(file)
	// partTwo(file)
}

func partOne(file *os.File) {
	scans := loadScan(file)

	counter := 0
	prevDepth := scans[0]
	currentDepth := scans[1]

	for i := 1; i < len(scans); i++ {
		currentDepth = scans[i]

		if currentDepth > prevDepth {
			counter++
		}

		prevDepth = currentDepth
	}

	fmt.Println("Part One Answer: ")
	fmt.Println(counter)
}

func partTwo(file *os.File) {
	scans := loadScan(file)
	var windows []Window

	for i := 0; i < len(scans); i++ {
		if i+1 > len(scans)-1 || i+2 > len(scans)-1 {
			break
		} else {
			windows = append(windows, Window{val1: scans[i], val2: scans[i+1], val3: scans[i+2], sum: scans[i] + scans[i+1] + scans[i+2]})
		}
	}

	counter := 0
	prevWindow := windows[0]
	currentWindow := windows[1]

	for i := 1; i < len(windows); i++ {
		currentWindow = windows[i]

		if currentWindow.sum > prevWindow.sum {
			counter++
		}

		prevWindow = currentWindow
	}

	fmt.Println("Part Two Answer:")
	fmt.Println(counter)
}

func loadScan(file *os.File) []int {
	var result []int

	buff := bufio.NewScanner(file)

	for buff.Scan() {
		val, e := strconv.Atoi(buff.Text())

		if e == nil {
			result = append(result, val)
		} else {
			log.Fatal(e)
		}
	}

	return result
}
