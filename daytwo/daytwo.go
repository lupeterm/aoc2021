package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
}

func getLines() ([]string, []int) {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var direction []string
	var steps []int
	for scanner.Scan() {
		tuple := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(tuple[1])
		check(err)
		steps = append(steps, val)
		direction = append(direction, tuple[0])
	}
	file.Close()
	return direction, steps
}

func partone(dir []string, steps []int) {
	width := 0
	depth := 0
	for i := 0; i < len(steps); i++ {
		switch dir[i] {
		case "forward":
			width += steps[i]
		case "down":
			depth += steps[i]
		case "up":
			depth -= steps[i]

		}
	}
	fmt.Printf("part one: %d\n", width*depth)
}

func parttwo(dir []string, steps []int) {
	width := 0
	depth := 0
	aim := 0
	for i := 0; i < len(steps); i++ {
		switch dir[i] {
		case "forward":
			width += steps[i]
			depth += aim * steps[i]
		case "down":
			aim += steps[i]
		case "up":
			aim -= steps[i]

		}
	}

	fmt.Printf("part two: %d\n", width*depth)
}

func main() {
	directions, steps := getLines()
	partone(directions, steps)
	parttwo(directions, steps)
}
