package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type instruction struct {
	move int
	from int
	to   int
}

type stack []int

func main() {
	var (
		partOneAnswer string
	)
	f, err := os.Open("./inputs/data.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	partOneAnswer = partOne(f)
	fmt.Printf("The Answer to part one is: %v\n", partOneAnswer)

}

func partOne(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	stacks := createStacks(scanner)
	// instructions := createInstructions(scanner)

	fmt.Println(stacks)
	fmt.Println("---------------------")
	// fmt.Println(instructions)

	return "CMZ"
}

func createStacks(scanner *bufio.Scanner) [][]string {
	var stacks [][]string
	x := make([][]string, 50)
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range line {
			reLetter := regexp.MustCompile("[a-zA-Z]+")
			if reLetter.Match([]byte(string(v))) {
				x[i] = append(x[i], string(v))

			}
		}
		if len(line) == 0 {
			break
		}
	}
	for _, v := range x {
		if v != nil {
			stacks = append(stacks, v)
		}
	}

	return stacks
}
func remove(slice [][]string, s int) [][]string {
	if s > len(slice)-1 {
		return slice
	}
	return append(slice[:s], slice[s+1:]...)
}

func createInstructions(scanner *bufio.Scanner) []instruction {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return []instruction{}

}
