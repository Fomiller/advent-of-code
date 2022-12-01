package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type calories int

// func toCalorie(x int) calories {
// 	return calories(x)

// }

type elf struct {
	id            int
	totalCalories calories
	snacks        []calories
	totalSnacks   int
}

func (e *elf) setTotalCalories() {
	var totalCalories calories
	for _, v := range e.snacks {
		totalCalories += v
	}
	e.totalCalories = totalCalories
}

func (e *elf) setTotalSnacks() int {
	e.totalSnacks = len(e.snacks)
	return e.totalSnacks
}

func (e *elf) getTotalCalories() calories {
	var totalCalories calories
	for _, v := range e.snacks {
		totalCalories += v
	}
	e.totalCalories = totalCalories
	return e.totalCalories
}

func (e *elf) getTotalSnacks() int {
	e.totalSnacks = len(e.snacks)
	return e.totalSnacks
}

func main() {
	var (
		count int = 1
	)

	f, err := os.Open("./inputs/data.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	elves := []elf{}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := strings.Join(strings.Split(scanner.Text(), "\n"), "")
		elf := elf{id: count}

		if s != "" {
			x, err := strconv.Atoi(s)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(elf.snacks)
			elf.snacks = append(elf.snacks, calories(x))
			fmt.Println(elf)
		} else {
			count += 1
			fmt.Println(elf)
		}

		elf.setTotalCalories()
		elf.setTotalSnacks()
		elves = append(elves, elf)
	}

	fmt.Printf("Total Elves: %v\n", len(elves))
	fmt.Println(elves[0])
	fmt.Println(elves[len(elves)-1])

}
