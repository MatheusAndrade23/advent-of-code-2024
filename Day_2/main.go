package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	isIncreasing, isDecreasing := true, true

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if diff > 0 {
			isDecreasing = false
		} else {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func canBeMadeSafe(levels []int) bool {

	for i := 0; i < len(levels); i++ {
		modified := append([]int{}, levels[:i]...)

		modified = append(modified, levels[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}
	return false
}

func main() {
	reports, err := os.Open("input.txt")

	if(err != nil){
		panic(err)
	}

	defer reports.Close()

	totalSafeReports := 0
	totalPossibleSafeReports := 0

	scanner := bufio.NewScanner(reports)
	for scanner.Scan(){
		line := scanner.Text()
		nums := strings.Fields(line)

		levels := make([]int, len(nums))
		for i, num := range nums {
			levels[i], err = strconv.Atoi(num)

			if(err != nil){
				panic(err)
			}
		}
	
		if isSafe(levels) {
			totalSafeReports++
		}

		if (isSafe(levels) || canBeMadeSafe(levels)) {
			totalPossibleSafeReports++
		}
	}

	fmt.Println("First Challenge: ", totalSafeReports)
	fmt.Println("Second Challenge: ", totalPossibleSafeReports)
}