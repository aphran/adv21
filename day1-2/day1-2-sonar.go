package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"

    "github.com/aphran/adv21/common"
)

func initRow(rSlice []int, mSize int) {
    for i := 0; i < mSize; i++ {
        rSlice[i] = -1
    }
}

func initGroups(gSlice [][]int, gStat []bool, mSize int) {
    for k := 0; k < mSize; k++ {
        gStat[k] = false
        gSlice[k] = make([]int, mSize)
        initRow(gSlice[k], mSize)
    }
}

func printGroups(gSlice [][]int, gStat []bool, mSize int) {
    fmt.Println("[ R  RowIn? x  y  z ]")
    for k := 0; k < mSize; k++ {
        fmt.Printf("[ %v  %v ", k, gStat[k])
        for _, v := range gSlice[k] {
            fmt.Printf("%v ", v)
        }
        fmt.Printf("]\n")
    }
}

func addNumber(gSlice [][]int, gStat []bool, mSize int, num int) (int) {
    var rowSum int
    newSum := -1
    var rowUsage int
    //Find 3 rows with room
    tookRowIn := false
    var addedNum bool
    //fmt.Println("Attempting to add number", num)

    for k := 0; k < mSize; k++ {
        //Check if this sliding window (row) is out
        //and take it in if we haven't taken any new
        //sliding windows in yet this time around
        if !gStat[k] && !tookRowIn {
            tookRowIn = true
            gStat[k] = true
        }

        //fmt.Println(fmt.Sprintf("Row %v status:%v", k, gStat[k]))
        //Track sum and usage as we traverse each row
        rowSum = 0
        rowUsage = 0
        addedNum = false
        for i := 0; i < mSize; i++ {
            //Is there room here?
            if gSlice[k][i] == -1 {
                //Is this row in and does the current number
                //still need to be added somewhere?
                if gStat[k] && !addedNum {
                    gSlice[k][i] = num
                    addedNum = true
                    rowUsage++
                    //fmt.Println("Just added number")
                    //if row k is now full we need to
                    //sum it up, clean it off and take it out
                    if rowUsage == mSize {
                        rowSum += gSlice[k][i]
                        newSum = rowSum
                        initRow(gSlice[k], mSize)
                        gStat[k] = false
                        //fmt.Println("Row", k, "is full, cleaned it! Had sum:", rowSum)
                    }
                }
            } else {
                rowUsage++
            }
            if gSlice[k][i] >= 0 {
                rowSum += gSlice[k][i]
            }
        }

    }
    return newSum
}

func main() {
    file, err := common.OpenDayData(1)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    biggers := 0
    var currentNumber int
    var currentSum int
    var previousSum int
    havePrevious := false

    //Adjust sliding window measurement size
    const mSize = 3
    //Track sliding window status (in/out)
    gStat := make([]bool, mSize)
    //Shelf storage for sliding windows
    gSlice := make([][]int, mSize)
    //gSlice[0][0] = -1
    //Initialize shelf as empty (-1) and out (false)
    initGroups(gSlice, gStat, mSize)
    //printGroups(gSlice, gStat, mSize)

    scans := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        scans++
        //printGroups(gSlice, gStat, mSize)
        currentLine := scanner.Text()
        currentNumber, err = strconv.Atoi(currentLine)
        if err != nil {
            fmt.Println("Converting line text to number did not work at all! ", err)
        }

        //Deal with sliding windows, if we get a result
        //from this we'll know there is a new sum to consider
        //which comes from a sliding window that was just closed
        currentSum = addNumber(gSlice, gStat, mSize, currentNumber)
        if currentSum >= 0 {
            if havePrevious {
                if currentSum > previousSum {
                    fmt.Println(currentSum, "is bigger!")
                    biggers = biggers + 1
                } else {
                    fmt.Println("         ", currentSum, "is not bigger :(")
                }
            } else {
                fmt.Println("**", currentSum, "is the first sum **")
            }
            previousSum = currentSum
            havePrevious = true
        }
    }
    fmt.Println("We found", biggers, "bigger measurement sliding window sums!")
    os.Exit(0)
}
