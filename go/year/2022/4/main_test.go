package main

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	want := 2

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

func testParttwo(t *testing.T) {
	want := 5

	f, err := os.Open("./inputs/test-data.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	result := partTwo(f)
	if want != result {
		t.Fatalf("want: %v does not match result: %v", want, result)
	}

}
