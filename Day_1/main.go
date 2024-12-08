package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pairs, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer pairs.Close()

	var lefts, rights []int

	var totalDistance, similarityPoints int

	scanner := bufio.NewScanner(pairs)
	for scanner.Scan(){
		rawData := scanner.Text()
		space := strings.Index(rawData, " ")
		left := rawData[:space]
		right := rawData[space+3:]

		intLeft, err := strconv.Atoi(left)

		if(err != nil){
			panic(err)
		}

		intRight, err := strconv.Atoi(right)

		if(err != nil){
			panic(err)
		}
		
		lefts = append(lefts, intLeft)
		rights = append(rights, intRight)
	}

	sort.Ints(rights)
	sort.Ints(lefts)

	for i:= 0; i < len(lefts); i++ {
		distance := rights[i] - lefts[i]

		if distance < 0 {
			distance = distance * -1
		}

		totalDistance += distance
	}

	for i:= 0; i < len(lefts); i++ {
		times := 0
		for j:= 0; j < len(rights); j++ {
			if lefts[i] == rights[j] {
				times++
			}
		}
		similarityPoints += lefts[i] * times
	}

	fmt.Println("First Challenge: ", totalDistance)
	fmt.Println("Second Challenge: ", similarityPoints)
}