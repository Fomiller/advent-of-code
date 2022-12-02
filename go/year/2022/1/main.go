package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./inputs/data.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var elfSnacks []int
	snacks := 0

	for _, v := range lines {
		if len(v) == 0 {
			elfSnacks = append(elfSnacks, snacks)
			snacks = 0
			continue
		}
		calories, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		snacks += calories
	}
	elfSnacks = append(elfSnacks, snacks)

	elf, calorieCount := findHighestCalorieCount(elfSnacks)
	fmt.Printf("Elf number %v has the highest calorie count with %v calories.\n", elf, calorieCount)
}

func findHighestCalorieCount(counts []int) (int, int) {
	elf := 0
	highestCalories := 0

	for i, calories := range counts {
		if highestCalories < calories {
			highestCalories = calories
			elf = i
		}
	}
	return elf, highestCalories
}
