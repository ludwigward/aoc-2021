package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	partA()
	partB()
}

func partA() {
	file, err := os.Open("./input_day1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var row, increases, previous, next int

	for scanner.Scan() {
		if row == 0 {
			previous, err = strconv.Atoi(scanner.Text())
		} else {
			next, err = strconv.Atoi(scanner.Text())

			if next > previous {
				increases++
			}
		}

		row++
		previous = next
	}

	fmt.Println("Increases part 1: ", increases)
}

func partB() {
	file, err := os.Open("./input_day1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var row, tmp, previous, next, increases int

	values := make([]int, 0)

	for scanner.Scan() {
		tmp, err = strconv.Atoi(scanner.Text())
		values = append(values, tmp)
	}

	for row < len(values) {
		if row > 2 {
			previous = values[row-3] + values[row-2] + values[row-1]
			next = values[row-2] + values[row-1] + values[row]

			if next > previous {
				increases++
			}
		}

		row++
	}

	fmt.Println("Increases part 2: ", increases)
}
