package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Procedure struct {
    source int
    target int
    count int
}

type apply func(Procedure)

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func parseRow(row string) (procedure Procedure, err error) {
    split := strings.Split(row, " ");
    paramsAsString := []string{split[1], split[3], split[5]}
    paramsAsInt, err := stringsToIntegers(paramsAsString)

    if err != nil {
        return procedure, err
    }

    return Procedure{paramsAsInt[1], paramsAsInt[2], paramsAsInt[0]}, nil
}

func applyProcedure(procedure Procedure) {
    fmt.Sprintln("%v\n", procedure)
}

func procedureScanner(file string, fn apply) (err error) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
        return err
	}
    defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
        procedure, err := parseRow(row)
        if err != nil {
            return err
        }
        fmt.Println(procedure)
        fn(procedure)
	}
    return nil
}

func parseStacks(file string) (stacks [][]rune, err error) {
    readFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
        return stacks, err
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)


    for fileScanner.Scan() {
        row := fileScanner.Text()
        stacksCount := len(row)/4 + 1
        stacks = make([][]rune, stacksCount)
        fmt.Println(stacksCount)
        for i := range stacks {
            if letter := rune(row[i*4 + 1]); !unicode.IsSpace(letter) {
                stacks[i] = append(stacks[i], letter);
            }
        }
    }

    return stacks, nil
}

func main() {
    stacks, err := parseStacks("cargo.txt")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Print(stacks)

    err = procedureScanner("procedure.txt", applyProcedure)
    if err != nil {
        fmt.Println(err)
    }
}
