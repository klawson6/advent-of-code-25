package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := strings.Split(getInput(), ",")
	result := make(chan uint64)
	// Evaluate each range
	for _, rng := range ranges {
		go evalRange(rng, result)
	}

	total := uint64(0)
	for i := 0; i < len(ranges); i++ {
		total += <-result
	}

	// Print results
	fmt.Printf("Total of invalid IDs: %d\n", total)
}

func getInput() string {
	// Get input from command line arg
	if len(os.Args) < 2 {
		panic("No input file provided.")
	}
	filename := os.Args[1]

	// Read whole file to string
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func evalRange(rng string, result chan uint64) {
	total := uint64(0)
	fromTo := strings.Split(rng, "-")
	if len(fromTo) != 2 {
		panic("A range in the test data has an invalid format")
	}

	from, err := strconv.ParseInt(fromTo[0], 10, 64)
	if err != nil {
		panic(err)
	}
	to, err := strconv.ParseInt(fromTo[1], 10, 64)
	if err != nil {
		panic(err)
	}

	for i := from; i <= to; i++ {
		currentNum := uint64(i)
		currentStr := strconv.FormatUint(currentNum, 10)
		if evalStr(currentStr) {
			total += currentNum
		}
	}
	result <- total
}

func evalStr(str string) bool {
	length := len(str)
	for chunkLength := 1; chunkLength <= length/2; chunkLength++ {
		if length%chunkLength == 0 {
			if walkStr(str, chunkLength) {
				return true
			}
		}
	}
	return false
}

func walkStr(str string, chunkLength int) bool {
	for j := 0; j < len(str)-chunkLength; j = j + chunkLength {
		if str[j:j+chunkLength] != str[(j+chunkLength):(j+2*chunkLength)] {
			return false
		}
	}
	fmt.Printf("Invalid: %s\n", str)
	return true
}
