package main

import (
	"regexp"
	"strconv"
	s "strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var numRegex = regexp.MustCompile(`(\d)`)
var alphaNumRegex = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
var numReplace = s.NewReplacer("oneight", "18", "threeight", "38", "fiveight", "58", "nineight", "98", "twone", "21", "eightwo", "82")
var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	inputArr := s.Split(input, "\n")
	sumCalVals := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		for _, str := range inputArr {
			str = numReplace.Replace(str)
			numList := alphaNumRegex.FindAllString(str, -1)

			first, last := numList[0], numList[len(numList)-1]

			if val, ok := numMap[first]; ok {
				first = val
			}

			if val, ok := numMap[last]; ok {
				last = val
			}

			val, _ := strconv.Atoi(string(first) + string(last))
			sumCalVals += val
		}
		return sumCalVals
	}
	// solve part 1 here
	for _, str := range inputArr {
		substr := numRegex.FindAllString(str, -1)
		val, _ := strconv.Atoi(string(substr[0]) + string(substr[len(substr)-1]))
		sumCalVals += val
	}
	return sumCalVals
}
