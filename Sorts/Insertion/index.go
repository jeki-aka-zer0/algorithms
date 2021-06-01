package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter comma separated list for sorting: ")

    inputStr, _ := reader.ReadString('\n')

    listRaw := strings.Split(inputStr, ",")
    listRawLen := len(listRaw)
    list := make([]int, 0, listRawLen)
    for i := 0; i < listRawLen; i++ {
        trimmedVal := strings.TrimSpace(listRaw[i])
        if numberPure, err := strconv.Atoi(trimmedVal); err == nil {
            list = append(list, numberPure)
        }
    }

    fmt.Println(insertionSort(list))
}

func insertionSort(list []int) []int {
    for i := 1; i < len(list); i++ {
        curVal := list[i]
        j := i - 1;
        for ;j >= 0 && list[j] > curVal; j-- {
            list[j + 1] = list[j]
        }
        list[j + 1] = curVal
    }

    return list
}