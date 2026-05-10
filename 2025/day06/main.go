package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var problems [][]string
var operators []string

func main() {
	lines := strings.Split(input, "\n")
	problems, operators = parseProblems(lines)

	fmt.Printf("part 1: %d\n", part(1))
	fmt.Printf("part 2: %d\n", part(2))
}

func part(part int) int {
	sum := 0

	if part == 1 {
		for i, operands := range problems {
			result := 0
			for j, operandString := range operands {
				operand, _ := strconv.Atoi(strings.TrimSpace(operandString))
				// avoid 0 * operand
				if j == 0 {
					result = operand
					continue
				}
				if operators[i] == "+" {
					result += operand
				} else { // *
					result *= operand
				}
			}
			sum += result
		}
	} else { // part 2
		for i, operands := range problems {
			cephalopodOperands := []int{}
			result := 0
			// transform operands by string index
			for operandDigitIndex := 0; operandDigitIndex < len(operands[0]); operandDigitIndex++ {
				cephalopodOperandSB := strings.Builder{}
				for _, operand := range operands {
					cephalopodOperandSB.WriteByte(operand[operandDigitIndex])
				}
				cephalopodOperand, _ := strconv.Atoi(strings.TrimSpace(cephalopodOperandSB.String()))
				if cephalopodOperand != 0 {
					cephalopodOperands = append(cephalopodOperands, cephalopodOperand)
				}
			}
			// then operate
			for j, cephOperand := range cephalopodOperands {
				// avoid 0 * operand
				if j == 0 {
					result = cephOperand
					continue
				}
				if operators[i] == "+" {
					result += cephOperand
				} else { // *
					result *= cephOperand
				}
			}
			sum += result
		}
	}
	return sum
}

func parseProblems(lines []string) ([][]string, []string) {
	// init operators
	operatorLine := lines[len(lines)-1]
	reOperators := regexp.MustCompile(`\+|\*`)
	operators = reOperators.FindAllString(operatorLine, -1)
	operatorIndexRanges := reOperators.FindAllIndex([]byte(operatorLine), -1)
	var operatorIndexes []int
	for _, oir := range operatorIndexRanges {
		operatorIndexes = append(operatorIndexes, oir[0])
	}
	// append final operator to help with operand slicing
	operatorIndexes = append(operatorIndexes, len(operatorLine))

	// init problems
	var problems [][]string
	problemLines := lines[:len(lines)-1]
	for i, currOperatorI := range operatorIndexes {
		// break on final index
		if i == len(operatorIndexes)-1 {
			break
		}
		problems = append(problems, []string{})
		nextOperatorI := operatorIndexes[i+1]
		// populate operands for each problem column
		for _, l := range problemLines {
			operand := l[currOperatorI:nextOperatorI]
			problems[i] = append(problems[i], operand)
		}
	}
	return problems, operators
}
