package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type item struct {
	item string
}

func (i item) getPriority() int {
	var p int
	for ii, v := range alphabet {
		if i.item == string(v) {
			p = ii + 1
			break
		}
	}
	return p
}

func main() {
	var totalPriority int

	f, err := os.Open("./inputs/data.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		rucksacks := strings.Split(scanner.Text(), "")
		middle := (len(rucksacks) / 2)
		sack1 := rucksacks[:middle]
		sack2 := rucksacks[middle:]

		s, err := getPriorityItem(sack1, sack2)
		if err != nil {
			panic(err)
		}

		totalPriority += s.getPriority()
	}
	fmt.Println(totalPriority)
}

func getPriorityItem(list1, list2 []string) (item, error) {
	for _, v := range list1 {
		for _, vv := range list2 {
			if v == vv {
				return item{v}, nil
			}
		}
	}
	return item{}, errors.New("no like characters")
}
