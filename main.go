package main
import (
	i "image"
	"image/gif"
	"image/color"
	sc "strconv"
	"os"
	"math"
)


var img *i.RGBA
var x int
var col color.RGBA

func main () {
	var depth float64
	depth = 9
	x = int(math.Pow(2, depth - 1))

	r := 255
	b := 255
	g := 255
	for n := 1; n < len(os.Args); n++ {
		for _, lett := range os.Args[n][1:] {
			switch lett {
				case 'r':
					r, _ = sc.Atoi(os.Args[n + 1])
				case 'g':
					g, _ = sc.Atoi(os.Args[n + 1])
				case 'b':
					b, _ = sc.Atoi(os.Args[n + 1])
			}
		}
		n += 1
	}
	col = color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
	upLeft := i.Point{0, 0}
	lowRight := i.Point{x * 8 + 1, x * 8 + 9}


	img = i.NewRGBA(i.Rectangle{upLeft, lowRight})

	for z := x * 8 + 1; z >= 0; z-- {
		for y := x * 8 + 9; y >= 0; y-- {
			img.Set(z, y, col)
		}
	}

	draw()

	opts := gif.Options{255, nil, nil}
	f, _ := os.Create("image.gif")
	gif.Encode(f, img, &opts)
}
