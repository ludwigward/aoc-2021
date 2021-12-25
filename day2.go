package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input_day2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([]int, 0)
	directions := make([]string, 0)
	var stepSize int

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		stepSize, err = strconv.Atoi(instruction[1])
		values = append(values, stepSize)
		directions = append(directions, instruction[0])
	}

	partA(values, directions)
	partB(values, directions)
}

func partA(values []int, directions []string) {
	var depth, position, row int

	for row < len(values) {
		if directions[row] == "forward" {
			position += values[row]
		} else if directions[row] == "up" {
			depth -= values[row]
		} else if directions[row] == "down" {
			depth += values[row]
		}

		row++
	}

	fmt.Println("Part 1: ", depth, position, depth*position)
}

func partB(values []int, directions []string) {
	var depth, position, aim, row int

	for row < len(values) {
		if directions[row] == "forward" {
			position += values[row]
			depth -= values[row] * aim
		} else if directions[row] == "up" {
			aim += values[row]
		} else if directions[row] == "down" {
			aim -= values[row]
		}

		row++
	}

	fmt.Println("Part 2: ", depth, position, depth*position)
}
