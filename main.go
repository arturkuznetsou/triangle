package main
import (
	i "image"
	"image/png"
	"os"
	"math"
)


var img *i.RGBA
func triangulate (y int, x1 int, x2 int) {
	if math.Abs(float64(x2 - x1)) > 7 {
	}
}
func setup () {
}
func main () {
	var depth float64
	depth = 9
	x := int(math.Pow(2, depth - 1))
	upLeft := i.Point{0, 0}
	lowRight := i.Point{x * 8 + 1, x * 8 + 9}
	img = i.NewRGBA(i.Rectangle{upLeft, lowRight})

	hline(1, x * 8 - 1, x * 8)
	dline(1, x * 8, x * 8)
	dline(x * 4, 1, -x * 8)


	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
