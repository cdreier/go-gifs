# gif-tools

create your own beautiful gifs

note: frame count controls the animation speed, more frames = slower animation


## here are some starting points

nice! with pastel stripe background
```
go run . -f=12 -x=200 -y=150 -c=FF99C8 -bg=FCF6BD -txt=nice! -tslope=0.3
```

wobbly PTAL with whirly background
```
go run . -f=14 -x=300 -y=300 -bg=3a86ff -c=ffbe0b -e=whirl -txt=PTAL -te=wobbly -tsize=120
```

TY with rotating star on purple background
```
go run . -f=18 -x=300 -y=300 -bg=3c096c -e=star -txt=TY -tsize=120 -tslope=-0.2
```

our labdrador charly, waiting for me to go out
```
go run . -f=6 -x=400 -y=400 -c=4392f1 -bg=4b0082 -i=charly_waiting.png
```


## all options

```
GLOBAL OPTIONS:
   --http                                                   starting a webservice instead of cli (default: false)
   --filename value, --out value, --name value              the filename of the gif (default: "out.gif")
   --width value, -x value                                  the gifs width (default: 300)
   --height value, -y value                                 the gifs height (default: 300)
   --frames value, -f value                                 number of frames, this controls the speed of the animation (default: 6)
   --background value, --bg value                           background color (default: "#fff")
   --color value, -c value                                  foreground color (default: "#ff0000")
   --effect value, -e value                                 text effect, possible values are stripes, star or whirl (default: "stripes")
   --text value, --txt value                                your text
   --txteffect value, --te value                            text effect, possible values are wobbly or rotate
   --txtsize value, --tsize value                           font size (default: 80)
   --txtslope value, --tslope value                         slope of the text (default: 0)
   --image value, --imagepath value, --img value, -i value  path to image
   --imageEffect value, --imgEffect value, --ie value       the image animation, possible values are wobbly or rotate
   --help, -h                                               show help
```