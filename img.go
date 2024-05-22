package main

import (
	"image/draw"
	"log"

	"github.com/fogleman/gg"
)

func (d *GifDrawer) Image(path string) {

	img, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	resizedImg := d.resizedImage(img)

	imgPosX := 0
	imgPosY := 0

	imgPosY = (d.h / 2) - (resizedImg.Bounds().Size().Y / 2)
	imgPosX = (d.w / 2) - (resizedImg.Bounds().Size().X / 2)

	for _, f := range d.frames {
		dc := d.PrepareContext(0)
		dc.DrawImage(resizedImg, imgPosX, imgPosY)
		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
	}
}

func (d *GifDrawer) RotatingImage(path string) {

	img, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	resizedImg := d.resizedImage(img)

	imgPosX := 0
	imgPosY := 0

	imgPosY = (d.h / 2) - (resizedImg.Bounds().Size().Y / 2)
	imgPosX = (d.w / 2) - (resizedImg.Bounds().Size().X / 2)

	rotationSteps := make([]float64, len(d.frames))
	step := radians360deg / float64(len(d.frames))
	for f := range d.frames {
		rotationSteps[f] = float64(f) * step
	}

	for i, f := range d.frames {
		dc := d.PrepareContext(0)
		dc.RotateAbout(rotationSteps[i], d.wf/2, d.hf/2)
		dc.DrawImage(resizedImg, imgPosX, imgPosY)
		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
	}
}
func (d *GifDrawer) WobblyImage(path string) {

	img, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	resizedImg := d.resizedImage(img)

	imgPosX := 0
	imgPosY := 0

	imgPosY = (d.h / 2) - (resizedImg.Bounds().Size().Y / 2)
	imgPosX = (d.w / 2) - (resizedImg.Bounds().Size().X / 2)

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
		dc := d.PrepareContext(0)
		dc.RotateAbout(rotationSteps[i], d.wf/2, d.hf/2)
		dc.DrawImage(resizedImg, imgPosX, imgPosY)
		draw.Draw(f, f.Bounds(), dc.Image(), f.Bounds().Min, draw.Over)
	}
}
