package main

import (
    "fmt"
    "os"
)

func main() {
    for index, args := range os.Args[1:] {
        fmt.Printf("%d, %s\n", index, args)
    }
}
