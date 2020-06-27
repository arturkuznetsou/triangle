package main
import (
	i "image"
	"image/gif"
	"image/color"
	sc "strconv"
	"os"
	"math"
	"log"
)


var img *i.RGBA
var x int
var col color.RGBA
var bgcol color.RGBA

func main () {
	depth := 9


	col = ParseHex("000")
	bgcol = ParseHex("fff")

	for n := 1; n < len(os.Args); n++ {
		arg := os.Args[n]
		switch arg{
		case "color":
			col = ParseHex(os.Args[n + 1])
			n += 1
		case "background":
			bgcol = ParseHex(os.Args[n + 1])
			n += 1
		case "depth":
			depth, _ = sc.Atoi(os.Args[n + 1])
			n += 1
		default:
			l := log.New(os.Stderr, "", 0)
			l.Fatalf("Invalid option: '%s'.\nTry 'gensir help' for more information.", arg)
			n += 1
		}
	}





	x = int(math.Pow(2, float64(depth - 1)))
	upLeft := i.Point{0, 0}
	lowRight := i.Point{x * 8 + 1, x * 8 + 9}


	img = i.NewRGBA(i.Rectangle{upLeft, lowRight})

	for z := x * 8 + 1; z >= 0; z-- {
		for y := x * 8 + 9; y >= 0; y-- {
			img.Set(z, y, bgcol)
		}
	}

	draw()

	opts := gif.Options{255, nil, nil}
	f, _ := os.Create("image.gif")
	gif.Encode(f, img, &opts)
}
