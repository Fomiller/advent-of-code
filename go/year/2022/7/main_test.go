package main

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	want := 95437
	f, err := os.Open("./inputs/test-data.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	result := partOne(f)

	if want != result {
		t.Fatalf("want: %v does not match result: %v", want, result)
	}
}

// func TestPartTwo(t *testing.T) {
// 	want := 19
// 	f, err := os.Open("./inputs/test-data-part-2.txt")
// 	defer f.Close()

// 	if err != nil {
// 		panic(err)
// 	}
// 	result := partTwo(f)

// 	if want != result {
// 		t.Fatalf("want: %v does not match result: %v", want, result)
// 	}
// }
