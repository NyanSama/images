// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 93.
//!+main

//Surface computes an SVG web server
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type function func(x, y float64) float64

var (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       float64
	zscale        float64
	angle         = math.Pi / 6
	f             function
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		draw(w, r)
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func draw(out http.ResponseWriter, in *http.Request) {
	setconf(in)
	svg(out, f)
}

func setconf(in *http.Request) {
	fmt.Println("Set config's")
	xwidth, _ := strconv.Atoi(in.FormValue("width"))
	if xwidth != 0 {
		width = xwidth
	}
	xheight, _ := strconv.Atoi(in.FormValue("height"))
	if xheight != 0 {
		height = xheight
	}
	xtypes := in.FormValue("type")
	if len(xtypes) == 0 {
		f = orignal
	} else {
		switch xtypes {
		case "eggbox":
			f = eggbox
		case "saddle":
			f = saddle
		case "orignal":
			f = orignal
		case "moguls":
			f = moguls
		default:
			f = orignal
		}
	}
	xyscale = float64(width) / 2.0 / xyrange
	zscale = float64(height) * 0.4
	fmt.Printf("Set config's as height[%d],width[%d],func[%s]\n", height, width, xtypes)
}

func svg(w http.ResponseWriter, f function) {
	zmin, zmax := getminmax(f)
	fmt.Println("Start draw")
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	w.Write([]byte(s))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, f)
			bx, by, bz := corner(i, j, f)
			cx, cy, cz := corner(i, j+1, f)
			dx, dy, dz := corner(i+1, j+1, f)
			maxz := math.Max(math.Max(az, bz), math.Max(cz, dz))
			minz := math.Min(math.Min(az, bz), math.Min(cz, dz))
			// for Ex3.1
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			s = fmt.Sprintf("<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color(maxz, minz, zmax, zmin),
				ax, ay, bx, by, cx, cy, dx, dy)
			w.Write([]byte(s))
		}
	}
	w.Write([]byte(fmt.Sprintln("</svg>")))
	fmt.Println("Draw end")
}

func getminmax(f function) (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i <= cells; i++ {
		for j := 0; j <= cells; j++ {
			x := xyrange * (float64(i)/float64(cells) - 0.5)
			y := xyrange * (float64(j)/float64(cells) - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}
	return
}
func color(maxz, minz, zmax, zmin float64) string {
	color := ""
	if math.Abs(maxz) > math.Abs(minz) {
		red := math.Exp(math.Abs(maxz)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(minz)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}
func corner(i, j int, f function) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func orignal(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r != 0 {
		return math.Sin(r) / r
	}
	return 0
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 12.0
	b := 10.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}
func moguls(x, y float64) float64 {
	return 0
}
