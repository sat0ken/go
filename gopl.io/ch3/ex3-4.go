package main

import (
    "fmt"
    "math"
    "net/http"
    "log"
    "bufio"
)

const (
    width, height = 600, 320
    cells = 100
    xyrange = 30.0
    xyscale = width / 2 / xyrange
    zscale = height * 0.4
    angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    http.HandleFunc("/", plotHandler)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}


func plotHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "image/svg+xml")
    header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
    out := bufio.NewWriter(w)
    out.WriteString(header)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
            if math.IsNaN(ax) || math.IsNaN(ay) ||
                math.IsNaN(bx) || math.IsNaN(by) ||
                math.IsNaN(cx) || math.IsNaN(cy) ||
                math.IsNaN(dx) || math.IsNaN(dy) {
                    continue
            }
            color := "#000000"
            if az > 0 && bz > 0 && cz > 0 && dz > 0 {
                color = "#ff0000"
            } else if az < 0 && bz < 0 && cz < 0 && dz < 0 {
                color = "#0000ff"
            }
            polygon := fmt.Sprintf("<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                color, ax, ay, bx, by, cx, cy, dx, dy)
            out.WriteString(polygon)
		}
	}
    out.WriteString("</svg>")
    out.Flush()
}

func corner(i, j int) (float64, float64, float64) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)

    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale

    return  sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}
