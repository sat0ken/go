package main

import (
    "fmt"
    "log"
    "net/http"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "strconv"
    "os"
    //"time"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    param := r.URL.Query()
    i, err := strconv.Atoi(param["cycles"][0])
    if err != nil {
        fmt.Fprintf(os.Stderr, "strconv err: %v",param["cycles"][0])
    }
	lissajous(w, float64(i))
}

func lissajous(out io.Writer, cycleNum float64) {
    const (
        //cycles  = cycleNum
        res     = 0.001
        size    = 100
        nframes = 64
        delay   = 8
    )
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycleNum*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
