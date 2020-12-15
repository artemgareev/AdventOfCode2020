package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type PasswordPuzzle struct {
	min        int
	max        int
	targetChar rune
	password   string
}

func main() {
	var validCounterPart1, validCounterPart2 int

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	line := readLine(reader)
	for line != "" {
		in := strings.Split(line, " ")
		minMaxRange := strings.Split(in[0], "-")

		min, _ := strconv.Atoi(minMaxRange[0])
		max, _ := strconv.Atoi(minMaxRange[1])

		puzzle := PasswordPuzzle{
			min:        min,
			max:        max,
			targetChar: []rune(strings.TrimSuffix(in[1], ":"))[0],
			password:   in[2],
		}
		if isPasswordValidPart1(puzzle) {
			validCounterPart1++
		}
		if isPasswordValidPart2(puzzle) {
			validCounterPart2++
		}
		line = readLine(reader)
	}

	fmt.Println("Part #1:", validCounterPart1)
	fmt.Println("Part #2:", validCounterPart2)
}

// https://adventofcode.com/2020/day/2#part1
func isPasswordValidPart1(puzzle PasswordPuzzle) bool {
	repCount := 0
	for i := 0; i < len(puzzle.password); i++ {
		if puzzle.targetChar == rune(puzzle.password[i]) {
			repCount++
		}
	}
	return repCount >= puzzle.min && repCount <= puzzle.max
}

// https://adventofcode.com/2020/day/2#part2
func isPasswordValidPart2(puzzle PasswordPuzzle) bool {
	isFirstChar := rune(puzzle.password[puzzle.min-1]) == puzzle.targetChar
	isSecondChar := rune(puzzle.password[puzzle.max-1]) == puzzle.targetChar

	if (isFirstChar && !isSecondChar) || (!isFirstChar && isSecondChar) {
		return true
	}
	return false
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
