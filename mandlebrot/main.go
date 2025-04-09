package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

var (
	baseColor = color.RGBA{
		R: 35,
		G: 33,
		B: 54,
		A: 255,
	}
	irisColor = color.RGBA{
		R: 196,
		G: 167,
		B: 231,
		A: 255,
	}
)

type mandlebrot struct {
	img       *image.RGBA
	iScale    float64
	realScale float64
	iLower    float64
	iUpper    float64
	realLower float64
	realUpper float64
}

func new(width int, height int, iLower float64, iUpper float64, realLower float64, realUpper float64) *mandlebrot {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := range width {
		for y := range height {
			img.SetRGBA(x, y, baseColor)
		}
	}
	return &mandlebrot{
		img:       img,
		iLower:    iLower,
		iUpper:    iUpper,
		realLower: realLower,
		realUpper: realUpper,
		iScale:    math.Abs(iLower-iUpper) / float64(height),
		realScale: math.Abs(realLower-realUpper) / float64(width),
	}
}

func solve(real0 float64, i0 float64, real1 float64, i1 float64) []float64 {
	realPart := (real0 * real0) + (-1 * i0 * i0)
	imagPart := (real0 * i0) * 2
	realPart += real1
	imagPart += i1

	return []float64{realPart, imagPart}
}

func (m *mandlebrot) plot() {
	for x := range m.img.Rect.Max.X {
		for y := range m.img.Rect.Max.Y {
			iCoord := m.iUpper - float64(y)*m.iScale
			realCoord := m.realLower + float64(x)*m.realScale

			pointOutside := false
			cNumber := []float64{0, 0}
			its := 0

			for its < 25 && !pointOutside {
				cNumber = solve(cNumber[0], cNumber[1], realCoord, iCoord)
				pointOutside = math.Sqrt(cNumber[0]*cNumber[0]+cNumber[1]*cNumber[1]) > 2.0
				its++
			}

			if !pointOutside {
				m.img.SetRGBA(x, y, irisColor)
			}
		}
	}
}

func main() {
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatalf("unable to create image: %v", err)
	}

	m := new(2160, 2160, -1.0, 1.0, -1.5, 0.5)
	m.plot()
	if err := png.Encode(f, m.img); err != nil {
		f.Close()
		log.Fatalf("unable to encode image: %v", err)
	}

	if err := f.Close(); err != nil {
		log.Fatalf("unable to close file: %v", err)
	}
}
