package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Procedure struct {
    source int
    target int
    count int
}

type apply func(Procedure)

type CrateMover9000 struct {}
func (crane *CrateMover9000) move(stacks [][]string, procedure Procedure) {
    source := procedure.source - 1
    target := procedure.target - 1
    fmt.Println(stacks[source])
    fmt.Println(stacks[target])
    for i := 0; i < procedure.count; i++ {
        stacks[target] = append([]string{stacks[source][i]}, stacks[target]...)
    }
    stacks[source] = stacks[source][procedure.count:]
    fmt.Println(stacks[source])
    fmt.Println(stacks[target])
}

type CrateMover9001 struct {}
func (crane *CrateMover9001) move(stacks [][]string, procedure Procedure) {
    source := procedure.source - 1
    target := procedure.target - 1
    fmt.Println(stacks[source])
    fmt.Println(stacks[target])
    for i := procedure.count - 1; i >= 0; i-- {
        stacks[target] = append([]string{stacks[source][i]}, stacks[target]...)
    }
    stacks[source] = stacks[source][procedure.count:]
    fmt.Println(stacks[source])
    fmt.Println(stacks[target])
}

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
        fn(procedure)
	}
    return nil
}

func parseStacks(file string) (stacks [][]string, err error) {
    readFile, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
        return stacks, err
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        row := strings.Split(fileScanner.Text(), "")
        stacksCount := len(row)/4 + 1
        if stacks == nil {
            stacks = make([][]string ,stacksCount)
        }
        for i := range stacks {
            if letter := row[i*4 + 1]; strings.TrimSpace(letter) != "" {
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

    applyProcedureToStacks := func (stacks [][]string)  (apply) {
        return func (procedure Procedure) {
            crane := CrateMover9001{}
            crane.move(stacks, procedure)
        }
    }

    err = procedureScanner("procedure.txt", applyProcedureToStacks(stacks))
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(stacks)
    
    fmt.Println()
    for _, stack := range stacks {
        fmt.Print(string(stack[0]))
    }
}
