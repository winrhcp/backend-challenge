package piefiredie

import "fmt"

func decode(input string) string {
	inputLength := len(input)
	output := make([]int, inputLength+1)

	for i := range inputLength {
		if input[i] == '=' {
			output[i+1] = output[i]
		} else if input[i] == 'R' {
			output[i+1] = output[i] + 1
		}
	}

	for i := inputLength - 1; i >= 0; i-- {
		if input[i] == 'L' {
			if output[i] <= output[i+1] {
				if output[i] <= output[i+1] {
					output[i] = output[i+1] + 1
				}
			}
		}
	}

	if inputLength > 0 && input[0] == '=' {
		output[0] = output[1]
	}

	result := ""
	for _, digit := range output {
		result += fmt.Sprintf("%d", digit)
	}

	return result
}

func RunPieFireDie() {
	inputs := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}

	for _, input := range inputs {
		fmt.Printf("input = %s and output = %s \n", input, decode(input))
	}
}
