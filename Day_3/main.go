package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	memory, err := os.Open("input.txt")

	if(err != nil){
		panic(err)
	}

	defer memory.Close()

	var totalMemory string

	scanner := bufio.NewScanner(memory)
	for scanner.Scan(){
		line := scanner.Text()
		totalMemory = totalMemory + line
	}

	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRegex.FindAllStringSubmatch(totalMemory, -1)

	total := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		
		total += x * y
	}

	fmt.Println("First Challenge: ", total)

	sentenceRegex := regexp.MustCompile(`(?:do\(\)|^)(?:[^d]|d[^o]|do[^n]|don[^']|don'[^t]|don't[^(\)])*mul\((\d{1,3}),(\d{1,3})\)`)

	total = 0

	sentences := sentenceRegex.FindAllString(totalMemory, -1)

	for _, sentence := range sentences {
		mulMatches := mulRegex.FindAllStringSubmatch(sentence, -1)

		for _, match := range mulMatches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += x * y
		}
	}

	fmt.Println("Second Challenge: ", total)
}