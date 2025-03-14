package main

import "fmt"

func decode(input string) string {
	inputLength := len(input)
	output := make([]int, inputLength+1)

	for i := range inputLength {
		if input[i] == '=' {
			output[i+1] = output[i]
		} else if input[i] == 'R' {
			output[i+1] = output[i] + 1
		} else if input[i] == 'L' {
			if output[i] <= output[i+1] {
				output[i] = output[i+1] + 1
				for j := i - 1; j >= 0; j-- {
					if input[j] == 'L' && output[j] <= output[j+1] {
						output[j] = output[j+1] + 1
					}
					if input[j] == '=' {
						output[j] = output[j+1]
					}
					if input[j] == 'R' {
						break
					}
				}
			}
		}
	}

	result := ""
	for _, digit := range output {
		result += fmt.Sprintf("%d", digit)
	}

	return result
}

func main() {
	inputs := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}

	for _, input := range inputs {
		fmt.Printf("input = '%s' and output = '%s' \n", input, decode(input))
	}
}
