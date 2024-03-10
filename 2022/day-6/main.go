package main

import (
    "bufio"
    "fmt"
    "os"
)

func areCharactersDistinct(slice []rune) bool {
    charCount := make(map[rune]int)

    for _, char := range slice {
        charCount[char]++
        if charCount[char] > 1 {
            return false
        }
    }
    return true
}

func main() {
    file, err := os.Open("signal.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    index := 0
    window := make([]rune, 0)
    windowSize := 14

    for {
        char, _, err := reader.ReadRune()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            fmt.Println("Error:", err)
            return
        }

        window = append(window, char)
        index++

        if len(window) > windowSize {
            window = window[1:]
        }

        if len(window) == windowSize && areCharactersDistinct(window) {
            break
        }
    }
    fmt.Println(index)
}
