package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(IsAnagram("gaggi", "gagig"))
    fmt.Println(IsAnagram("foobar", "foorar"))
}

func IsAnagram(a, b string) bool {
    for i := range a {
        if strings.Index(b, string(a[i])) == -1 || strings.Index(a, string(b[i])) == -1 {
            return false
        }
    }
    return true
}
