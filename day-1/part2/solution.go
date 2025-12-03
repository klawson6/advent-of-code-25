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
	clickCount := 0

	// Walk through our input
	for scanner.Scan() {
		turn := scanner.Text()
		i, err := strconv.Atoi(turn[1:])
		if err != nil {
			panic(err)
		}
		quotient := i / 100
		remainder := i % 100
		clickCount += quotient
		if turn[0] == 'R' {
			if (pos + remainder) >= 100 {
				fmt.Printf("Click going from: %d to %d\n", pos, (pos+i)%100)
				clickCount++
			}
			pos = (pos + i) % 100
		} else {
			if pos == 0 {
				if (pos - remainder) <= -100 {
					fmt.Printf("Click going from: %d to %d\n", pos, (pos-i)%100)
					clickCount++
				}
			} else {
				if (pos - remainder) <= 0 {
					fmt.Printf("Click going from: %d to %d\n", pos, (pos-i)%100)
					clickCount++
				}
			}
			pos = (pos - i) % 100
		}
		if pos < 0 {
			pos = 100 + pos
		}
	}

	fmt.Printf("Position: %d\n", pos)
	fmt.Printf("Number of clicks: %d\n", clickCount)
}
