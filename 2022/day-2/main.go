package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseGameStrategy(gameStrategy string) [][]string {
	return lo.Map(strings.Split(gameStrategy, "\n"), func(data string, index int) []string {
		return strings.Split(data, " ")
	})
}

func parseExchangeWithWinStrategy(exchange []string) (playerHand string, opponentHand string, result string) {
	expectedResultMap := map[string]string{"X": "lose", "Y": "draw", "Z": "win"}
	opponentShapeMap := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	handNeededForResult := map[string]map[string]string{
		"rock": {
			"draw": "rock",
			"win":  "paper",
			"lose": "scissors",
		},
		"paper": {
			"lose": "rock",
			"draw": "paper",
			"win":  "scissors",
		},
		"scissors": {
			"win":  "rock",
			"lose": "paper",
			"draw": "scissors",
		},
	}

	opponentHand = opponentShapeMap[exchange[0]]
	result = expectedResultMap[exchange[1]]
	playerHand = handNeededForResult[opponentHand][result]

	return playerHand, opponentHand, result
}

func parseExchangeWithHandStrategy(exchange []string) (playerHand string, opponentHand string, result string) {
	playerShapeMap := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	opponentShapeMap := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}

	exchangeResultMap := map[string]map[string]string{
		"rock": {
			"rock":     "draw",
			"paper":    "lose",
			"scissors": "win",
		},
		"paper": {
			"rock":     "win",
			"paper":    "draw",
			"scissors": "lose",
		},
		"scissors": {
			"rock":     "lose",
			"paper":    "win",
			"scissors": "draw",
		},
	}

	opponentHand = opponentShapeMap[exchange[0]]
	playerHand = playerShapeMap[exchange[1]]
	result = exchangeResultMap[playerHand][opponentHand]

	return playerHand, opponentHand, result
}

func scoreExchangeWithWinStrategy(exchange []string, index int) int {
	playerHand, _, result := parseExchangeWithWinStrategy(exchange)

	scoreForShape := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	scoreForExchange := map[string]int{"lose": 0, "draw": 3, "win": 6}

	return scoreForShape[playerHand] + scoreForExchange[result]
}

func scoreExchangeWithHandStrategy(exchange []string, index int) int {
	playerHand, _, result := parseExchangeWithHandStrategy(exchange)

	scoreForShape := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	scoreForExchange := map[string]int{"lose": 0, "draw": 3, "win": 6}

	return scoreForShape[playerHand] + scoreForExchange[result]
}

func main() {
	gameStrategy, err := os.ReadFile("input.txt")
	check(err)

	parsedGameStrategy := parseGameStrategy(string(gameStrategy))

	totalScore := lo.Sum(lo.Map(parsedGameStrategy, scoreExchangeWithWinStrategy))

	fmt.Println(totalScore)
}
