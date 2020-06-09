package main
import (
	i "image"
	"image/png"
	"image/color"
	"os"
	"math"
)


var img *i.RGBA
var x int
var col color.Gray16

func main () {
	var depth float64
	depth = 10
	col = color.Black
	x = int(math.Pow(2, depth - 1))
	upLeft := i.Point{0, 0}
	lowRight := i.Point{x * 8 + 1, x * 8 + 9}
	img = i.NewRGBA(i.Rectangle{upLeft, lowRight})

	draw()

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
