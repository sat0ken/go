package main

import (
    "fmt"
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "os"
    "bufio"
    "strconv"
)

var opt = flag.String("n", "256", "Generate sha256/384/512")

func main() {
    flag.Parse()
    algo, err := strconv.Atoi(*len)
    if err != nil {
        algo = 256
    }
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        text := input.Text()
        switch algo {
        case 256:
            fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
        case 384:
            fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
        case 512:
            fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
        default:
            fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
        }
    }
}
