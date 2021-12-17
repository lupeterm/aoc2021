package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partone(lines []int) {
	var asc int
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			asc++
		}
	}
	fmt.Printf("part one: %d\n", asc)
}

func sumslice(slice []int) int {
	var sum int
	for _, e := range slice {
		sum += e
	}
	return sum
}

func parttwo(lines []int) {
	var asc int
	for i := 2; i < len(lines); i++ {
		prev := sumslice(lines[i-2 : i+1])
		curr := sumslice(lines[i-1 : i+2])
		if prev < curr {
			asc++
		}
	}
	fmt.Printf("parttwo: %d\n", asc)
}

func getLines() []int {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check(err)
		lines = append(lines, val)
	}
	file.Close()
	return lines
}

func main() {
	lines := getLines()
	partone(lines)
	parttwo(lines)
}
