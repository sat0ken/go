package main

import (
    "fmt"
    "os"
    "strconv"

    "ch2/tempconv"
    "ch2/weightconv"
    "ch2/lenconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        f, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "parse %v to float: %v", arg, err)
        }
        fmt.Printf("%s = %s, %s = %s\n",
            tempconv.Fahrenheit(f), tempconv.FToC(tempconv.Fahrenheit(f)),
            tempconv.Celsius(f), tempconv.CToF(tempconv.Celsius(f)))
        fmt.Printf("%s = %s, %s = %s\n",
            weightconv.Pound(f), weightconv.PToK(weightconv.Pound(f)),
            weightconv.Kilogram(f), weightconv.KToP(weightconv.Kilogram(f)))
        fmt.Printf("%s = %s, %s = %s\n",
            lenconv.Foot(f), lenconv.FToM(lenconv.Foot(f)),
            lenconv.Meter(f), lenconv.MToF(lenconv.Meter(f)))
    }
}
