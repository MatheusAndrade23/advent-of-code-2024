package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func buildPrecedenceMap(rules [][]int) map[int]map[int]bool {
	precedence := make(map[int]map[int]bool)
	for _, rule := range rules {
		a, b := rule[0], rule[1]
		if precedence[a] == nil {
			precedence[a] = make(map[int]bool)
		}
		precedence[a][b] = true
	}
	return precedence
}

func compareOrderOfNumbers(precedence map[int]map[int]bool, a, b int) bool {
	if precedence[a] != nil && precedence[a][b] {
		return true
	}
	if precedence[b] != nil && precedence[b][a] {
		return false
	}
	return false
}

func checkRowOrder(rules [][]int, row []int) bool {
	precedence := buildPrecedenceMap(rules)
	for i := 0; i < len(row); i++ {
		for j := i + 1; j < len(row); j++ {
			a := row[i]
			b := row[j]

			if !compareOrderOfNumbers(precedence, a, b) {
				return false
			}
		}
	}
	return true
}

func fixRowOrder(rules [][]int, row []int) []int {

	precedence := buildPrecedenceMap(rules)

	less := func(a, b int) bool {
		if precedence[a] != nil && precedence[a][b] {
			return true
		}

		if precedence[b] != nil && precedence[b][a] {
			return false
		}

		return false
	}


	sort.SliceStable(row, func(i, j int) bool {
		return less(row[i], row[j])
	})

	return row
}

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	var rules [][]int
	var pages [][]int

	inputSeparator := false

	for scanner.Scan() {
		line := scanner.Text()
		
		if(line == ""){
			inputSeparator = true
			continue
		}

		if(!inputSeparator){
			parts := strings.Split(line, "|")
			left, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				panic(err)
			}
			rules = append(rules, []int{left, right})
		} else {

			fields := strings.Split(line, ",")
			var row []int

			for _, field := range fields {
				num, err := strconv.Atoi(strings.TrimSpace(field))
				if err != nil {
					panic(err)
				}
				row = append(row, num)
			}
			if len(row) > 0 {
				pages = append(pages, row)
			}
		}
	}

	total:= 0
	secondTotal := 0

	for _, page := range pages {
		if checkRowOrder(rules, page) {
			middleIndex := (len(page) - 1) / 2
			total += page[middleIndex]
		} else {
			fixedPage := fixRowOrder(rules, page)
			middleIndex := (len(fixedPage) - 1) / 2
			secondTotal += fixedPage[middleIndex]
		}
	}

	fmt.Println("First Challenge: ", total)
	fmt.Println("Second Challenge: ", secondTotal)
}