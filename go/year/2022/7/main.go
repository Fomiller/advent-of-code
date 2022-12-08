package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./inputs/data.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	partOneAnswer := partOne(f)
	fmt.Printf("The Answer to part one is: %v\n", partOneAnswer)

	// f.Seek(0, 0)

	// partTwoAnswer := partTwo(f)
	// fmt.Printf("The Answer to part Two is: %v\n", partTwoAnswer)
}

func partOne(f *os.File) int {
	commands := parseCommands(f)
	fmt.Println(commands[1].output[0])

	return 95437
}

// func partTwo() int {
// 	return 0
// }

func parseCommands(f *os.File) []command {
	count := 0
	var cmd command
	var commands []command

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "$") {
			cmd = command{}
			cmd.input = strings.Split(line, " ")[1:]
			commands = append(commands, cmd)
			count++
		} else {
			cmd.output = append(cmd.output, line)
			commands[count-1] = cmd
		}

	}
	return commands
}
