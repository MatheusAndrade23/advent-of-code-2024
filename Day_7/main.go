package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation []int

// func calculateWithOnly2Operators(equation Equation, index int, current int, result *bool){

// 	shouldResult := equation[0]
// 	numbers := equation[1:]

// 	if index == len(numbers) {
// 		if current == shouldResult{
// 			*result = true
// 		}
// 		return
// 	}

// 	next := numbers[index]

// 	calculateWithOnly2Operators(equation, index+1, current*next, result)

// 	calculateWithOnly2Operators(equation, index+1, current+next, result)
// }

func calculateWith3Operators(equation Equation, index int, current int, result *bool){

	shouldResult := equation[0]
	numbers := equation[1:]

	if index == len(numbers) {
		if current == shouldResult{
			*result = true
		}
		return
	}

	next := numbers[index]

	calculateWith3Operators(equation, index+1, current*next, result)

	calculateWith3Operators(equation, index+1, current+next, result)

	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, next))
	calculateWith3Operators(equation, index+1, concat, result)
}

func isEquationValid(equation Equation) bool {

	shouldResult := equation[0]
	numbers := equation[1:]

	totalPossibleCombinations := int(math.Pow(2, float64(len(numbers)-1)))

	for i := 0; i < totalPossibleCombinations; i++ {

		result := numbers[0]
		binary := fmt.Sprintf("%0*b", len(numbers)-1, i) // each binary is a possible combination

		for j, bit := range binary {

			if bit == '1' {
					result += numbers[j+1] 
			} else if bit == '0' {
					result *= numbers[j+1] 
			}
		}

		if result == shouldResult{
			return true
		}
	}

	return false
}

func tryCalculate(equation Equation) bool {
	canBeCalculated := false

	calculateWith3Operators(equation, 1, equation[1], &canBeCalculated)

	return canBeCalculated
}

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	var equations []Equation

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		result := fields[0][:len(fields[0]) - 1]

		intResult, err := strconv.Atoi(result)

		if err != nil{
			panic(err)
		}

		equation := make([]int, 0, len(fields)-1)
		equation = append(equation, intResult)

		for _, field := range fields[1:] {
			intField, err := strconv.Atoi(field)

			if err != nil{
				panic(err)
			}

			equation = append(equation, intField)
		}

		equations = append(equations, equation)
	}

	sum := 0
	sumWithThirdOperator := 0

	for _, equation := range equations {

		validWith2Operators := isEquationValid(equation)
		validWith3Operators := tryCalculate(equation)

		if validWith2Operators {
			sum += equation[0]
		}

		if !validWith2Operators && validWith3Operators{
			sumWithThirdOperator+=equation[0]
		}
	}

	fmt.Println(sum)
	fmt.Println(sum + sumWithThirdOperator)
}