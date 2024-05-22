package main

import (
	"image/draw"
)

func (d *GifDrawer) StripeBackground() {

	// stripes
	stripes := 4
	stripeW := d.wf / float64(stripes*2)

	offsets := make([]float64, len(d.frames))
	// generate offsets
	for f := range d.frames {
		// this is a bit tricky, as we want to cycle through one white and one red stripe
		// depending on the direction,
		// we begin in the negative stripe width, and add for every step the double amount of a single step

		// offsets[f] = -float64(stripeW) + ((stripeW / float64(len(d.frames)) * 2) * float64(f))
		offsets[f] = float64(stripeW) - ((stripeW / float64(len(d.frames)) * 2) * float64(f))
	}

	for o, f := range d.frames {
		dc := d.PrepareContext(0)

		dc.SetColor(d.background)
		dc.Clear()
		dc.Rotate(0.7)

		for i := 0; i < stripes*3; i++ {
			if i%2 == 0 {
				dc.SetColor(d.primaryColor)
			} else {
				dc.SetColor(d.background)
			}
			dc.DrawRectangle(offsets[o]+(float64(i)*stripeW), -d.hf*4, stripeW, d.hf*8)
			dc.Fill()
		}

		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
		// dc.SavePNG(fmt.Sprintf("out%d.png", o))
	}

}
