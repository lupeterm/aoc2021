package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
}

func parttwo(most string, least string, lines []string) {
	oxygen := get_oxy(lines, 0)
	scrubber := get_scr(lines, 0)
	fmt.Println(oxygen)
	fmt.Println(scrubber)

	ox, err := strconv.ParseInt(oxygen, 2, 64)
	check(err)
	scr, err := strconv.ParseInt(scrubber, 2, 64)
	check(err)
	fmt.Printf("ogygen: %d, scrubber: %d \n", ox, scr)
	fmt.Printf("parttwo: %d", ox*scr)
}

func get_oxy(lines []string, index int) string {
	var ones, zeros []string
	if len(lines) == 1 {
		return lines[0]
	}
	for _, bits := range lines {
		if bits[index] == 49 {
			ones = append(ones, bits)
		} else {
			zeros = append(zeros, bits)
		}
	}
	if len(ones) >= len(zeros) {
		return get_oxy(ones, index+1)
	}
	return get_oxy(zeros, index+1)
}

func get_scr(lines []string, index int) string {
	var ones, zeros []string
	if len(lines) == 1 {
		return lines[0]
	}
	for _, bits := range lines {
		if bits[index] == 49 {
			ones = append(ones, bits)
		} else {
			zeros = append(zeros, bits)
		}
	}
	if len(zeros) <= len(ones) {
		return get_scr(zeros, index+1)
	}
	return get_scr(ones, index+1)
}

func partone(lines []string) (string, string) {
	var gamma, eps int64
	var length int
	coms := make([]int, len(lines[0]))
	for _, bits := range lines {
		for i, rune := range bits {
			coms[i] += int(rune % 48)
		}
	}
	length = len(coms)
	for i := length - 1; i >= 0; i-- {
		if coms[i] > 500 {
			gamma += 1 << (length - i - 1)
			// coms[i] = 1
		} else {
			eps += 1 << (length - i - 1)
			// coms[i] = 0
		}
	}
	most := strconv.FormatInt(gamma, 2)
	least := strconv.FormatInt(eps, 2)
	fmt.Println(coms)
	fmt.Println(most)
	fmt.Printf("00%s\n", least)
	fmt.Printf("partone: %d\n\n", gamma*eps)
	return most, least
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

func main() {
	lines := getLines()
	most, least := partone(lines)
	parttwo(most, least, lines)

}
