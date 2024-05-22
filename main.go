package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "embed"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/icza/gox/imagex/colorx"
	"github.com/urfave/cli/v2"
)

const radians360deg = 6.2

//go:embed impact/impact.ttf
var ImpactFont []byte

func main() {

	app := cli.NewApp()
	app.Usage = "Create animated gifs with text and images"
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "http",
			Value: false,
			Usage: "starting a webservice instead of cli",
		},
		&cli.StringFlag{
			Name:    "filename",
			Aliases: []string{"out", "name"},
			Usage:   "the filename of the gif",
			Value:   "out.gif",
		},
		&cli.IntFlag{
			Name:    "width",
			Aliases: []string{"x"},
			Usage:   "the gifs width",
			Value:   300,
		},
		&cli.IntFlag{
			Name:    "height",
			Aliases: []string{"y"},
			Usage:   "the gifs height",
			Value:   300,
		},
		&cli.IntFlag{
			Name:    "frames",
			Aliases: []string{"f"},
			Usage:   "number of frames, this controls the speed of the animation",
			Value:   6,
		},
		&cli.StringFlag{
			Name:    "background",
			Aliases: []string{"bg"},
			Usage:   "background color",
			Value:   "#fff",
		},
		&cli.StringFlag{
			Name:    "color",
			Aliases: []string{"c"},
			Usage:   "foreground color",
			Value:   "#ff0000",
		},
		&cli.StringFlag{
			Name:    "effect",
			Aliases: []string{"e"},
			Usage:   "text effect, possible values are stripes, star or whirl",
			Value:   "stripes",
		},
		&cli.StringFlag{
			Name:    "text",
			Aliases: []string{"txt"},
			Usage:   "your text",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "txteffect",
			Aliases: []string{"te"},
			Usage:   "text effect, possible values are wobbly or rotate",
			Value:   "",
		},
		&cli.IntFlag{
			Name:    "txtsize",
			Aliases: []string{"tsize"},
			Usage:   "font size",
			Value:   80,
		},
		&cli.Float64Flag{
			Name:    "txtslope",
			Aliases: []string{"tslope"},
			Usage:   "slope of the text",
			Value:   0.0,
		},
		&cli.StringFlag{
			Name:    "image",
			Aliases: []string{"imagepath", "img", "i"},
			Usage:   "path to image",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "imageEffect",
			Aliases: []string{"imgEffect", "ie"},
			Usage:   "the image animation, possible values are wobbly or rotate",
			Value:   "",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

type urlParamInput interface {
	string | int | float64
}

func urlParam[T urlParamInput](r *http.Request, key string, fallback T) T {
	val := r.URL.Query().Get(key)
	if val == "" {
		return fallback
	}
	if _, ok := any(fallback).(int); ok {
		tmp, _ := strconv.ParseInt(val, 10, 32)
		i := int(tmp)
		return any(i).(T)
	}
	if _, ok := any(fallback).(float64); ok {
		tmp, _ := strconv.ParseFloat(val, 32)
		i := float64(tmp)
		return any(i).(T)
	}
	return any(val).(T)
}

func urlParamMax(r *http.Request, key string, fallback, max int) int {
	val := urlParam(r, key, fallback)
	if val > max {
		return max
	}
	return val
}

//go:embed index.html
var indexHTML string

func router() error {
	r := chi.NewRouter()

	r.Get("/", func(wr http.ResponseWriter, r *http.Request) {
		render.HTML(wr, r, indexHTML)
	})

	r.Get("/gif", func(wr http.ResponseWriter, r *http.Request) {

		w := urlParamMax(r, "w", 512, 2048)
		h := urlParamMax(r, "h", 512, 2048)
		f := urlParamMax(r, "f", 6, 30)
		bg := urlParam(r, "bg", "#fff")
		c := urlParam(r, "c", "#ff0000")
		effect := urlParam(r, "effect", "stripes")
		txt := urlParam(r, "txt", "")
		txtsize := urlParam(r, "txtsize", 80)
		txtslope := urlParam(r, "txtslope", 0.0)
		txtEffect := urlParam(r, "txteffect", "")

		if !strings.HasPrefix(bg, "#") {
			bg = fmt.Sprintf("#%s", bg)
		}
		if !strings.HasPrefix(c, "#") {
			c = fmt.Sprintf("#%s", c)
		}

		bgc, _ := colorx.ParseHexColor(bg)
		cc, _ := colorx.ParseHexColor(c)

		d := NewGif(w, h, f).WithColors(Colors{
			background:   bgc,
			primaryColor: cc,
		}).Drawer()

		applyEffect(d, effect)
		applyText(d, txt, txtEffect, float64(txtsize), txtslope)

		wr.Header().Add("Content-Type", "image/gif")
		d.Write(wr)
	})

	log.Println("starting on 8080")
	return http.ListenAndServe(":8080", r)
}

func applyEffect(d *GifDrawer, effect string) {
	switch effect {
	case "stripes":
		d.StripeBackground()
	case "star":
		d.StarBackground()
	case "whirl":
		d.WhirlBackground()
	}
}
func applyText(d *GifDrawer, txt, txtEffect string, txtsize, txtslope float64) {
	if txt != "" {
		switch txtEffect {
		case "":
			d.Text(txt, float64(txtsize), float64(txtslope))
		case "wobbly":
			d.WobblyText(txt, txtsize)
		case "rotate":
			d.RotatingText(txt, txtsize)
		}
	}
}

func run(c *cli.Context) error {

	if c.Bool("http") {
		return router()
	}

	w := c.Int("width")
	h := c.Int("height")
	f := c.Int("frames")
	bg := c.String("background")
	col := c.String("color")
	effect := c.String("effect")
	txt := c.String("text")

	txtsize := c.Int("txtsize")
	txtslope := c.Float64("txtslope")
	txtEffect := c.String("txteffect")
	imgPath := c.String("image")
	imgEffect := c.String("imageEffect")
	out := c.String("filename")

	if !strings.HasPrefix(bg, "#") {
		bg = fmt.Sprintf("#%s", bg)
	}
	if !strings.HasPrefix(col, "#") {
		col = fmt.Sprintf("#%s", col)
	}

	bgc, _ := colorx.ParseHexColor(bg)
	cc, _ := colorx.ParseHexColor(col)

	// setting up gif with colors
	g := NewGif(w, h, f).WithColors(Colors{
		background:   bgc,
		primaryColor: cc,
	})

	if imgPath != "" {
		g.WithImageColors(imgPath)
	}

	d := g.Drawer()

	applyEffect(d, effect)

	if imgPath != "" {
		switch imgEffect {
		case "":
			d.Image(imgPath)
		case "wobbly":
			d.WobblyImage(imgPath)
		case "rotate":
			d.RotatingImage(imgPath)
		}
	}

	applyText(d, txt, txtEffect, float64(txtsize), txtslope)

	return d.Save(out)
}
