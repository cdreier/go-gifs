package main

import (
	"image/draw"
	"math"
)

type starPoint struct {
	X, Y float64
}

func starPolygon(n int) []starPoint {
	result := make([]starPoint, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = starPoint{math.Cos(a), math.Sin(a)}
	}
	return result
}

func (d *GifDrawer) StarBackground() {

	rotationSteps := make([]float64, len(d.frames))
	step := radians360deg / float64(len(d.frames))
	for f := range d.frames {
		rotationSteps[f] = float64(f) * step
	}

	smallerSide := d.hf
	if d.wf < smallerSide {
		smallerSide = d.wf
	}
	padding := float64(smallerSide / 20)

	for i, f := range d.frames {
		dc := d.PrepareContext(0)

		dc.SetColor(d.background)
		dc.Clear()

		n := 5
		points := starPolygon(n)

		dc.Translate(d.wf/2, d.hf/2)
		dc.Scale((smallerSide-padding)/2, (smallerSide-padding)/2)
		dc.Rotate(rotationSteps[i])

		for i := 0; i < n+1; i++ {
			index := (i * 2) % n
			p := points[index]
			dc.LineTo(p.X, p.Y)
		}
		dc.SetLineWidth(10)
		dc.SetHexColor("#FFCC00")
		dc.StrokePreserve()
		dc.SetHexColor("#FFE43A")
		dc.Fill()

		// dc.SavePNG(fmt.Sprintf("out%d.png", i))
		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
	}

}
