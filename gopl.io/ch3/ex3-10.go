package main

import (
    "fmt"
    "bytes"
)

func main() {
    fmt.Println(comma("12345"))
    fmt.Println(comma("678"))
}

func comma(s string) string {
    var buf bytes.Buffer
    n := len(s)
    if n <= 3 {
        return s
    }
    buf.WriteString(s[:n-3] + "," +s[n-3:])
    return buf.String()
}

