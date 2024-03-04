package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseElfCalories(unparsedElfCalories string, index int) int {
	elfCalories := strings.Split(unparsedElfCalories, "\n")
	totalCalories := 0
	for _, calories := range elfCalories {
		caloriesAsInt, err := strconv.ParseInt(calories, 0, 0)
		check(err)
		totalCalories += int(caloriesAsInt)
	}

	return totalCalories
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	elvesWithCalories := lo.Map[string, int](strings.Split(string(data), "\n\n"), parseElfCalories)
	sort.Ints(elvesWithCalories)

	topCalories := elvesWithCalories[len(elvesWithCalories)-1]
	topThreeHighestCalories := elvesWithCalories[len(elvesWithCalories)-3:]

	fmt.Println(topCalories, lo.Sum(topThreeHighestCalories))
}
