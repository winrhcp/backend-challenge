package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxPath(data [][]int) int {
	n := len(data)

	for row := n - 2; row >= 0; row-- {
		for col := range data[row] {
			data[row][col] += max(data[row+1][col], data[row+1][col+1])
		}
	}

	return data[0][0]
}

func main() {
	var data [][]int

	dataPath := "hard.json"
	rawData, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(rawData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// fmt.Println("data:", data)

	result := findMaxPath(data)
	fmt.Println("The most valuable route has a total value of:", result)
}
