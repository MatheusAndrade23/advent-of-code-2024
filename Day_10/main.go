package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func findHeadTrails (topographicMap [][]int) []Point {
	headTrails := make([]Point, 0)

	for i := 0; i < len(topographicMap); i++ {
		for j := 0; j < len(topographicMap[0]); j ++ {
			
			if topographicMap[i][j] == 0 {
				point := Point{
					x: i,
					y: j,
				}

				headTrails = append(headTrails, point)
			}
		}
	}

	return headTrails
}

func followTrail(point Point, topographicMap [][]int, directions []Point, peaksFound [][]bool) (score int, rating int) {

	headTrailScore := 0
	headTrailRating := 0

	currentHigh := topographicMap[point.x][point.y]

	if currentHigh == 9 {
		
		if !peaksFound[point.x][point.y] {
			peaksFound[point.x][point.y] = true

			return 1, 1
		}	

		return 0, 1
	}
	

	for _, direction := range directions {

		newX := point.x + direction.x
		newY := point.y + direction.y

		newPoint := Point{
			x: newX,
			y: newY,
		}

		if newX >= len(topographicMap) || newY >= len(topographicMap[0]) || newY < 0 || newX < 0 {
			continue
		}

		if topographicMap[newX][newY] == (currentHigh + 1) {
			score, rating := followTrail(newPoint, topographicMap, directions, peaksFound)

			headTrailScore += score
			headTrailRating += rating
		}
	}

	return headTrailScore, headTrailRating
}

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	topographicMap := make([][]int, 0)

	directions := []Point{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}

	for scanner.Scan() {
		line := scanner.Text()
		
		intLine := make([]int, 0)

		for _, byte := range line {
			number, _ := strconv.Atoi(string(byte))

			intLine = append(intLine, number)
		}

		topographicMap = append(topographicMap, intLine)
	}

	headTrails := findHeadTrails(topographicMap)

	totalScore := 0
	totalRating := 0

	for _, headTrail := range headTrails {

		peaksFound := make([][]bool, len(topographicMap))

		for i := range peaksFound {
			peaksFound[i] = make([]bool, len(topographicMap[0]))
		}

		score, rating := followTrail(headTrail, topographicMap, directions, peaksFound)
		totalScore += score
		totalRating += rating
	}

	fmt.Println("First Challenge: ", totalScore)
	fmt.Println("Second Challenge: ", totalRating)
}