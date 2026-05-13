// p1 good
// p2 uses sig digits for an incorrect solution

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
	setProblems(lines)

	fmt.Printf("part 1: %d\n", part(1))
	fmt.Printf("part 2: %d\n", part(2))
}

func part(part int) int {
	pcount := len(problems[0])
	plen := len(problems)
	sum := 0

	if part == 1 {
		// problem length
		for i := range pcount {
			op := operators[i]
			result := 0
			for j := range plen {
				n, _ := strconv.Atoi(problems[j][i])
				if j == 0 {
					result = n
					continue
				}

				if op == "+" {
					result += n
				} else { // *
					result *= n
				}
			}
			sum += result
		}
		return sum
	} else { // p2
		// for each problem..
		for i := range pcount {
			op := operators[i]
			result := 0

			// find longest number
			maxDigits := 0
			for j := range plen {
				s := problems[j][i]
				if len(s) > maxDigits {
					maxDigits = len(s)
				}
			}

			// for each significant digit..
			operands := []int{}
			for digit := range maxDigits {
				operandSB := strings.Builder{}
				// for each number..
				for j := range plen {
					s := problems[j][i]
					if len(s) > digit {
						n, _ := strconv.Atoi(s)
						// add digit to operand string
						if digit == 0 {
							operandSB.WriteByte(s[len(s)-1])
						} else {
							operandSB.WriteString(strconv.Itoa(getSignificantDigit(n, digit)))
						}
					}
				}
				operand, _ := strconv.Atoi(operandSB.String())
				operands = append(operands, operand)

				for operandIndex, n := range operands {
					if operandIndex == 0 {
						result = n
						continue
					}

					if op == "+" {
						result += n
					} else { // *
						result *= n
					}
				}
			}
			sum += result
		}
		return sum
	}
}
func setProblems(lines []string) {
	rDigits := regexp.MustCompile(`\d+`)
	rOperators := regexp.MustCompile(`\+|\*`)

	for _, l := range lines {
		matches := rDigits.FindAllString(l, -1)
		if len(matches) > 0 {
			problems = append(problems, matches)
		} else {
			// last line is operators
			matches := rOperators.FindAllString(l, -1)
			operators = matches
		}
	}
}

func getSignificantDigit(n, digit int) int {
	for range digit {
		n /= 10
	}
	return n % 10
}
