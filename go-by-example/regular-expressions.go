package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    var p = fmt.Println
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    p(match)

    r, _ := regexp.Compile("p([a-z]+)ch")
    p(r.MatchString("peach"))
    p(r.FindString("peach punch"))
    p(r.FindStringIndex("peach punch"))
    p(r.FindStringSubmatch("peach punch"))
    p(r.FindStringSubmatchIndex("peach punch"))
    p(r.FindAllString("peach punch pinch", -1))
    p(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    p(r.FindAllString("peach punch pinch", 2))
    p(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    p(r)

    p(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    p(string(out))
}
