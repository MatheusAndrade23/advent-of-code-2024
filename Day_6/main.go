package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate90degreesToRight(position [2]int) [2]int {
	x, y := position[0], position[1]

	switch {
	case x == 0 && y == 1:
		return [2]int{1, 0}
		
	case x == 1 && y == 0:
		return [2]int{0, -1}

	case x == 0 && y == -1:
		return [2]int{-1, 0}

	case x == -1 && y == 0:
		return [2]int{0, 1}
	}

	return position
}

func simulateGuard(roomMap []string, startPosition [2]int, obstruction [2]int) bool {
	direction := [2]int{-1,0}
	position := startPosition

	placesVisited := make(map[[3]int]bool)

	guardLeftTheRoom := false

	for !guardLeftTheRoom {
		newX, newY := position[0] + direction[0], position[1] + direction[1]

		if newX < 0 || newX >= len(roomMap) || newY < 0 || newY >= len(roomMap[0]){
			guardLeftTheRoom = true
			break
		}
		
		if(roomMap[newX][newY] == '#' || ([2]int{newX, newY} == obstruction)){
			direction = rotate90degreesToRight(direction)
		} else {
			position = [2]int{newX, newY}
		}

		state := [3]int{position[0], position[1], direction[0]*10 + direction[1]}

		if placesVisited[state] {
			return true 
		}

		placesVisited[state] = true
	}

	return false
}

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	var roomMap []string
	var startPosition [2]int

	for scanner.Scan() {
		line := scanner.Text()

		roomMap = append(roomMap, line)
	}

	placesChecked := make([][]bool, len(roomMap))

	for i := range placesChecked {
		placesChecked[i] = make([]bool, len(roomMap[0]))
	}

	var guardFound bool

	for i := 0; i < len(roomMap); i++ {
		for j := 0; j < len(roomMap[i]); j++ {
			if roomMap[i][j] == '^' {
				placesChecked[i][j] = true
				startPosition = [2]int{i, j}
				guardFound = true
				break
			}
		}
		if guardFound {
			break
		}
	}

	direction := [2]int{-1,0}
	position := startPosition

	guardLeftTheRoom := false

	for !guardLeftTheRoom {
		newX, newY := position[0] + direction[0], position[1] + direction[1]

		if newX < 0 || newX >= len(roomMap) || newY < 0 || newY >= len(roomMap[0]){
			guardLeftTheRoom = true
			break
		}
		
		if(roomMap[newX][newY] == '#'){
			direction = rotate90degreesToRight(direction)
		} else {
			placesChecked[newX][newY] = true
			position = [2]int{newX, newY}
		}
	}

	totalPlacesChecked := 0

	for i := 0; i < len(placesChecked); i++{
		for j := 0; j < len(placesChecked[0]); j++{
			if placesChecked[i][j] {
				totalPlacesChecked++
			}
		}
	}

	fmt.Println("First Challenge: ", totalPlacesChecked)

	totalPossibleLoops := 0

	for i := 0; i < len(roomMap); i++ {
		for j := 0; j < len(roomMap[i]); j++ {

			if roomMap[i][j] == '.' && [2]int{i, j} != startPosition {

				if simulateGuard(roomMap, startPosition, [2]int{i, j}) {
					totalPossibleLoops++
				}
			}
		}
	}

	fmt.Println("Second Challenge: ", totalPossibleLoops)
}