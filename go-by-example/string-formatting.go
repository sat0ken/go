package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

var pf = fmt.Printf

func main() {
    p := point{1, 2}

    pf("%v\n", p)
    pf("%+v\n", p)
    pf("%#v\n", p)

    pf("%T\n", p)
    pf("%d\n", true)
    pf("%d\n", 123)

    pf("%b\n", 14)
    pf("%c\n", 33)

    pf("%x\n", 456)
    pf("%f\n", 78.9)

    pf("%e\n", 123400000.0)
    pf("%E\n", 123400000.0)

    pf("%s\n", "\"string\"")
    pf("%q\n", "\"string\"")

    fmt.Printf("%x\n", "hex this")
    fmt.Printf("%p\n", &p)

    fmt.Printf("|%6d|%6d|\n", 12, 345)

    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
