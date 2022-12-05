package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	upper, lower int
}

func main() {
	f, err := os.Open("./inputs/data.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	completelyOverlappingSections := partOne(f)

	ff, err := os.Open("./inputs/data.txt")
	defer ff.Close()
	if err != nil {
		panic(err)
	}
	overlappingSections := partTwo(ff)

	fmt.Printf("Total completely overlapping sections: %v\n", completelyOverlappingSections)
	fmt.Printf("Total overlapping sections: %v\n", overlappingSections)
}

func parseSections(sections string) (section, error) {
	var section section
	var err error

	s := strings.Split(sections, "-")
	section.lower, err = strconv.Atoi(s[0])
	section.upper, err = strconv.Atoi(s[1])
	if err != nil {
		return section, err
	}
	return section, nil
}

func isOverlappingCompletely(s1, s2 section) bool {
	if s2.lower >= s1.lower && s2.upper <= s1.upper {
		return true
	} else if s1.lower >= s2.lower && s1.upper <= s2.upper {
		return true
	}
	return false
}

func isOverlapping(s1, s2 section) bool {
	if s2.lower <= s1.upper && s2.upper >= s1.lower {
		return true
	} else if s1.lower <= s2.upper && s1.lower >= s2.upper {
		return true
	}
	return false
}

func partOne(f *os.File) int {
	var totalOverlappingSections int
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		s1, err := parseSections(assignments[0])
		s2, err := parseSections(assignments[1])
		if err != nil {
			panic(err)
		}

		if isOverlappingCompletely(s1, s2) {
			totalOverlappingSections += 1
		}
	}
	return totalOverlappingSections
}

func partTwo(f *os.File) int {
	var overlappingSections int
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		s1, err := parseSections(assignments[0])
		s2, err := parseSections(assignments[1])
		if err != nil {
			panic(err)
		}

		if isOverlappingCompletely(s1, s2) || isOverlapping(s1, s2) {
			overlappingSections += 1
		}
	}
	return overlappingSections
}
