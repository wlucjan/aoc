package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func findFailedType(rucksuckContents string) (string, error) {
	middle := len(rucksuckContents) / 2
	containers := []string{rucksuckContents[:middle], rucksuckContents[middle:]}

	for _, char := range containers[0] {
		if strings.Contains(containers[1], string(char)) {
			return string(char), nil
		}
	}

	return "", errors.New("No failed type found")
}

func findBadge(rucksucks []string) (string, error) {
	for _, char := range rucksucks[0] {
		if strings.Contains(rucksucks[1], string(char)) && strings.Contains(rucksucks[2], string(char)) {
			return string(char), nil
		}
	}

	return "", errors.New("No badge found")
}

func prioritizeType(typeToPrioritize string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return strings.Index(alphabet, typeToPrioritize) + 1
}

func readFile(file string) (int, int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	failedTypePriorities := []int{}

	elvesGroupRucksucks := []string{}

	badgePriorities := []int{}

	for fileScanner.Scan() {
		// first task
		rucksuck := fileScanner.Text()
		failedType, err := findFailedType(rucksuck)
		if err != nil {
			fmt.Println(err)
		} else {
			priority := prioritizeType(failedType)

			failedTypePriorities = append(failedTypePriorities, priority)
		}

		elvesGroupRucksucks = append(elvesGroupRucksucks, rucksuck)

		if len(elvesGroupRucksucks) == 3 {
			badge, err := findBadge(elvesGroupRucksucks)
			if err != nil {
				fmt.Println(err)
			} else {
				badgePriorities = append(badgePriorities, prioritizeType(badge))
			}

			elvesGroupRucksucks = []string{}
		}

	}

	return lo.Sum(failedTypePriorities), lo.Sum(badgePriorities)
}

func main() {
	fmt.Println(readFile("input.txt"))
}
