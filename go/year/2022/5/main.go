package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	instructions := createInstructions(scanner)

	newStacks := crane(stacks, instructions)
	fmt.Println(newStacks)
	topCrates := getTopCrates(newStacks)

	return topCrates
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

func createInstructions(scanner *bufio.Scanner) []instruction {
	var instructions []instruction
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("[0-9]+")
		i := re.FindAllString(line, -1)
		if i == nil {
			continue
		}

		x := instruction{}
		x.move, _ = strconv.Atoi(i[0])
		x.from, _ = strconv.Atoi(i[1])
		x.to, _ = strconv.Atoi(i[2])
		instructions = append(instructions, x)
	}
	return instructions
}

func crane(stacks [][]string, instructions []instruction) [][]string {
	fmt.Println(instructions)
	for _, v := range instructions {
		moves := 0
		fmt.Printf("m: %v\n", v.move)
		for moves < v.move {
			// fmt.Println(len(stacks))
			// fmt.Println(v.from)
			fromStack := stacks[v.from-1]
			fmt.Printf("f: %v\n", fromStack)
			crate, fromStack := fromStack[0], fromStack[1:]
			fmt.Printf("c: %v\n", crate)
			fmt.Printf("f: %v\n", fromStack)

			toStack := stacks[v.to-1]
			fmt.Printf("t: %v\n", toStack)
			toStack = append([]string{crate}, toStack...)
			fmt.Printf("t: %v\n", toStack)

			stacks[v.from-1] = fromStack
			stacks[v.to-1] = toStack
			fmt.Printf("s: %s\n", stacks)

			fmt.Println("-------------")
			moves++
		}
	}

	return stacks
}

func getTopCrates(stacks [][]string) string {
	var output string
	for i := range stacks {
		output += stacks[i][0]
	}
	return output
}
