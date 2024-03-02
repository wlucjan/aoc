package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitLink(s string, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

type Range struct {
	left  int
	right int
}

func (thisRange Range) contains(otherRange Range) bool {
	return thisRange.left <= otherRange.left && thisRange.right >= otherRange.right
}

func (thisRange Range) overlaps(otherRange Range) bool {
	return (thisRange.left <= otherRange.left && thisRange.right >= otherRange.left) || (thisRange.left <= otherRange.right && thisRange.right >= otherRange.right)
}

func createRange(rangeAsString string) Range {
	left, right := splitLink(rangeAsString, "-")

	convertedLeft, err := strconv.Atoi(left)

	if err != nil {
		fmt.Println(err)
	}

	convertedRight, err := strconv.Atoi(right)

	if err != nil {
		fmt.Println(err)
	}

	return Range{convertedLeft, convertedRight}
}

func evaluateRow(row string) (contained bool, overlaped bool) {
	firstElfSections, secondElfSections := splitLink(row, ",")
	firstElfRange := createRange(firstElfSections)
	secondElfRange := createRange(secondElfSections)

	return firstElfRange.contains(secondElfRange) || secondElfRange.contains(firstElfRange), firstElfRange.overlaps(secondElfRange) || secondElfRange.overlaps(firstElfRange)
}

func readFile(file string) (int, int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	containedRows := []string{}
	overlapedRows := []string{}

	for fileScanner.Scan() {
		row := fileScanner.Text()
		contains, overlaps := evaluateRow(row)

		if contains {
			containedRows = append(containedRows, row)
		}

		if overlaps {
			overlapedRows = append(overlapedRows, row)
		}
	}

	return len(containedRows), len(overlapedRows)
}

func main() {
	fmt.Println(readFile("input.txt"))
}
