package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var commands []string

	file, err := os.Open("./day2-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buff := bufio.NewScanner(file)

	for buff.Scan() {
		commands = append(commands, buff.Text())
	}

	pilotSub(commands)
}

func pilotSub(commands []string) {
	hPos := 0
	vPos := 0
	aim := 0 // part 2

	for i := 0; i < len(commands); i++ {
		command := strings.Fields(commands[i])
		direction, magnitude := command[0], command[1]
		magVal, err := strconv.Atoi(magnitude)

		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			hPos += magVal
			vPos += (aim * magVal) // part 2
		case "up":
			aim -= magVal
		case "down":
			aim += magVal
		default:
			log.Fatal("Error parsing command")
		}
	}

	fmt.Printf("hPos: %d, vPos: %d, answer: %d", hPos, vPos, hPos*vPos)
}
