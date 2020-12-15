package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	columns := make([][]rune, 0)

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	line := readLine(reader)
	for line != "" {
		in := strings.Split(line, "")

		column := make([]rune, 0)
		for i := 0; i < len(in); i++ {
			column = append(column, []rune(in[i])[0])
		}
		columns = append(columns, column)
		line = readLine(reader)
	}
	fmt.Println("Part #1", getTreeNumber(columns, 3, 1))
	fmt.Println("Part #2", getTreeNumber(columns, 1, 1)*
		getTreeNumber(columns, 3, 1)*
		getTreeNumber(columns, 5, 1)*
		getTreeNumber(columns, 7, 1)*
		getTreeNumber(columns, 1, 2))
}

const treeSign rune = '#'

func getTreeNumber(treeMap [][]rune, stepRight, stepDown int) int {
	maxRightStep := len(treeMap[0]) - 1
	var x, y, treesNumber = stepRight, stepDown, 0

	for y < len(treeMap) {
		currX := getShiftedX(maxRightStep, x)
		if treeMap[y][currX] == treeSign {
			treesNumber++
		}
		x += stepRight
		y += stepDown
	}

	return treesNumber
}

func getShiftedX(maxX int, currentX int) int {
	for currentX > maxX {
		currentX = currentX - maxX - 1
	}

	return currentX
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
