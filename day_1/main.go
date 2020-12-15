package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var expenses = map[int]interface{}{}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	line := readLine(reader)
	for line != "" {
		expense, _ := strconv.Atoi(line)
		expenses[expense] = struct{}{}
		line = readLine(reader)
	}
	fmt.Println("#task 1:", findMulOfTargetSumOf2(2020))
	fmt.Println("#task 2:", findMulOfTargetSumOf3(2020))
}

func findMulOfTargetSumOf3(targetSum int) int {
	for expenseVal, _ := range expenses {
		target := targetSum - expenseVal
		targetMul := findMulOfTargetSumOf2(target)
		if targetMul != 0 {
			return expenseVal * targetMul
		}
	}
	return 0
}

func findMulOfTargetSumOf2(targetSum int) int {
	for key, _ := range expenses {
		if _, ok := expenses[targetSum-key]; ok {
			return key * (targetSum - key)
		}
	}
	return 0
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
