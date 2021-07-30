package main

import (
    "fmt"
)

func main() {
    var arr = []int{1, 2, 3, 4}
    fmt.Println(rotate(arr, 4))
    fmt.Println(rotate2(arr, 4))
}

func rotate(arr []int, num int) []int {
    var rotated []int
    rotated = append(rotated, arr[num:]...)
    rotated = append(rotated, arr[:num]...)
    return rotated
}

func rotate2(arr []int, num int) []int {
    n := len(arr)
    rotated := make([]int, n)
    for i := range arr {
        rotated[i] = arr[(i+num)%n]
    }
    return rotated
}
