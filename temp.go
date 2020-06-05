package main
import (
	i "image"
	//"image/color"
	"image/png"
	"os"
	"math"
	"fmt"
)


var img *i.RGBA
func triangulate (y int, x1 int, x2 int) {
	if math.Abs(float64(x2 - x1)) > 7 {
		xa := (x1 + x2) / 2
		triangulate(y, x1, xa)
		triangulate(y, xa + 2, x2)
		triangulate(y - xa - 1, xa / 2 + 1, 3 / 4 * xa)
	}
}
func setup () {
}
func main () {
	var depth float64
	fmt.Scan(&depth)
	x := int(math.Pow(2, depth - 1))
	upLeft := i.Point{0, 0}
	lowRight := i.Point{x * 7, x * 8}
	img = i.NewRGBA(i.Rectangle{upLeft, lowRight})
	// Encode as PNG.
	triangulate(x * 7, 0, x * 8)
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
