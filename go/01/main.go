package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// pull data from input file
	// split from "   "
	// convert to integers and sort numerically
	// add difference of lowest integers for each line consecutively for answer
	var left = make([]int, 0)
	var right = make([]int, 0)
	var answer int = 0

	counter := 0

	file, err := os.Open("input_data.txt")
	if err != nil {
		fmt.Printf("Error Reading File: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
		nums := strings.Split(scanner.Text(), "   ")
		for i, value := range nums {
			//fmt.Printf("[%v] %v\n", i, value)
			if i == 0 {
				left = strToIntSorted(value)
			} else {
				right = strToIntSorted(value)
			}
		}
		fmt.Printf("Current: %v\n", answer)
		answer += getDifferenceIntSlice(left, right)
	}
	fmt.Println("Answer: ", answer)
	fmt.Printf("Processed %v Lines\n", counter)
}

func strToIntSorted(str string) []int {
	var answer = make([]int, 0)
	num := strings.Split(str, "")
	for _, value := range num {
		nums, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("error converting to int", err)
			continue
		}
		answer = append(answer, nums)
	}
	return answer
}

func getDifferenceIntSlice(num1 []int, num2 []int) int {
	// sort slices into ascending order
	sort.Ints(num1)
	sort.Ints(num2)

	var answer int = 0

	for i := range num1 {
		answer += int(math.Abs(float64(num1[i] - num2[i])))
		//fmt.Printf("Num1:%v - Num2:%v\n", num1[i], num2[i])
	}
	fmt.Println("Line Answer: ", answer)
	return answer
}
