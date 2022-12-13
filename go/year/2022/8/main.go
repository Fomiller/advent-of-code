package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./inputs/test-data.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	partOneAnswer := partOne(f)
	fmt.Printf("The Answer to part one is: %v\n", partOneAnswer)
}

func partOne(f *os.File) int {
	input := parseInput(f)
	for _, v := range input {
		for ii, vv := range v {
			isVisibleEW(v, ii, vv)
			isVisibleNS(input, ii, ii)
		}
	}
	return 21
}

func parseInput(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines := [][]int{}

	for scanner.Scan() {
		lineStr := strings.Split(scanner.Text(), "")
		lineInt := convertSliceStringToInt(lineStr)
		lines = append(lines, lineInt)
	}
	return lines
}

func isVisibleEW(s []int, idx int, value int) bool {
	west := s[:idx]
	east := s[idx+1:]
	fmt.Printf("visible from west: %v\n", isTallestInRow(west, value))
	fmt.Printf("visible from east: %v\n", isTallestInRow(east, value))
	fmt.Println("---------------")
	return true
}

func isVisibleNS(s [][]int, idx int, value int) bool {
	north := s[:idx]
	south := s[idx+1:]
	fmt.Printf("visible from north: %v\n", isTallestInColumn(north, value))
	fmt.Printf("visible from south: %v\n", isTallestInColumn(south, value))
	fmt.Println("---------------")
	return true
}

func convertSliceStringToInt(s []string) []int {
	x := []int{}
	for _, v := range s {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		x = append(x, v)
	}
	return x
}

func isTallestInRow(s []int, t int) bool {
	isTallest := true
	for _, v := range s {
		if v >= t {
			isTallest = false
		}
	}
	return isTallest
}

func isTallestInColumn(rows [][]int, t int) bool {
	isTallest := true
	// columns := [][]int{}
	for i, v := range rows {
		for ii, _ := range v {
			fmt.Println(rows[i][ii])
		}
		fmt.Println(v[0])
	}
	// fmt.Println(columns)
	return isTallest
}
