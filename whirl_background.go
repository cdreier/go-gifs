package main

import (
	"image/draw"
	"math"
)

func (d *GifDrawer) WhirlBackground() {

	biggerSide := d.hf
	if d.wf > biggerSide {
		biggerSide = d.wf
	}
	N := biggerSide * 2

	rotationSteps := make([]float64, len(d.frames))
	step := radians360deg / float64(len(d.frames))
	for f := range d.frames {
		rotationSteps[f] = float64(f) * step
	}

	for i, f := range d.frames {
		dc := d.PrepareContext(0)

		dc.SetColor(d.background)
		dc.Clear()
		dc.SetColor(d.primaryColor)

		dc.RotateAbout(rotationSteps[i], d.wf/2, d.hf/2)

		for i := 0; i <= int(N); i++ {
			t := float64(i) / N
			dd := t*float64(biggerSide)*0.7 + 10
			a := t * math.Pi * 2 * 20
			x := d.wf/2 + math.Cos(a)*dd
			y := d.hf/2 + math.Sin(a)*dd
			r := t * 8
			dc.DrawCircle(x, y, r)
		}
		dc.Fill()
		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
	}
}
