package main

import (
    "fmt"
    "sort"
)

type bylength []string

func (s bylength) Len() int {
    return len(s)
}

func (s bylength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s bylength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(bylength(fruits))
    fmt.Println(fruits)
}
