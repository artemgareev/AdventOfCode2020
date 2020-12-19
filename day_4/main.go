package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var allowedMasks = map[uint8]interface{}{
	255: struct{}{}, // 1111 1111 | 255
	127: struct{}{}, // 0111 1111 | 255^(1<<7)
}

var passportElementMasks = map[string]uint8{
	"byr": 1 << 0, //(Birth Year)
	"iyr": 1 << 1, //(Issue Year)
	"eyr": 1 << 2, //(Expiration Year)
	"hgt": 1 << 3, //(Height)
	"hcl": 1 << 4, //(Hair Color)
	"ecl": 1 << 5, //(Eye Color)
	"pid": 1 << 6, //(Passport ID)
	"cid": 1 << 7, //(Country ID)
}

var pidReg = regexp.MustCompile(`^\d{9}$`)
var eclReg = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var hclReg = regexp.MustCompile(`^#([0-9a-f]{6})$`)
var hgtCMReg = regexp.MustCompile(`^(?:15[0-9]|1[6-8][0-9]|19[0-3])cm$`)
var hgtINReg = regexp.MustCompile(`^(?:59|6[0-9]|7[0-6])in$`)
var byrReg = regexp.MustCompile(`(?:192[0-9]|19[3-9][0-9]|200[0-2])`)
var iyrReg = regexp.MustCompile(`(?:201[0-9]|2020)`)
var eyrReg = regexp.MustCompile(`(?:202[0-9]|2030)`)

type validator func(input string) bool

var passportElementValidators = map[string]validator{
	//(Birth Year) - four digits; at least 1920 and at most 2002.
	"byr": func(input string) bool {
		return byrReg.MatchString(input)
	},
	//(Issue Year) - four digits; at least 2010 and at most 2020.
	"iyr": func(input string) bool {
		return iyrReg.MatchString(input)
	},
	//(Expiration Year) - four digits; at least 2020 and at most 2030.
	"eyr": func(input string) bool {
		return eyrReg.MatchString(input)
	},
	//(Height) - a number followed by either cm or in:
	//- If cm, the number must be at least 150 and at most 193.
	//- If in, the number must be at least 59 and at most 76.
	"hgt": func(input string) bool {
		return hgtCMReg.MatchString(input) || hgtINReg.MatchString(input)
	},
	//(Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	"hcl": func(input string) bool {
		return hclReg.MatchString(input)
	},
	//(Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	"ecl": func(input string) bool {
		return eclReg.MatchString(input)
	},
	//(Passport ID) - a nine-digit number, including leading zeroes.
	"pid": func(input string) bool {
		return pidReg.MatchString(input)
	},
	//(Country ID) - ignored, missing or not.
	"cid": func(input string) bool {
		return true
	},
}

var dataUnitReg = regexp.MustCompile(`(\w{3}:[\#\w]*)`)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	passportsData := strings.Split(string(bytes), "\n\n")

	var validPassportsNumber1, validPassportsNumber2 int
	for _, passportData := range passportsData {
		if isPassportValidPart1(passportData) {
			validPassportsNumber1++
		}
		if isPassportValidPart2(passportData) {
			validPassportsNumber2++
		}
	}
	fmt.Println("Part #1: ", validPassportsNumber1)
	fmt.Println("Part #2: ", validPassportsNumber2)
}

func isPassportValidPart1(passportRaw string) bool {
	matches := dataUnitReg.FindAllString(passportRaw, -1)

	var resultMask uint8
	for _, match := range matches {
		matchParts := strings.Split(match, ":")
		key, _ := matchParts[0], matchParts[1]

		keyMask, ok := passportElementMasks[key]
		if ok {
			resultMask |= keyMask
		}
	}
	_, ok := allowedMasks[resultMask]

	return ok
}

func isPassportValidPart2(passportRaw string) bool {
	matches := dataUnitReg.FindAllString(passportRaw, -1)

	var resultMask uint8
	for _, match := range matches {
		matchParts := strings.Split(match, ":")
		key, val := matchParts[0], matchParts[1]

		valValidator, _ := passportElementValidators[key]
		if !valValidator(val) {
			return false
		}
		keyMask, ok := passportElementMasks[key]
		if ok {
			resultMask |= keyMask
		}
	}
	_, ok := allowedMasks[resultMask]

	return ok
}
