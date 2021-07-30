package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Println(bitCount(c1, c2))
}

func bitCount(a, b [sha256.Size]byte) int {
    var sum int
    for i := 0; i < len(a); i++ {
        c := a[i] ^ b[i]
        sum += int(PopCount(c))
    }
    return sum
}

func PopCount(b byte) uint8 {
    var pop uint8
    for i := uint8(0); i < uint8(8); i++ {
        if b&1 == 1 {
            pop++
        }
        b = b >> 1
    }
    return pop
}
