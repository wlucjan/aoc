package main

import (
	"bufio"
	"fmt"
	"os"
)

func procedureScanner(file string) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	for fileScanner.Scan() {
		row := fileScanner.Text()

	}
}

func main() {
	fmt.Println()
}
