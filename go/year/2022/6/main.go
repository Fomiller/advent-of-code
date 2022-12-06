package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	packetMarker  = 4
	messageMarker = 14
)

func main() {
	f, err := os.Open("./inputs/data.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	partOneAnswer := partOne(f)
	fmt.Printf("The Answer to part one is: %v\n", partOneAnswer)

	f.Seek(0, 0)

	partTwoAnswer := partTwo(f)
	fmt.Printf("The Answer to part Two is: %v\n", partTwoAnswer)
}

func partOne(f *os.File) int {
	return CommunicationSystem(f, packetMarker)
}

func partTwo(f *os.File) int {
	return CommunicationSystem(f, messageMarker)
}

func CommunicationSystem(f *os.File, marker int) int {
	r := bufio.NewReader(f)
	reader := bufio.Reader(*r)

	//count of starting byte index
	count := 0
	for {
		duplicateChars := false
		uniqueChars := 0
		// look at next 4 bytes in string
		b, err := reader.Peek(marker)
		if err != nil {
			break
		}
		// loop over []byte to check for unique chars
		for _, v := range string(b) {
			// if we found any duplicates break out of loop
			if duplicateChars == true {
				break
			}
			// check []bytes for duplicate chars
			if strings.Count(string(b), string(v)) == 1 {
				uniqueChars++
			} else {
				duplicateChars = true
			}
		}
		// if we found 4 uniqueChars we are done
		if uniqueChars == marker {
			break
		} else {
			reader.Read(b[:1])
			count += 1
		}
	}
	return count + marker
}
