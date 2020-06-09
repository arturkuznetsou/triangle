package main
import (
	"image/color"
)


/*
 * Function for drawing horizontal lines
 */
func hline (x1 int, x2 int, y int) {
	for n := x1; n < x2; n++ {
			img.Set(n, y, color.Black)
	}
}
/*
 *
 * Recursive function for splitting a triangle
 * into four smaller ones
 *
 */
func triangulate (x1 int, x2 int, y int) {
	xma := x2 - x1 + 1
	if xma > 9 {
		xa := (x1 + x2) / 2
		hline(x1 + xma / 4, x2 - xma / 4, y - xma / 2)
		dline(x1 + xma / 4 - 1, y - xma / 2 + 1, -xma / 2)
		dline(x2 - xma / 2 + 1, y, xma / 2)
		triangulate(x1, x1 + (x2 - x1) / 2, y)
		triangulate(xa + 1, x2, y)
		triangulate(x1 + xma / 4, x2 - xma / 4, y - xma / 2)
	}
}
/*
 *
 * Draw sierpinski triangle
 *
 */
func draw () {
	hline(1, x * 8 - 1, x * 8)
	dline(1, x * 8, x * 8)
	dline(x * 4, 1, -x * 8)
	triangulate(1, x * 8, x * 8)
}
/*
 *
 *
 * Function for drawing diagonal lines.
 * Starts at x, y and moves diagonaly
 * for each integer in height
 * Always goes to the right.
 *
 *
 */

func dline (x int, y int, height int) {
	add := 1
	n := 0
	height += height % 2
	if height < 0 {
		add = -1
	}
	for n != height {
		img.Set(x, y - n, color.Black)
		img.Set(x, y - n - add, color.Black)
		x += 1;
		n += 2 * add
	}
}

