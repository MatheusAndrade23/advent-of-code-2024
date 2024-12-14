package main

import (
	"bufio"
	"fmt"
	"os"
)

func countXMAS(grid [][]rune)int {
	word := "XMAS"
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)

	possibleDirections := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, direction := range possibleDirections {
				dr, dc := direction[0], direction[1]
				found := true

				for k := 0; k < wordLen; k++ {
					r := i + k * dr
					c := j + k * dc

					if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != rune(word[k]) {
						found = false
						break
					}
				}

				if found {
					count++
				}
			}
		}
	}

	return count;
}

type XDirection struct {
	Direction1 [2]int
	Direction2 [2]int
}

func countX_MAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	possibleDirections := [][2]int{
		{1, 1},
		
	}

	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'A' {
				for _, direction := range possibleDirections {
					dr1, dc1 := direction[0], direction[1]
					dr2, dc2 := -dr1, -dc1

					r1 := i + dr1
					r2 := i + dr2
					c1 := j + dc1
					c2 := j + dc2

					if r1 >= 0 && r1 < rows && c1 >= 0 && c1 < cols && r2 >= 0 && r2 < rows && c2 >= 0 && c2 < cols {
						if (grid[r1][c1] == 'M' && grid[r2][c2] == 'S') {
              if (grid[r2][c1] == 'S' && grid[r1][c2] == 'M'){
								count++
							}
            }

						if (grid[r1][c1] == 'S' && grid[r2][c2] == 'M') {
              if (grid[r2][c1] == 'M' && grid[r1][c2] == 'S'){
								count++
							}
            }

						if (grid[r1][c1] == 'M' && grid[r2][c2] == 'M') {
              if (grid[r2][c1] == 'S' && grid[r1][c2] == 'S'){
								count++
							}
            }

						if (grid[r1][c1] == 'S' && grid[r2][c2] == 'S') {
              if (grid[r2][c1] == 'M' && grid[r1][c2] == 'M'){
								count++
							}
            }
          }
				}
			}
		}
	}

	return count
}

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	fmt.Println("First Challenge: ", countXMAS(grid))
	fmt.Println("Second Challenge: ", countX_MAS(grid))
}