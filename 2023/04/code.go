package main

import (
	"math"
	re "regexp"
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
		inputArr := make([]int, len(inputSlice))
		totalPoints := 0
		for i, card := range inputSlice {
			winningNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[0], -1)
			haveNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[1], -1)

			matches := intersect(winningNums, haveNums)

			for j := i + 1; j <= i+len(matches); j++ {
				inputArr[j] += (1 + inputArr[i])
			}

			totalPoints += (1 + inputArr[i])
		}
		return totalPoints
	}
	// solve part 1 here
	totalPoints := 0
	for _, card := range inputSlice {
		winningNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[0], -1)
		haveNums := numRegex.FindAllString(s.Split(s.Split(card, ": ")[1], "| ")[1], -1)

		matches := intersect(winningNums, haveNums)

		totalPoints += int(math.Pow(float64(2), float64(len(matches)-1)))
	}

	return totalPoints
}

func intersect(a []string, b []string) []string {
	set := make([]string, 0)
	hash := make(map[string]bool)

	for _, val := range a {
		hash[val] = true
	}

	for _, val := range b {
		if hash[val] {
			set = append(set, val)
		}
	}

	return set
}
