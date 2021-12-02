package main

import (
    "bufio"
    "fmt"
    "strconv"

    "github.com/aphran/adv21/common"
)

func main() {
    file, err := common.OpenDayData(1)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    biggers := 0
    var currentNumber int
    var previousNumber int
    havePrevious := false

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currentLine := scanner.Text()
        currentNumber, err = strconv.Atoi(currentLine)
        if err != nil {
            fmt.Println("Converting line text to number did not work at all! ", err)
        }
        if havePrevious {
            if currentNumber > previousNumber {
                fmt.Println(currentNumber, "is bigger!")
                biggers = biggers + 1
            } else {
                fmt.Println("         ", currentNumber, "is not bigger :(")
            }
        } else {
            fmt.Println("**", currentNumber, "is the first number **")
        }
        previousNumber = currentNumber
        havePrevious = true
    }
    fmt.Println("We found", biggers, "bigger numbers!")
}
