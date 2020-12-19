package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getSeatNumberByBoardingPass(boardingPass string) int {
	lowerBound, upperBound := 0, 1<<7
	leftBound, rightBound := 0, 1<<3

	for i := 0; i < len(boardingPass); i++ {
		upDownHalf := (upperBound - lowerBound) >> 1
		if boardingPass[i] == 'F' {
			upperBound -= upDownHalf
		}
		if boardingPass[i] == 'B' {
			lowerBound += upDownHalf
		}

		leftRightHalf := (rightBound - leftBound) >> 1
		if boardingPass[i] == 'L' {
			rightBound -= leftRightHalf
		}
		if boardingPass[i] == 'R' {
			leftBound += leftRightHalf
		}
	}

	return lowerBound*8 + leftBound
}

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	boardingPasses := strings.Split(string(bytes), "\n")

	highestPass := 0
	for _, pass := range boardingPasses {
		passSeatNumber := getSeatNumberByBoardingPass(pass)
		if passSeatNumber > highestPass {
			highestPass = passSeatNumber
		}
	}
	fmt.Println(highestPass)
}
