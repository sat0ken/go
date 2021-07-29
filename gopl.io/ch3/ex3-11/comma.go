package main

import (
    "fmt"
    "strings"
)

func main() {
    s :=  []string{"1", "12", "123", "1234", "12345", "123456", "+1", "-1", "1234.5678", "-1234.5678"}
    for _, v := range s {
        fmt.Println(comma(v))
    }
}

func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        return comma(s[:dot]) + "."  +comma(s[dot+1:])
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}
