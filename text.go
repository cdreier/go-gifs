package main

import (
	"image"
	"image/draw"
)

func (d *GifDrawer) Text(txt string, fontSize, rot float64) {

	for _, f := range d.frames {
		img := d.textAsImage(txt, fontSize, rot)
		draw.Draw(f, f.Bounds(), img, f.Bounds().Min, draw.Over)
	}
}

func (d *GifDrawer) textAsImage(txt string, fontSize, rot float64) image.Image {
	dc := d.PrepareContext(fontSize)
	n := 6 // "stroke" size
	dc.SetRGB(0, 0, 0)
	dc.RotateAbout(rot, d.wf/2, d.hf/2)
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := d.wf/2 + float64(dx)
			y := d.hf/2 + float64(dy)
			dc.DrawStringAnchored(txt, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(txt, d.wf/2, d.hf/2, 0.5, 0.5)
	return dc.Image()
}

func (d *GifDrawer) RotatingText(txt string, fontSize float64) {

	rotationSteps := make([]float64, len(d.frames))
	step := radians360deg / float64(len(d.frames))
	for f := range d.frames {
		rotationSteps[f] = float64(f) * step
	}

	for i, f := range d.frames {
		img := d.textAsImage(txt, fontSize, rotationSteps[i])
		draw.Draw(f, f.Bounds(), img, f.Bounds().Min, draw.Over)
	}
}

func (d *GifDrawer) WobblyText(txt string, fontSize float64) {

	rotationSteps := make([]float64, len(d.frames))
	from := -0.3
	to := 0.3
	for f := range d.frames {
		if f < len(d.frames)/2 {
			rotationSteps[f] = float64(f) * (from / float64(len(d.frames)))
		} else {
			rotationSteps[f] = float64(f) * (to / float64(len(d.frames)))
		}
	}

	for i, f := range d.frames {
		img := d.textAsImage(txt, fontSize, rotationSteps[i])
		draw.Draw(f, f.Bounds(), img, f.Bounds().Min, draw.Over)
	}
}
