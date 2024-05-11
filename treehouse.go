package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: treehouse <filename>")
		return
	}

	filename := os.Args[1]
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Error converting character to integer:", err)
				return
			}
			row = append(row, digit)
		}
		numbers = append(numbers, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rowSize := len(numbers)
	colSize := len(numbers[0])
	var count int = colSize * rowSize
	for i := 1; i < (rowSize - 1); i++ {
		// Loop over columns
		for j := 1; j < (colSize - 1); j++ {
			canSee := []bool{true, true, true, true} // Represent 4 direction {left,right,upper,lower}
			// Check left side
			for left := j - 1; left >= 0; left-- {
				if numbers[i][left] >= numbers[i][j] {
					canSee[0] = false
					break
				}
			}
			// Check right side
			for right := j + 1; right < colSize; right++ {
				if numbers[i][right] >= numbers[i][j] {
					canSee[1] = false
					break
				}
			}
			// Check upper side
			for upper := i - 1; upper >= 0; upper-- {
				if numbers[upper][j] >= numbers[i][j] {
					canSee[2] = false
					break
				}
			}
			// Check lower side
			for lower := i + 1; lower < rowSize; lower++ {
				if numbers[lower][j] >= numbers[i][j] {
					canSee[3] = false
					break
				}
			}
			// Can't see all angle
			allFalse := true
			for _, value := range canSee {
				if value {
					allFalse = false
					break
				}
			}
			if allFalse {
				count--
			}
		}
	}

	fmt.Print(count)
}
