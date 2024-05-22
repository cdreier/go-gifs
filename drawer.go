package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/ericpauley/go-quantize/quantize"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

type Gif struct {
	frames       []*image.Paletted
	w            int
	h            int
	wf           float64
	hf           float64
	background   color.Color
	primaryColor color.Color
	palette      []color.Color
}

type GifDrawer struct {
	*Gif
}

var gifPalette = []color.Color{
	color.Black,
	color.RGBA{255, 0, 0, 255},    // red
	color.RGBA{255, 228, 58, 255}, // star yellow
	color.RGBA{255, 204, 0, 255},  // star orange
	color.White,
}

func NewGif(w, h, frames int) *Gif {
	d := &Gif{
		w:            w,
		h:            h,
		wf:           float64(w),
		hf:           float64(h),
		frames:       make([]*image.Paletted, frames),
		background:   color.White,
		primaryColor: color.Black,
		palette:      gifPalette,
	}
	return d
}

type Colors struct {
	background   color.Color
	primaryColor color.Color
}

func (d *Gif) Drawer() *GifDrawer {
	for i := range d.frames {
		d.frames[i] = image.NewPaletted(image.Rect(0, 0, d.w, d.h), d.palette)
	}
	return &GifDrawer{
		d,
	}
}

func (d *Gif) resizedImage(img image.Image) image.Image {
	smallerSide := d.h
	if d.w < smallerSide {
		smallerSide = d.w
	}

	imgW := 0
	imgH := 0
	if img.Bounds().Size().X > img.Bounds().Size().Y {
		// landscape
		imgW = smallerSide
	} else {
		// portrait
		imgH = smallerSide
	}

	return imaging.Resize(img, imgW, imgH, imaging.Lanczos)
}
func (d *Gif) WithImageColors(path string) *Gif {
	img, err := gg.LoadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	resizedImg := d.resizedImage(img)
	q := quantize.MedianCutQuantizer{}
	targetPalette := make([]color.Color, 0, 256-len(d.palette))
	d.palette = append(d.palette, q.Quantize(targetPalette, resizedImg)...)

	return d
}

func (d *Gif) WithColors(c Colors) *Gif {
	if c.primaryColor != nil {
		d.primaryColor = c.primaryColor
	}
	if c.background != nil {
		d.background = c.background
	}
	d.palette = append(d.palette, d.primaryColor, d.background)
	return d
}

func (d *Gif) Save(filename string) error {
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	return d.Write(f)
}

func (d *Gif) Write(wr io.Writer) error {
	return gif.EncodeAll(wr, &gif.GIF{
		Image: d.frames,
		Delay: delays(len(d.frames), 0),
	})
}

func (d *Gif) PrepareContext(fontSize float64) *gg.Context {
	dc := gg.NewContext(d.w, d.h)

	if fontSize > 0 {
		// setfontface
		f, err := truetype.Parse(ImpactFont)
		if err != nil {
			log.Fatal(err)
		}
		face := truetype.NewFace(f, &truetype.Options{
			Size: fontSize,
			// Hinting: font.HintingFull,
		})
		dc.SetFontFace(face)
	}
	return dc
}

func delays(size, del int) []int {
	ds := make([]int, size)
	for d := range ds {
		ds[d] = del
	}
	return ds
}
