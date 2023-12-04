package main

import (
	re "regexp"
	"strconv"
	s "strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var gameRe = re.MustCompile(`(Game \d{1,3})`)
var redRe = re.MustCompile(`(\d+ red)`)
var greenRe = re.MustCompile(`(\d+ green)`)
var blueRe = re.MustCompile(`(\d+ blue)`)

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
	sumIds := 0
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sumPowers := 0
		for _, row := range inputSlice {
			games := s.Split(s.Split(row, ": ")[1], "; ")
			minR, minG, minB := 0, 0, 0
			for _, game := range games {
				r, _ := strconv.Atoi(s.Split(redRe.FindString(game), " ")[0])
				g, _ := strconv.Atoi(s.Split(greenRe.FindString(game), " ")[0])
				b, _ := strconv.Atoi(s.Split(blueRe.FindString(game), " ")[0])

				minR = max(r, minR)
				minG = max(g, minG)
				minB = max(b, minB)
			}
			sumPowers += (minR * minG * minB)
		}
		return sumPowers
	}
	// solve part 1 here
	bag := [3]int{12, 13, 14}
	for _, row := range inputSlice {
		id, _ := strconv.Atoi(s.Split(gameRe.FindString(row), " ")[1])
		games := s.Split(s.Split(row, ": ")[1], "; ")

		isValid := true
		for _, game := range games {
			r, _ := strconv.Atoi(s.Split(redRe.FindString(game), " ")[0])
			g, _ := strconv.Atoi(s.Split(greenRe.FindString(game), " ")[0])
			b, _ := strconv.Atoi(s.Split(blueRe.FindString(game), " ")[0])

			if r > bag[0] || g > bag[1] || b > bag[2] {
				isValid = false
				break
			}
		}

		if isValid {
			sumIds += id
		}
	}
	return sumIds
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
