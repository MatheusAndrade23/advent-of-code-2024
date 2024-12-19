package main

import (
	"bufio"
	"fmt"
	"os"
)

func determineValidAntiNodesQuantityWithTwoAntennas(x1 int, x2 int, y1 int, y2 int, antennasMap []string, placesWithAntiNodeChecked [][]bool) int{

	xDifference := x2 - x1
	yDifference := y2 - y1

	node1x := x1 - xDifference
	node1y := y1 - yDifference

	node2x := x2 + xDifference
	node2y := y2 + yDifference

	nodesValidQuantity := 0

	if node1x >= 0 && node1x < len(antennasMap) && node1y >= 0 && node1y < len(antennasMap[0]) && !placesWithAntiNodeChecked[node1x][node1y] {
		placesWithAntiNodeChecked[node1x][node1y] = true
		nodesValidQuantity++
	}

	if node2x >= 0 && node2x < len(antennasMap) && node2y >= 0 && node2y < len(antennasMap[0]) && !placesWithAntiNodeChecked[node2x][node2y] {
		placesWithAntiNodeChecked[node2x][node2y] = true
		nodesValidQuantity++
	}

	return nodesValidQuantity
}

func determineValidAntiNodesQuantityWithTwoAntennasConsideringResonantHarmonics(x1 int, x2 int, y1 int, y2 int, antennasMap []string, placesWithAntiNodeChecked [][]bool) int{
	xDifference := x2 - x1
	yDifference := y2 - y1

	nodesValidQuantity := 0

	for {
		if x1 >= 0 && x1 < len(antennasMap) && y1 >= 0 && y1 < len(antennasMap[0]) {
			if(!placesWithAntiNodeChecked[x1][y1]) {
				placesWithAntiNodeChecked[x1][y1] = true
				nodesValidQuantity++
			}

			x1 = x1 - xDifference
			y1 = y1 - yDifference
		} else {
			break
		}
	}

	for {
		if x2 >= 0 && x2 < len(antennasMap) && y2 >= 0 && y2 < len(antennasMap[0]) {
			if(!placesWithAntiNodeChecked[x2][y2]) {
				placesWithAntiNodeChecked[x2][y2] = true
				nodesValidQuantity++
			}

			x2 = x2 + xDifference
			y2 = y2 + yDifference
		} else {
			break
		}
	}

	return nodesValidQuantity
}


func main() {

	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	var antennasMap []string
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		antennasMap = append(antennasMap, line)
	}

	placesWithAntiNodeChecked := make([][]bool, len(antennasMap))

	for i := range placesWithAntiNodeChecked {
		placesWithAntiNodeChecked[i] = make([]bool, len(antennasMap[0]))
	}


	secondPlacesWithAntiNodeChecked := make([][]bool, len(antennasMap))

	for i := range secondPlacesWithAntiNodeChecked {
		secondPlacesWithAntiNodeChecked[i] = make([]bool, len(antennasMap[0]))
	}

	totalNodePoints := 0
	totalNodePointsConsideringResonantHarmonics := 0

	for i := 0; i < len(antennasMap); i++ {
		for j := 0; j < len(antennasMap[i]); j++ {
			if antennasMap[i][j] != '.' {
					for k := 0; k < len(antennasMap); k++ {
						for l := 0; l < len(antennasMap[k]); l++ {
							if antennasMap[k][l] == antennasMap[i][j] && !(k == i && l == j)  {
								validAntiNodesQuantity := determineValidAntiNodesQuantityWithTwoAntennas(i, k, j, l, antennasMap, placesWithAntiNodeChecked)
								validAntiNodesQuantityConsideringResonantHarmonics := determineValidAntiNodesQuantityWithTwoAntennasConsideringResonantHarmonics(i, k, j, l, antennasMap, secondPlacesWithAntiNodeChecked)
								totalNodePoints += validAntiNodesQuantity
								totalNodePointsConsideringResonantHarmonics += validAntiNodesQuantityConsideringResonantHarmonics
							}
						} 
					}
				
			}
		} 
	}

	fmt.Println("First Challenge", totalNodePoints)
	fmt.Println("Second Challenge", totalNodePointsConsideringResonantHarmonics)
}