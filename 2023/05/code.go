package main

import (
	"math"
	re "regexp"
	"strconv"
	s "strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

type record struct {
	sourceRange [2]int
	destStart   int
}

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
	almanac := s.Split(input, "\n\n")
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	tables := make([][]record, len(almanac)-1)
	seeds := toIntArr(numRegex.FindAllString(almanac[0], -1))
	for i, table := range almanac[1:] {
		tables[i] = makeConverter(s.Split(table, "\n")[1:])
	}

	lowest := math.MaxInt
	for _, location := range seeds {

		for _, table := range tables {
			location = runConversion(table, location)
		}

		if location < lowest {
			lowest = location
		}
	}
	// solve part 1 here
	return lowest
}

func toIntArr(strs []string) []int {
	out := make([]int, len(strs))

	for i, val := range strs {
		n, _ := strconv.Atoi(val)
		out[i] = n
	}

	return out
}

func makeConverter(s []string) []record {
	m := make([]record, len(s))

	for i, row := range s {
		x := toIntArr(numRegex.FindAllString(row, -1))
		m[i] = record{[2]int{x[1], x[1] + (x[2] - 1)}, x[0]}
	}

	return m
}

func runConversion(t []record, s int) int {
	for _, r := range t {
		sr := r.sourceRange
		d := r.destStart
		if s >= sr[0] && s <= sr[1] {
			return d + (s - sr[0])
		}
	}
	return s
}
