package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }
    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }
    fname := fmt.Sprintf("fetchall_%d.txt", time.Now())
    f, err := os.Create(fname)
    if err != nil {
        ch <- fmt.Sprintf("creating file %s: %v", fname, err)
        return
    }

    nbytes, err := io.Copy(f, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
