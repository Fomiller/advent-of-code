package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type hand struct {
	hand  string
	score int
}

var (
	rock     = hand{hand: "rock", score: 1}
	paper    = hand{hand: "paper", score: 2}
	scissors = hand{hand: "scissors", score: 3}
)

func main() {
	var totalScore int

	f, err := os.Open("./inputs/data.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), "")
		totalScore += playGame(round)
	}

	fmt.Printf("Your total score is: %v\n", totalScore)
}

func playGame(round []string) int {
	var score int

	opponent := getHand(round[0])
	player := getHand(round[2])

	if player == opponent {
		score += 3
	} else if win(player, opponent) {
		score += 6
	} else {
		score += 0
	}
	score += player.score

	return score
}

func win(p hand, o hand) bool {
	if p == rock && o == scissors || p == scissors && o == paper || p == paper && o == rock {
		return true
	} else {
		return false
	}
}

func getHand(hand string) hand {
	if hand == "A" || hand == "X" {
		return rock
	} else if hand == "B" || hand == "Y" {
		return paper
	} else {
		return scissors
	}
}
