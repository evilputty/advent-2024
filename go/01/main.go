package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// pull data from input file
	// split from "   "
	// convert to integers and sort numerically
	// add difference of lowest integers for each line consecutively for answer

	file, err := os.Open("input_data.txt")
	if err != nil {
		fmt.Printf("Error Reading File: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		get_answer(strings.Split(scanner.Text(), "   "))
	}
}

func get_answer(line []string) int {
	// split the string data coming in like "12345   12345"
	for i, line := range line{
		num, err := strconv.Atoi(line)
		if err != nil {}
			fmt.Println("Error Converting to integer", err)
			continue
		}
		nums[i] = num
	}

	
	var answer int // placeholder
	return answer
}
