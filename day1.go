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

	fmt.Println("Increases:", increases)
}
