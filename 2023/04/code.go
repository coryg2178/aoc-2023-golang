package main

import (
	re "regexp"
	"slices"
	s "strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var numRegex = re.MustCompile(`[\d]+`)

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
	inputSlice := s.Split(input, "\n")
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	totalPoints := 0
	for _, card := range inputSlice {
		winningNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[0], -1)
		haveNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[1], -1)

		points := 0
		for _, num := range haveNums {
			if slices.Contains(winningNums, num) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		totalPoints += points
	}

	return totalPoints
}
