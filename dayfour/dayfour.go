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

func getLines() []string {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func helper(s string) []int {
	var arr []int
	ss := strings.Fields(s)
	for _, val := range ss {
		a, e := strconv.Atoi(val)
		check(e)
		arr = append(arr, a)
	}
	return arr
}

func getBoards(lines []string) [][][]int {
	var boards [][][]int
	for i := 2; i < len(lines); i += 6 {
		sublines := lines[i : i+5]
		var b [][]int
		for _, l := range sublines {
			b = append(b, helper(l))
		}
		boards = append(boards, b)
	}
	return boards
}

func sumslice(slice []int) int {
	var sum int
	for _, e := range slice {
		sum += e
	}
	return sum
}

func gotfive(b [][]int) bool {
	for i := 0; i < len(b); i++ {
		if sumslice(b[i]) == -5 {
			return true
		}
	}
	for row := 0; row < len(b); row++ {
		subsum := 0
		for col := 0; col < len(b); col++ {
			subsum += b[col][row]
		}
		if subsum == -5 {
			return true
		}
	}
	return false
}

func getsum(b [][]int) int {
	var sum int
	for _, v := range b {
		for _, x := range v {
			if x != -1 {
				sum += x
			}
		}
	}
	return sum
}

func game(b [][][]int, inputs []int, lines []string) {
	order := make([]int, len(b))
	winners := 0
	for _, num := range inputs {
		for i, board := range b {
			for j, row := range board {
				for k, n := range row {
					if num == n {
						b[i][j][k] = -1
					}
				}
			}
			if gotfive(board) {
				if order[i] != -1 {
					winners++
					order[i] = -1
					if winners == 1 {
						fmt.Printf("Day 4 Part one: %d\n", getsum(board)*num)
					}
				}
				if winners == len(b) {
					fmt.Printf("Day 4 Part two: %d\n", getsum(board)*num)
					return
				}
			}
		}
	}
}

func main() {
	lines := getLines()
	first := lines[0]
	vals := strings.Split(first, ",")
	var input []int
	for _, val := range vals {
		a, e := strconv.Atoi(val)
		check(e)
		input = append(input, a)
	}
	boards := getBoards(lines)
	// printboards(boards)
	game(boards, input, lines)
}
