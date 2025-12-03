package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get input from command line arg
	if len(os.Args) < 2 {
		panic("No input file provided.")
	}
	filename := os.Args[1]

	// Open file stream
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Init counters
	pos := 50
	zeroCount := 0

	// Walk through our input
	for scanner.Scan() {
		turn := scanner.Text()
		i, err := strconv.Atoi(turn[1:])
		if err != nil {
			panic(err)
		}
		if turn[0] == 'R' {
			pos = (pos + i) % 100
		} else {
			pos = (pos - i) % 100
		}
		if pos < 0 {
			pos = 100 + pos
		} else if pos == 0 {
			zeroCount++
		}
	}

	fmt.Printf("Position: %d\n", pos)
	fmt.Printf("Number of zeros: %d\n", zeroCount)
}
