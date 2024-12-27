package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	//"common/input_data"
)

const INPUT_LINES int = 1000

func main() {
	// pull data from input file
	// split from "   "
	// convert to integers and sort numerically
	// add difference of lowest integers for each line consecutively for answer
	var left = make([]int, INPUT_LINES)
	var right = make([]int, INPUT_LINES)
	var answer int = 0
	var answer2 int = 0

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
			//convert str to int
			value, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Err converting to integer", err)
			}
			if i%2 == 0 {
				left = append(left, value)
				continue
			} else {
				right = append(right, value)
				continue
			}
		}
	}

	// sort before we do any funny biz
	sort.Ints(left)
	sort.Ints(right)

	answer = getDifferenceIntSlice(left, right)
	fmt.Println("Part 1 Answer: ", answer)

<<<<<<< HEAD
	answer2 := 0

	for _, value := range left {
		answer2 += (value * heatMap(right, value))
	}
=======
	for _, v := range left {
		answer2 += (v * getHits(right, v))
	}

>>>>>>> 2b2947bba74578fea7217a336f27b4920d084998
	fmt.Println("Part 2 Answer: ", answer2)
}

func getDifferenceIntSlice(num1 []int, num2 []int) int {
	var answer int = 0

	for i := range num1 {
		answer += int(math.Abs(float64(num1[i] - num2[i])))
	}
	return answer
}

<<<<<<< HEAD
func heatMap(slice []int, cmp int) int {
=======
func getHits(slice []int, cmp int) int {
>>>>>>> 2b2947bba74578fea7217a336f27b4920d084998
	count := 0
	for _, value := range slice {
		if value == cmp {
			count++
		}
<<<<<<< HEAD
=======
		if value > cmp {
			break
		}
>>>>>>> 2b2947bba74578fea7217a336f27b4920d084998
	}
	return count
}
