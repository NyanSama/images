// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type function func(x, y float64) float64

func svg(w io.Writer, f function) {
	zmin, zmax := getminmax(f)
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
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
			fmt.Fprintf(w, "<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color(maxz, minz, zmax, zmin),
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func getminmax(f function) (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i <= cells; i++ {
		for j := 0; j <= cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
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
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
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
func main() {
	usage := "Pleas input type: orignal,eggbox,moguls,saddle"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdout, usage)
		os.Exit(1)
	}
	var f function
	switch os.Args[1] {
	case "eggbox":
		f = eggbox
	case "saddle":
		f = saddle
	case "orignal":
		f = orignal
	case "moguls":
		f = moguls
	}
	svg(os.Stdout, f)
}

//!-
